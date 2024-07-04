package directives

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"go.uber.org/zap"
	"waterfall-backend/internal/models/access"
	"waterfall-backend/internal/modules/features/logger"
)

func IgnoreReleaseDate(ctx context.Context, _ interface{}, next graphql.Resolver, availableFields []string) (res interface{}, err error) {
	// В любом случае добавляем к разрешенным полям "__typename"
	availableFields = append(availableFields, "__typename")

	// Получаем все запрошенные поля
	requestedFields := graphql.CollectAllFields(ctx)

	availableFieldsMap := make(map[string]struct{})
	for _, field := range availableFields {
		availableFieldsMap[field] = struct{}{}
	}

	// Проверяем наличие в запросе каких-нибудь не доступных полей
	foundNotAvailableField := false
	for _, requestedField := range requestedFields {
		if _, ok := availableFieldsMap[requestedField]; !ok {
			// Найдено поле, не помеченное как доступное
			foundNotAvailableField = true
		}
	}

	lg := logger.GetFromCtx(ctx).With(zap.Strings("fields", availableFields), zap.String(logger.Operation, "IgnoreReleaseDate"))

	// Если запрос содержит только доступные поля -> разрешаем игнорировать дату релиза
	if !foundNotAvailableField {
		ctx = access.SetIgnoreReleaseDate(ctx)
		lg.Debug("В запросе присутствуют только доступные поля")
	} else {
		lg.Debug("В запросе присутствует недоступное поле")
	}

	return next(ctx)
}
