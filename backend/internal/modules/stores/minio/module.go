package minio

import (
	"go.uber.org/fx"
)

var Module = fx.Module("minio",
	fx.Provide(NewClient))
