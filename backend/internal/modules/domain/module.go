package domain

import (
	"go.uber.org/fx"
)

var (
	ServicesModule = fx.Module("domain",
		fx.Provide(NewServices),
	)
)
