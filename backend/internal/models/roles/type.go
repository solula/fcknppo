package roles

import (
	"slices"
	"waterfall-backend/internal/constants/err_const"
)

type Type string // Тип роли

const (
	Admin   Type = "admin"   // Администратор
	Premium Type = "premium" // Платный пользователь
	Free    Type = "free"    // Бесплатный пользователь
	Guest   Type = "guest"   // Гость (неавторизованный пользователь)
)

func (t Type) Values() []string {
	return []string{
		string(Admin),
		string(Premium),
		string(Free),
		string(Guest),
	}
}

func (t Type) Validate() error {
	if !slices.Contains(t.Values(), string(t)) {
		return err_const.ErrInvalidRole
	}
	return nil
}
