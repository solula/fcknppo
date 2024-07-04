package email

import (
	"go.uber.org/fx"
	"waterfall-backend/internal/modules/services/email/service"
)

var (
	ServiceModule = fx.Module("email-service",
		service.Module,
	)
)
