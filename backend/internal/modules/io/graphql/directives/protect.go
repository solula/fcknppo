package directives

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"go.uber.org/zap"
	"waterfall-backend/internal/models/access"
	logger2 "waterfall-backend/internal/modules/features/logger"
)

func Protect(ctx context.Context, _ interface{}, next graphql.Resolver, protectedFields []string) (res interface{}, err error) {
	// Получаем все запрошенные поля
	requestedFields := graphql.CollectAllFields(ctx)

	protectedFieldsMap := make(map[string]struct{})
	for _, field := range protectedFields {
		protectedFieldsMap[field] = struct{}{}
	}

	// Проверяем наличие в запросе каких-нибудь защищенных полей
	foundProtectedField := false
	for _, requestedField := range requestedFields {
		if _, ok := protectedFieldsMap[requestedField]; ok {
			// Найдено защищенное поле
			foundProtectedField = true
		}
	}

	lg := logger2.GetFromCtx(ctx).With(zap.Strings("fields", protectedFields), zap.String(logger2.Operation, "Protect"))

	// Если запрос содержит только защищенные поля -> устанавливаем необходимость защитить этот запрос
	if foundProtectedField {
		ctx = access.SetProtectionNeeded(ctx)
		lg.Debug("В запросе присутствуют защищенные поля")
	} else {
		lg.Debug("В запросе нет защищенных полей")
	}

	return next(ctx)
}
