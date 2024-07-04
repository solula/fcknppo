//1

package dto

import (
	"time"
	"waterfall-backend/internal/models/roles"
	"waterfall-backend/internal/pkg/optional"
)

type User struct {
	Uuid          string     // Uuid пользователя
	CreatedAt     time.Time  // Дата создания
	UpdatedAt     time.Time  // Дата обновления
	DeletedAt     *time.Time // Дата удаления
	Email         *string    // Email
	Fullname      string     // Полное имя
	Username      string     // Имя пользователя в системе
	Score         int        // Количество баллов
	EmailVerified bool       // Признак подтверждения почты
}

// ExtendedUser расширенный пользователь с ролями и правами доступа
type ExtendedUser struct {
	User
	PasswordHash *string // Хэш пароля

	Roles       Roles       // Роли
	Permissions Permissions // Права доступа
}

type Users []*User

type UserPasswordCreate struct {
	Roles        []roles.Type // Роли пользователя
	Email        string       // Почта
	PasswordHash *string      // Хэш пароля (не обязательно)
}

type UserServiceCreate struct {
	Roles    []roles.Type // Роли пользователя
	Email    *string      // Почта (не обязательно)
	VkId     *int64       // ID в VK (не обязательно)
	Fullname string       // Полное имя
}

type UserUpdate struct {
	Fullname optional.Variable[string] // Полное имя
	Username optional.Variable[string] // Имя пользователя в системе
	VkId     optional.Variable[int64]  // ID в VK
}
