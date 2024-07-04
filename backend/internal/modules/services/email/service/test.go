package service

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.uber.org/fx"
	"waterfall-backend/internal/modules/services/email/dto"
)

func InvokeEmailService(
	service *EmailService,
	lifecycle fx.Lifecycle,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := service.SendVerificationEmail(ctx, "yura.solopov@gmail.com", &dto.Verification{
					VerificationToken: "token",
				}); err != nil {
					log.Errorf("Ошибка отправки письма: %v", err)
					return
				}
				log.Info("Письмо отправлено")
			}()
			return nil
		},
	})
}
