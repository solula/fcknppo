package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/api/idtoken"
	"io"
	"math"
	"net/http"
	"net/mail"
	"strconv"
	"time"
	"unicode"
	"waterfall-backend/internal/constants/err_const"
	files_const "waterfall-backend/internal/constants/files"
	"waterfall-backend/internal/models"
	"waterfall-backend/internal/models/auth"
	"waterfall-backend/internal/models/files"
	"waterfall-backend/internal/models/permissions"
	"waterfall-backend/internal/models/roles"
	"waterfall-backend/internal/models/session"
	"waterfall-backend/internal/modules/domain/auth/dto"
	user "waterfall-backend/internal/modules/domain/user/dto"
	"waterfall-backend/internal/modules/features/config"
	"waterfall-backend/internal/modules/features/logger"
	email "waterfall-backend/internal/modules/services/email/dto"
	fs "waterfall-backend/internal/modules/services/fs/dto"
	"waterfall-backend/internal/modules/services/fs/i_fs"
	"waterfall-backend/internal/modules/services/token"
	"waterfall-backend/internal/pkg/transaction"
	"waterfall-backend/internal/utils"
	"waterfall-backend/internal/utils/password"
	"waterfall-backend/internal/utils/ptr"
)

// Новые пользователи получают роль бесплатного пользователя
var defaultRoles = []roles.Type{roles.Free}

type IUserRepo interface {
	GetExtendedByUuid(ctx context.Context, uuid string) (*user.ExtendedUser, error)
	GetExtendedByEmail(ctx context.Context, email string) (*user.ExtendedUser, error)
	GetExtendedByVKId(ctx context.Context, vkId int64) (*user.ExtendedUser, error)
	CheckEmailExists(ctx context.Context, email string) (bool, error)
	CreateWithPassword(ctx context.Context, dtm *user.UserPasswordCreate) (*user.User, error)
	CreateWithService(ctx context.Context, dtm *user.UserServiceCreate) (*user.ExtendedUser, error)
	Update(ctx context.Context, uuid string, dtm *user.UserUpdate) (*user.User, error)
	VerifyEmail(ctx context.Context, uuid string) error
	Tx(ctx context.Context) (transaction.TxRepo, transaction.Tx, error)
}

type IRoleRepo interface {
	GetWithPermissionsById(ctx context.Context, id roles.Type) (*user.Role, user.Permissions, error)
}

type IEmailService interface {
	SendVerificationEmail(ctx context.Context, toEmail string, dtm *email.Verification) error
}

type IFileStorageService interface {
	i_fs.IBucketCreator
	CreateFile(ctx context.Context, newFile *fs.NewFile, content io.Reader) (*fs.File, error)
	CopyFile(ctx context.Context, uuid string, toObjectRef *models.ObjectRef, newCreatorUuid string) error
}

type AuthService struct {
	userRepo     IUserRepo
	roleRepo     IRoleRepo
	fs           IFileStorageService
	emailService IEmailService
	tokenService *token.Service
	cfg          config.Config
}

func NewAuthService(userRepo IUserRepo, roleRepo IRoleRepo, fs IFileStorageService, emailService IEmailService, tokenService *token.Service, cfg config.Config) *AuthService {
	return &AuthService{
		userRepo:     userRepo,
		roleRepo:     roleRepo,
		fs:           fs,
		emailService: emailService,
		tokenService: tokenService,
		cfg:          cfg,
	}
}

// AuthAccessToken проверяет доступ к системе по токену доступа
func (r *AuthService) AuthAccessToken(_ context.Context, accessToken string) (*dto.AccessTokenPayload, error) {
	accessTokenClaims, err := r.tokenService.ParseAccessToken(accessToken)
	if err != nil {
		return nil, err
	}

	if accessTokenClaims.Payload == nil {
		return nil, fmt.Errorf("%w: данные токена не найдена", err_const.ErrInvalidToken)
	}

	return accessTokenClaims.Payload, nil
}

// GenerateGuestSession создает гостевую сессию
func (r *AuthService) GenerateGuestSession(ctx context.Context) (*session.Session, error) {
	role, perms, err := r.roleRepo.GetWithPermissionsById(ctx, roles.Guest)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить роль гостя: %w", err)
	}

	// Заполняем гостевого пользователя
	extendedUser := &user.ExtendedUser{
		User: user.User{
			Uuid:     "",
			Email:    ptr.String("guest"),
			Username: "guest",
		},
		Roles:       []*user.Role{role},
		Permissions: perms,
	}

	return r.generateSessionByUser(extendedUser), nil
}

// GenerateUserSession создает сессию пользователя
func (r *AuthService) GenerateUserSession(ctx context.Context, userUuid string) (*session.Session, error) {
	extendedUser, err := r.userRepo.GetExtendedByUuid(ctx, userUuid)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить пользователя: %w", err)
	}

	return r.generateSessionByUser(extendedUser), nil
}

// RefreshAccessToken продлевает токен доступа по паре токенов
func (r *AuthService) RefreshAccessToken(_ context.Context, tokenPair *dto.Tokens) (*dto.Tokens, error) {
	accessTokenClaims, err := r.tokenService.ParseAccessToken(tokenPair.AccessToken)
	if err != nil {
		return nil, err
	}

	refreshTokenClaims, err := r.tokenService.ParseRefreshToken(tokenPair.RefreshToken)
	if err != nil {
		return nil, err
	}

	if accessTokenClaims.ID != refreshTokenClaims.ID {
		return nil, fmt.Errorf("%w: ID токенов не совпадают", err_const.ErrInvalidToken)
	}

	tokenSession := accessTokenClaims.Payload
	if tokenSession == nil {
		return nil, fmt.Errorf("%w: сессия пользователя не найдена", err_const.ErrInvalidToken)
	}

	// При использовании токена обновления срок жизни заново не продлеваем
	refreshTokenLifeTime := time.Until(refreshTokenClaims.ExpiresAt.Time)

	// Генерируем пару токенов
	newTokenPair, err := r.tokenService.CreateTokenPair(tokenSession, r.cfg.JWTAccessTokenLifetime, refreshTokenLifeTime)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать пару токенов: %w", err)
	}

	return newTokenPair, nil
}

// SignInWithGoogle аутентификация с использованием аккаунта Google
func (r *AuthService) SignInWithGoogle(ctx context.Context, idToken string) (*dto.JWT, error) {
	lg := logger.GetFromCtx(ctx).With(zap.String(logger.Operation, "SignInWithGoogle"))

	payload, err := idtoken.Validate(ctx, idToken, r.cfg.GoogleClientID)
	if err != nil {
		return nil, fmt.Errorf("ошибка валидации токена Google: %w", err)
	}

	userCredentials, err := utils.MapToStruct[dto.GoogleUserCredentials](payload.Claims)
	if err != nil {
		return nil, err
	}

	if userCredentials.Email == "" {
		return nil, fmt.Errorf("%w: не указана электронная почта", err_const.ErrValidate)
	}
	if userCredentials.EmailVerified == false {
		return nil, fmt.Errorf("%w: электронная почта не подтверждена", err_const.ErrValidate)
	}

	var currentUser *user.ExtendedUser
	isNewUser := false
	err = transaction.WithTx(ctx, r.userRepo, func(txRepo IUserRepo) error {
		currentUser, err = txRepo.GetExtendedByEmail(ctx, userCredentials.Email)
		if err != nil {
			if errors.Is(err, err_const.ErrNotFound) {
				isNewUser = true
			} else {
				return err
			}
		}

		if isNewUser {
			lg.Info("Пользователя с указанной почтой не существует, регистрируем...", zap.String("email", userCredentials.Email))

			userCreateFull := &user.UserServiceCreate{
				Email:    &userCredentials.Email,
				Roles:    defaultRoles,
				Fullname: userCredentials.Name,
			}

			// Создаем пользователя
			currentUser, err = txRepo.CreateWithService(ctx, userCreateFull)
			if err != nil {
				return fmt.Errorf("не удалось создать пользователя: %w", err)
			}

			// Создаем пользователю бакет
			err = r.fs.CreateBucket(ctx, &models.ObjectRef{
				Type: models.ObjectTypeUsers,
				Ref:  currentUser.Uuid,
			})
			if err != nil {
				return err
			}
		}

		if currentUser.DeletedAt != nil {
			return fmt.Errorf("%w: для входа в систему необходимо восстановить аккаунт", err_const.ErrUserDeleted)
		}

		if currentUser.EmailVerified == false {
			// Вход при помощи стороннего сервиса подтверждает электронную почту
			lg.Info("Электронная почта не верифицирована, подтверждаем...", zap.String("email", userCredentials.Email))
			err = txRepo.VerifyEmail(ctx, currentUser.Uuid)
			if err != nil {
				return err
			}
			currentUser.EmailVerified = true
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	if isNewUser {
		// Загружаем аватар и обновляем пользователя
		_, err := r.createAvatar(ctx, userCredentials.Picture, currentUser.Uuid)
		if err != nil {
			return nil, fmt.Errorf("не удалось загрузить аватар: %w", err)
		}
	}

	// Проверяем, разрешена ли аутентификация
	err = checkAuthPermission(ctx, auth.Google, currentUser)
	if err != nil {
		return nil, err
	}

	tokenPayload := r.generateTokenPayloadByUser(currentUser)

	// Времена жизни токенов берем из конфига
	tokenPair, err := r.tokenService.CreateTokenPair(tokenPayload, r.cfg.JWTAccessTokenLifetime, r.cfg.JWTRefreshTokenLifetime)
	if err != nil {
		return nil, err
	}

	return &dto.JWT{
		Tokens:             *tokenPair,
		AccessTokenPayload: tokenPayload,
	}, nil
}

// SignInWithVK аутентификация с использованием аккаунта VK
func (r *AuthService) SignInWithVK(ctx context.Context, authCredentials *dto.VKAuthCredentials) (*dto.JWT, error) {
	lg := logger.GetFromCtx(ctx).With(zap.String(logger.Operation, "SignInWithVK"))

	vk := api.NewVK(r.cfg.VKServiceToken)

	authParams := params.NewAuthExchangeSilentAuthTokenBuilder().
		Token(authCredentials.SilentToken).
		UUID(authCredentials.Uuid).
		Params
	authResponse, err := vk.AuthExchangeSilentAuthToken(authParams)
	if err != nil {
		return nil, fmt.Errorf("ошибка аутентификации токена VK: %w", err)
	}

	usersGetParams := params.NewUsersGetBuilder().
		UserIDs([]string{strconv.Itoa(authResponse.UserID)}).
		Fields([]string{"photo_max"}).
		Params
	usersCredentials, err := vk.UsersGet(usersGetParams)
	if err != nil {
		return nil, err
	}
	if len(usersCredentials) != 1 {
		return nil, fmt.Errorf("пользователь с указанным VK ID не найден: %d", authResponse.UserID)
	}

	vkId := int64(authResponse.UserID)
	userCredentials := usersCredentials[0]

	var currentUser *user.ExtendedUser
	isNewUser := false
	err = transaction.WithTx(ctx, r.userRepo, func(txRepo IUserRepo) error {
		currentUser, err = txRepo.GetExtendedByVKId(ctx, vkId)
		if err != nil {
			if errors.Is(err, err_const.ErrNotFound) {
				lg.Info("Пользователя с указанным VK ID не существует, регистрируем...", zap.Int64("vkId", vkId))
				isNewUser = true
			} else {
				return err
			}
		}

		if isNewUser {
			userCreateFull := &user.UserServiceCreate{
				VkId:     &vkId,
				Roles:    defaultRoles,
				Fullname: fmt.Sprintf("%s %s", userCredentials.FirstName, userCredentials.LastName),
			}

			// Создаем пользователя
			currentUser, err = txRepo.CreateWithService(ctx, userCreateFull)
			if err != nil {
				return fmt.Errorf("не удалось создать пользователя: %w", err)
			}

			// Создаем пользователю бакет
			err = r.fs.CreateBucket(ctx, &models.ObjectRef{
				Type: models.ObjectTypeUsers,
				Ref:  currentUser.Uuid,
			})
			if err != nil {
				return err
			}
		}

		if currentUser.DeletedAt != nil {
			return fmt.Errorf("%w: для входа в систему необходимо восстановить аккаунт", err_const.ErrUserDeleted)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	if isNewUser {
		// Загружаем аватар и обновляем пользователя
		_, err := r.createAvatar(ctx, userCredentials.PhotoMax, currentUser.Uuid)
		if err != nil {
			return nil, fmt.Errorf("не удалось загрузить аватар: %w", err)
		}
	}

	// Проверяем, разрешена ли аутентификация
	err = checkAuthPermission(ctx, auth.VK, currentUser)
	if err != nil {
		return nil, err
	}

	tokenPayload := r.generateTokenPayloadByUser(currentUser)

	// Времена жизни токенов берем из конфига
	tokenPair, err := r.tokenService.CreateTokenPair(tokenPayload, r.cfg.JWTAccessTokenLifetime, r.cfg.JWTRefreshTokenLifetime)
	if err != nil {
		return nil, err
	}

	return &dto.JWT{
		Tokens:             *tokenPair,
		AccessTokenPayload: tokenPayload,
	}, nil
}

// SignInWithPassword аутентификация с использованием почты и пароля
func (r *AuthService) SignInWithPassword(ctx context.Context, userCredentials *dto.UserCredentials) (*dto.JWT, error) {
	if userCredentials.Email == "" {
		return nil, fmt.Errorf("%w: не указана электронная почта", err_const.ErrValidate)
	}

	currentUser, err := r.userRepo.GetExtendedByEmail(ctx, userCredentials.Email)
	if err != nil {
		if errors.Is(err, err_const.ErrNotFound) {
			return nil, fmt.Errorf("%w: пользователя с указанной почтой не существует", err_const.ErrAuthentication)
		}
		return nil, err
	}

	if currentUser.DeletedAt != nil {
		return nil, fmt.Errorf("%w: для входа в систему необходимо восстановить аккаунт", err_const.ErrUserDeleted)
	}

	// Проверяем, разрешена ли аутентификация с использованием пароля
	err = checkAuthPermission(ctx, auth.Password, currentUser)
	if err != nil {
		return nil, err
	}

	if currentUser.PasswordHash == nil {
		return nil, fmt.Errorf("%w: для аутентификации при помощи пароля должен быть задан пароль пользователя", err_const.ErrAuthentication)
	}

	// Валидация пароля
	if !password.CheckHash(userCredentials.Password, *currentUser.PasswordHash) {
		return nil, fmt.Errorf("%w: неверный пароль", err_const.ErrAuthentication)
	}

	tokenPayload := r.generateTokenPayloadByUser(currentUser)

	// Времена жизни токенов берем из конфига
	tokenPair, err := r.tokenService.CreateTokenPair(tokenPayload, r.cfg.JWTAccessTokenLifetime, r.cfg.JWTRefreshTokenLifetime)
	if err != nil {
		return nil, err
	}

	return &dto.JWT{
		Tokens:             *tokenPair,
		AccessTokenPayload: tokenPayload,
	}, nil
}

// SignUpWithPassword регистрация в системе с использованием пароля
func (r *AuthService) SignUpWithPassword(ctx context.Context, newUser *dto.UserCredentials) (*user.User, error) {
	err := r.emailValidation(ctx, newUser.Email)
	if err != nil {
		return nil, err
	}

	err = r.passwordComplexityValidation(newUser.Password)
	if err != nil {
		return nil, err
	}

	// Формируем хэш пароля
	passwordHash, err := password.GenerateHash(newUser.Password)
	if err != nil {
		return nil, err
	}

	createUser := &user.UserPasswordCreate{
		Email:        newUser.Email,
		Roles:        defaultRoles,
		PasswordHash: &passwordHash,
	}

	// Создаем пользователя
	signedUpUser, err := r.userRepo.CreateWithPassword(ctx, createUser)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать пользователя: %w", err)
	}

	// Создаем пользователю бакет
	err = r.fs.CreateBucket(ctx, &models.ObjectRef{
		Type: models.ObjectTypeUsers,
		Ref:  signedUpUser.Uuid,
	})
	if err != nil {
		return nil, err
	}

	// Создаем пользователю аватар, копируя аватарку по умолчанию
	err = r.fs.CopyFile(ctx, files_const.DefaultAvatarUuid, &models.ObjectRef{
		Type: models.ObjectTypeUsers,
		Ref:  signedUpUser.Uuid,
	}, signedUpUser.Uuid)
	if err != nil {
		return nil, err
	}

	// Отправляем письмо со ссылкой для подтверждения почты
	err = r.sendVerificationEmail(ctx, newUser.Email, signedUpUser.Uuid, r.cfg.EmailVerificationTokenLifetime)
	if err != nil {
		return nil, err
	}

	return signedUpUser, nil
}

// VerifyEmail подтверждение почты пользователем
func (r *AuthService) VerifyEmail(ctx context.Context, verification *dto.EmailVerification) error {
	tokenClaims, err := r.tokenService.ParseEmailVerificationToken(verification.Token)
	if err != nil {
		return fmt.Errorf("не удалось распарсить токен подтверждения почты: %w", err)
	}

	err = r.userRepo.VerifyEmail(ctx, tokenClaims.UserUuid)
	if err != nil {
		return fmt.Errorf("не удалось подтвердить почту: %w", err)
	}

	return nil
}

func (r *AuthService) generateTokenPayloadByUser(extendedUser *user.ExtendedUser) *dto.AccessTokenPayload {
	return &dto.AccessTokenPayload{
		SID:      uuid.NewString(),
		UserUuid: extendedUser.Uuid,
	}
}

func (r *AuthService) generateSessionByUser(extendedUser *user.ExtendedUser) *session.Session {
	var (
		rolesIds        []roles.Type
		permissionsIds  []permissions.Type
		minReleaseDelay time.Duration = math.MaxInt64
	)

	for _, role := range extendedUser.Roles {
		rolesIds = append(rolesIds, role.Id)
		if role.ReleaseDelay < minReleaseDelay {
			minReleaseDelay = role.ReleaseDelay
		}
	}
	for _, permission := range extendedUser.Permissions {
		permissionsIds = append(permissionsIds, permission.Id)
	}

	return &session.Session{
		SID:          uuid.NewString(),
		UserUuid:     extendedUser.Uuid,
		Email:        extendedUser.Email,
		Username:     extendedUser.Username,
		Roles:        rolesIds,
		Permissions:  permissionsIds,
		ReleaseDelay: minReleaseDelay,
	}
}

// checkAuthPermission проверяет разрешена ли аутентификация пользователю user
func checkAuthPermission(ctx context.Context, authType auth.Type, user *user.ExtendedUser) error {
	if user == nil {
		return fmt.Errorf("%w: пользователь не найден", err_const.ErrAuthentication)
	}

	if authType == auth.Password && user.EmailVerified == false {
		logger.GetFromCtx(ctx).Info("Электронная почта не верифицирована, доступ запрещен", zap.Stringp("email", user.Email))
		return fmt.Errorf("%w: для аутентификации необходимо подтвердить email или войти с помощью стороннего сервиса", err_const.ErrEmailNotVerified)
	}

	return nil
}

func (r *AuthService) emailValidation(ctx context.Context, email string) error {
	if email == "" {
		return fmt.Errorf("%w: не указана электронная почта", err_const.ErrInvalidEmail)
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("%w: %s", err_const.ErrInvalidEmail, err)
	}

	exists, err := r.userRepo.CheckEmailExists(ctx, email)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("%w: пользователь с таким email уже существует", err_const.ErrInvalidEmail)
	}

	return nil
}

// passwordComplexityValidation проверяет сложность пароль
func (r *AuthService) passwordComplexityValidation(pass string) error {
	runes := []rune(pass)

	// минимум 5 символов
	if len(runes) < 5 {
		return fmt.Errorf("%w: пароль должен иметь длину как минимум 5 символов", err_const.ErrPasswordTooEasy)
	}

	hasNumbers := false
	hasLetter := false

	for _, r := range runes {
		// минимум одна цифра
		if !hasNumbers && unicode.IsNumber(r) {
			hasNumbers = true
		}
		// минимум одна прописная или заглавная буква
		if !hasLetter && unicode.IsLetter(r) {
			hasLetter = true
		}
	}

	if !hasNumbers || !hasLetter {
		return fmt.Errorf("%w: пароль должен содежрать как минимум одну букву и одну цифру", err_const.ErrPasswordTooEasy)
	}
	return nil
}

// createAvatar создает аватар по ссылке на картинку
func (r *AuthService) createAvatar(ctx context.Context, avatarURL string, userUuid string) (string, error) {
	resp, err := http.Get(avatarURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// TODO: уточнить про mime тип файлов из сервисов
	newFile := &fs.NewFile{
		Filename:    "avatar.png",
		MIMEType:    "image/png",
		Description: "Аватар пользователя из стороннего сервиса",
		ObjectRef: &models.ObjectRef{
			Type: models.ObjectTypeUsers,
			Ref:  userUuid,
		},
		Type:        files.Avatar,
		CreatorUuid: &userUuid,
	}

	createdFile, err := r.fs.CreateFile(ctx, newFile, resp.Body)
	if err != nil {
		return "", err
	}

	return createdFile.Uuid, nil
}

func (r *AuthService) sendVerificationEmail(ctx context.Context, userEmail string, userUuid string, ttl time.Duration) error {
	tokenClaims := &dto.EmailVerificationTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
		},
		UserUuid: userUuid,
	}

	verificationToken, err := r.tokenService.CreateToken(tokenClaims)
	if err != nil {
		return err
	}

	verification := &email.Verification{
		VerificationToken: verificationToken,
	}
	err = r.emailService.SendVerificationEmail(ctx, userEmail, verification)
	if err != nil {
		return fmt.Errorf("не удалось отправить письмо для подтверждения почты: %w", err)
	}

	return nil
}
