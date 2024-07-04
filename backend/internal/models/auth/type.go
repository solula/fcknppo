package auth

import (
	"slices"
	"waterfall-backend/internal/constants/err_const"
)

type Type string // Тип аутентификации

const (
	Password Type = "password" // Аутентификация с использованием пароля
	Google   Type = "google"   // Аутентификация с использованием аккаунта Google
	VK       Type = "vk"       // Аутентификация с использованием аккаунта VK
)

func (t Type) Values() []string {
	return []string{
		string(Password),
		string(Google),
		string(VK),
	}
}

func (t Type) Validate() error {
	if !slices.Contains(t.Values(), string(t)) {
		return err_const.ErrInvalidRole
	}
	return nil
}
