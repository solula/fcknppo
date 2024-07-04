//1

package dto

import (
	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenPayload struct {
	SID      string // Уникальный идентификатор сессии
	UserUuid string // Uuid пользователя
}

// AccessTokenClaims тип данных токена доступа
type AccessTokenClaims struct {
	jwt.RegisteredClaims
	Payload *AccessTokenPayload `json:"payload"`
}

// RefreshTokenClaims тип данных токена обновления
type RefreshTokenClaims struct {
	jwt.RegisteredClaims
	_ struct{}
}

// EmailVerificationTokenClaims тип данных токена подтверждения почты
type EmailVerificationTokenClaims struct {
	jwt.RegisteredClaims
	UserUuid string `json:"user_uuid"`
}
