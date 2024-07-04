package public

import (
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
	"waterfall-backend/internal/modules/features/config"
	"waterfall-backend/internal/pkg/routers"
)

type Router routers.Router

func InitRouter(cfg config.Config, router routers.Router) Router {
	router.Static("/docs", "./docs")
	router.Static("/favicon.ico", "./docs/favicon.ico")

	mon := asynqmon.New(asynqmon.Options{
		RootPath: "/monitoring/asynq",
		RedisConnOpt: asynq.RedisClientOpt{
			Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
			Password: "",
			DB:       0,
		},
	})

	router.All(mon.RootPath()+"/*", mon)

	return router
}
