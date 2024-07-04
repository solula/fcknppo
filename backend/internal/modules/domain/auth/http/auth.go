package http

import (
	"context"
	"waterfall-backend/internal/modules/domain/auth/dto"
	"waterfall-backend/internal/modules/domain/auth/service"
	user_dto "waterfall-backend/internal/modules/domain/user/dto"
	"waterfall-backend/internal/modules/io/http/auth"
	_ "waterfall-backend/internal/pkg/http/error_handler/http_errors"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func InvokeAuthController(controller *AuthController, router auth.Router) {
	router.Post("/signup/password", controller.SignUpWithPassword)
	router.Post("/login", controller.Login)
	router.Post("/google", controller.SignInWithGoogle)
	router.Post("/vk", controller.SignInWithVK)
	router.Post("/refresh", controller.Refresh)
	router.Post("/verify-email", controller.VerifyEmail)
}

// Login вход в систему
// @tags Auth
// @security ApiKeyAuth
// @in header
// @summary Вход в систему
// @param credentials body dto.UserCredentials true "credentials"
// @success 200 {object} dto.JWT
// @failure 400 {object} http_errors.ErrorResponse
// @failure 401 {object} http_errors.ErrorResponse
// @failure 500 {object} http_errors.ErrorResponse
// @router /auth/login [POST]
func (controller *AuthController) Login(ctx context.Context, credentials *dto.UserCredentials) (*dto.JWT, error) {
	return controller.service.SignInWithPassword(ctx, credentials)
}

func (controller *AuthController) SignInWithGoogle(ctx context.Context, credentials GoogleAuthCredentials) (*dto.JWT, error) {
	return controller.service.SignInWithGoogle(ctx, credentials.IdToken)
}

func (controller *AuthController) SignInWithVK(ctx context.Context, credentials VKAuthCredentials) (*dto.JWT, error) {
	return controller.service.SignInWithVK(ctx, credentials.toDTO())
}

// SignUpWithPassword регистрация в системе
// @tags Auth
// @security ApiKeyAuth
// @in header
// @summary Регистрация нового пользователя (с использованием пароля)
// @param newUser body dto.UserCredentials true "newUser"
// @success 200 {object} dto.User
// @failure 400 {object} http_errors.ErrorResponse
// @failure 401 {object} http_errors.ErrorResponse
// @failure 500 {object} http_errors.ErrorResponse
// @router /auth/signup/password [POST]
func (controller *AuthController) SignUpWithPassword(ctx context.Context, newUser *dto.UserCredentials) (*user_dto.User, error) {
	return controller.service.SignUpWithPassword(ctx, newUser)
}

// Refresh обновление токена доступа
// @tags Auth
// @security ApiKeyAuth
// @in header
// @summary Обновление токена доступа
// @param tokenPair body dto.Tokens true "tokenPair"
// @success 200 {object} dto.User
// @failure 400 {object} http_errors.ErrorResponse
// @failure 401 {object} http_errors.ErrorResponse
// @failure 500 {object} http_errors.ErrorResponse
// @router /auth/refresh [POST]
func (controller *AuthController) Refresh(ctx context.Context, tokenPair *dto.Tokens) (*dto.Tokens, error) {
	return controller.service.RefreshAccessToken(ctx, tokenPair)
}

func (controller *AuthController) VerifyEmail(ctx context.Context, emailVerification *dto.EmailVerification) error {
	return controller.service.VerifyEmail(ctx, emailVerification)
}
