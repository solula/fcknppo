package http

//go:generate go run -mod=mod github.com/swaggo/swag/cmd/swag@v1.16.2 init --generalInfo ./http/docs.go --ot "json" --parseInternal --propertyStrategy pascalcase --dir ../ --output ../../../docs/swagger --parseDependency 1

// @host localhost:4000
// @title Сервис чтения книги
// @accept json
// @produce json
// @version 1.0
// @tag.name GraphQL
// @tag.description Роут GraphQL
// @tag.name Auth
// @tag.description Роуты аутентификации
// @tag.name Files
// @tag.description Роуты файлов
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
