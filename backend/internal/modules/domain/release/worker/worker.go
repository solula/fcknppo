package worker

import (
	"context"
	"waterfall-backend/internal/constants/typenames"
	"waterfall-backend/internal/modules/domain/release/dto"
	"waterfall-backend/internal/modules/domain/release/service"
	"waterfall-backend/internal/modules/io/worker"
	_ "waterfall-backend/internal/pkg/http/error_handler/http_errors"
)

type ReleaseController struct {
	service *service.ReleaseService
}

func NewReleaseController(service *service.ReleaseService) *ReleaseController {
	return &ReleaseController{
		service: service,
	}
}

func InvokeReleaseController(controller *ReleaseController, router worker.Router) {
	router.Handle(typenames.ReleaseCreate, controller.ReleaseCreate)
}

func (controller *ReleaseController) ReleaseCreate(ctx context.Context, release *dto.ReleaseCreate) error {
	_, err := controller.service.Create(ctx, release)

	return err
}
