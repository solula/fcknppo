package graphql

import (
	"go.uber.org/fx"
	"waterfall-backend/internal/modules/io/graphql/resolvers"
)

var Module = fx.Module("graphql",
	fx.Provide(resolvers.NewResolver),
	fx.Invoke(RegisterGraphQL))
