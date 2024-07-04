package dto

// UserCredentials учетные данные пользователя (для аутентификации и регистрации)
type UserCredentials struct {
	Email    string // Почта пользователя
	Password string // Пароль
}

// NewExternalUser данные нового пользователя для регистрации (получены из внешнего сервиса)
type NewExternalUser struct {
	Email    string // Почта
	Fullname string // Полное имя
}

// Tokens токены
type Tokens struct {
	AccessToken  string // Токен доступа
	RefreshToken string // Токен обновления токена доступа
}

// JWT ответ на успешную аутентификацию
type JWT struct {
	Tokens
	AccessTokenPayload *AccessTokenPayload // Сессия пользователя
}

type EmailVerification struct {
	Token string // Токен подтверждения почты
}
