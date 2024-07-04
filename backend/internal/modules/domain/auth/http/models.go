package http

import "waterfall-backend/internal/modules/domain/auth/dto"

type GoogleAuthCredentials struct {
	IdToken string `json:"credential"` // Токен доступа Google
}

type VKAuthCredentials struct {
	SilentToken string `json:"token"` // Токен для обменя на токен доступа
	Uuid        string `json:"uuid"`  // Uuid токена
}

func (c *VKAuthCredentials) toDTO() *dto.VKAuthCredentials {
	return &dto.VKAuthCredentials{
		SilentToken: c.SilentToken,
		Uuid:        c.Uuid,
	}
}
