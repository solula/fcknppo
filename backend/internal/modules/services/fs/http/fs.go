package http

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	nethttp "net/http"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/models/permissions"
	"waterfall-backend/internal/modules/io/http/api"
	"waterfall-backend/internal/modules/services/fs/dto"
	"waterfall-backend/internal/modules/services/fs/http/models"
	"waterfall-backend/internal/modules/services/fs/service"
	"waterfall-backend/internal/pkg/http/constants"

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

func InvokeFileStorageController(controller *FileStorageController, router api.Router) {
	router.Get("/files/:uuid", controller.GetFileByUuid, permissions.FilesRead)
	router.Post("/files", controller.CreateFile, permissions.FilesCreate)
}

// GetFileByUuid
// @tags Files
// @security ApiKeyAuth
// @in header
// @summary Получить файл по uuid
// @param uuid path string true "uuid файла"
// @success 200 {file} file
// @failure 400 {object} http_errors.ErrorResponse
// @failure 403 {object} http_errors.ErrorResponse
// @failure 404 {object} http_errors.ErrorResponse
// @failure 500 {object} http_errors.ErrorResponse
// @router /api/files/{uuid} [GET]
func (controller *FileStorageController) GetFileByUuid(c echo.Context) error {
	uuid := c.Param("uuid")
	if uuid == "" {
		return err_const.ErrUuidMissing
	}

	dtm, content, err := controller.service.GetFileByUuid(c.Request().Context(), uuid)
	if err != nil {
		return err
	}

	// Устанавливаем необходимый хедер
	c.Response().Header().Set(constants.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%s", dtm.Filename))

	// Устанавливаем тип контента
	mimeType := constants.MIMEOctetStream
	if len(dtm.MIMEType) != 0 {
		mimeType = dtm.MIMEType
	}

	return c.Blob(nethttp.StatusOK, mimeType, content)
}

// CreateFile
// @tags Files
// @security ApiKeyAuth
// @in header
// @summary Создать новый файл
// @param content formData file true "файл"
// @accept multipart/form-data
// @success 200 {object} dto.File
// @failure 400 {object} http_errors.ErrorResponse
// @failure 403 {object} http_errors.ErrorResponse
// @failure 404 {object} http_errors.ErrorResponse
// @failure 500 {object} http_errors.ErrorResponse
// @router /api/files [POST]
func (controller *FileStorageController) CreateFile(c echo.Context) error {
	// Достаем содержимое файла
	file, err := c.FormFile("content")
	if err != nil {
		return err
	}

	metadataString := c.FormValue("metadata")
	if err != nil {
		return err
	}

	var metadata *models.FileMetadata
	err = json.Unmarshal([]byte(metadataString), &metadata)
	if err != nil {
		return err
	}

	newFile := &dto.NewFile{
		Filename:    file.Filename,
		MIMEType:    file.Header.Get(constants.HeaderContentType),
		Description: metadata.Description,
		ObjectRef:   metadata.ObjectRef,
		Type:        metadata.Type,
	}

	// Работа с файлом
	fileReader, err := file.Open()
	if err != nil {
		return fmt.Errorf("не удалось открыть файл")
	}
	defer fileReader.Close()

	createdFile, err := controller.service.CreateTempFile(c.Request().Context(), newFile, fileReader)
	if err != nil {
		return err
	}

	return c.JSON(nethttp.StatusOK, createdFile)
}
