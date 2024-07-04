//1

package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/modules/domain/auth/dto"
	"waterfall-backend/internal/modules/features/config"
)

type Service struct {
	secret string
}

func NewService(cfg config.Config) *Service {
	return &Service{
		secret: cfg.JWTSecret,
	}
}

// ParseAccessToken парсит токен доступа
func (r *Service) ParseAccessToken(jwtToken string) (*dto.AccessTokenClaims, error) {
	var claims dto.AccessTokenClaims
	err := r.parseToken(jwtToken, &claims)
	if err != nil {
		return nil, err
	}

	return &claims, nil
}

// ParseRefreshToken парсит токен обновления
func (r *Service) ParseRefreshToken(jwtToken string) (*dto.RefreshTokenClaims, error) {
	var claims dto.RefreshTokenClaims
	err := r.parseToken(jwtToken, &claims)
	if err != nil {
		return nil, err
	}

	return &claims, nil
}

// ParseEmailVerificationToken парсит токен подтверждения почты
func (r *Service) ParseEmailVerificationToken(jwtToken string) (*dto.EmailVerificationTokenClaims, error) {
	var claims dto.EmailVerificationTokenClaims
	err := r.parseToken(jwtToken, &claims)
	if err != nil {
		return nil, err
	}

	return &claims, nil
}

// ValidateToken проверяет валидность токена
func (r *Service) ValidateToken(jwtToken string) error {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%w: неожиданный метод подписи: %v", err_const.ErrInvalidToken, token.Header["alg"])
		}
		return []byte(r.secret), nil
	})
	if err != nil {
		return fmt.Errorf("%w: %s", err_const.ErrInvalidToken, err)
	}

	if !token.Valid {
		return err_const.ErrInvalidToken
	}

	return nil
}

// CreateToken создает токен
func (r *Service) CreateToken(tokenClaims jwt.Claims) (string, error) {
	return r.createToken(tokenClaims)
}

// CreateTokenPair создает пару токенов
func (r *Service) CreateTokenPair(payload *dto.AccessTokenPayload, accessTokenExp time.Duration, refreshTokenExp time.Duration) (*dto.Tokens, error) {
	// Определяем общий uuid для токенов
	tokenUuid := uuid.NewString()
	// Дата создания токена
	issuedAt := time.Now()

	accessTokenClaims := dto.AccessTokenClaims{}
	accessTokenClaims.Payload = payload
	accessTokenClaims.IssuedAt = &jwt.NumericDate{Time: issuedAt}
	accessTokenClaims.ExpiresAt = &jwt.NumericDate{Time: issuedAt.Add(accessTokenExp)}
	accessTokenClaims.ID = tokenUuid

	accessTokenString, err := r.createToken(accessTokenClaims)
	if err != nil {
		return nil, err
	}

	refreshTokenClaims := dto.RefreshTokenClaims{}
	refreshTokenClaims.IssuedAt = &jwt.NumericDate{Time: issuedAt}
	refreshTokenClaims.ExpiresAt = &jwt.NumericDate{Time: issuedAt.Add(refreshTokenExp)}
	refreshTokenClaims.ID = tokenUuid

	refreshTokenString, err := r.createToken(refreshTokenClaims)
	if err != nil {
		return nil, err
	}

	return &dto.Tokens{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

// createToken создает токен
func (r *Service) createToken(tokenClaims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString([]byte(r.secret))
	if err != nil {
		return "", fmt.Errorf("%w: ошибка генерации токена", err)
	}

	return tokenString, nil
}

// parseToken парсит токен и заполняет claims
func (r *Service) parseToken(jwtToken string, claims jwt.Claims) error {
	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%w: неожиданный метод подписи: %v", err_const.ErrInvalidToken, token.Header["alg"])
		}
		return []byte(r.secret), nil
	})
	if err != nil {
		return fmt.Errorf("%w: %s", err_const.ErrInvalidToken, err)
	}

	if !token.Valid {
		return err_const.ErrInvalidToken
	}

	return nil
}
