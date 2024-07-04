package dto

type VKAuthCredentials struct {
	SilentToken string // Токен для обменя на токен доступа
	Uuid        string // Uuid токена
}
