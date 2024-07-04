package worker

import (
	"context"
	"waterfall-backend/internal/constants/typenames"
	"waterfall-backend/internal/modules/io/worker"
	"waterfall-backend/internal/modules/services/fs/service"

	_ "waterfall-backend/internal/pkg/http/error_handler/http_errors"
)

type FileStorageController struct {
	service *service.FileStorageService
}

func NewFileStorageController(service *service.FileStorageService) *FileStorageController {
	return &FileStorageController{
		service: service,
	}
}

func InvokeFileStorageController(controller *FileStorageController, router worker.Router) {
	router.Handle(typenames.FileDeleteTemp, controller.DeleteTempFile)
}

func (controller *FileStorageController) DeleteTempFile(ctx context.Context, fileUuid string) error {
	err := controller.service.DeleteTempFile(ctx, fileUuid)
	return err
}
