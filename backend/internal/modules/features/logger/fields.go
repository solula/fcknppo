package logger

import (
	"go.uber.org/zap"
	"waterfall-backend/internal/models/session"
)

func UserFields(ss session.Session) []zap.Field {
	return []zap.Field{
		zap.String("username", ss.Username),
		zap.Stringp("email", ss.Email),
	}
}

// Operation операция, с которой связан лог
const Operation = "operation"
