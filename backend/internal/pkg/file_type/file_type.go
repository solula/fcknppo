package file_type

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
	"waterfall-backend/internal/constants/err_const"
)

const (
	ErrWrongFileType = err_const.Const("некорректный тип файла")
)

const (
	TypeJSON     = "json"
	TypePDF      = "pdf"
	TypeXLSX     = "xlsx"
	TypePNG      = "png"
	TypeJPEG     = "jpeg"
	TypeSVG      = "svg"
	TypeMarkdown = "markdown"

	// TypeAll Любой тип файла допустим
	TypeAll = "*"
)

var fileExtensionsMap = map[string]string{
	".json":     TypeJSON,
	".pdf":      TypePDF,
	".xls":      TypeXLSX,
	".xlsx":     TypeXLSX,
	".zip":      TypeXLSX,
	".png":      TypePNG,
	".jpeg":     TypeJPEG,
	".jpg":      TypeJPEG,
	".svg":      TypeSVG,
	".md":       TypeMarkdown,
	".markdown": TypeMarkdown,
}

var fileContentTypesMap = map[string]string{
	"application/json": TypeJSON,
	"application/pdf":  TypePDF,
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": TypeXLSX,
	"application/zip": TypeXLSX,
	"image/png":       TypePNG,
	"image/jpeg":      TypeJPEG,
	"image/svg+xml":   TypeSVG,
	"text/markdown":   TypeMarkdown,
}

// Check проверяет, что тип файла и расширение файла совпадают с разрешенными
func Check(filename, mimeType string, availableTypes ...string) error {
	// Получаем тип файла по mimeType
	fileTypeFromMime, err := fileTypeByMime(mimeType)
	if err != nil {
		return wrap(err, filename)
	}

	// Получаем тип файла по его расширению
	fileTypeFromExtension, err := fileTypeByFilename(filename)
	if err != nil {
		return wrap(err, filename)
	}

	err = compare(fileTypeFromMime, fileTypeFromExtension, availableTypes...)
	if err != nil {
		return wrap(err, filename)
	}

	return nil
}

func fileTypeByMime(mimeType string) (string, error) {
	fileType, ok := fileContentTypesMap[mimeType]
	if !ok {
		return "", fmt.Errorf("неподдерживаемый MIME тип файла %s", mimeType)
	}

	return fileType, nil
}

func fileTypeByFilename(fileName string) (string, error) {
	fileExt := filepath.Ext(fileName)
	if fileExt == "" {
		return "", fmt.Errorf("у файла %s не указано расширение", fileName)
	}

	fileType, ok := fileExtensionsMap[fileExt]
	if !ok {
		return "", fmt.Errorf("неподдерживаемое расширение файла %s", fileExt)
	}

	return fileType, nil
}

func compare(fileTypeFromMime, fileTypeFromFilename string, availableTypes ...string) error {
	if fileTypeFromMime != fileTypeFromFilename {
		return fmt.Errorf("расширение файла %s не совпадает с MIME типом файла %s", fileTypeFromFilename, fileTypeFromMime)
	}

	fileType := fileTypeFromMime
	if len(availableTypes) == 1 && availableTypes[0] == TypeAll {
		return nil
	}

	if !slices.Contains(availableTypes, fileType) {
		return fmt.Errorf("тип %s не разрешен; разрешенные типы файлов: %s", fileType, strings.Join(availableTypes, ", "))
	}

	return nil
}

func wrap(err error, filename string) error {
	return fmt.Errorf("%w %s: %s", ErrWrongFileType, filename, err.Error())
}
