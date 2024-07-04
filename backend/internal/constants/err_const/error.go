package err_const

import "fmt"

// Константные ошибки

// Const тип используемый для константных ошибок, позволяет избегать возможных мутаций значений ошибок.
// Не рекомендуется использовать их для создания ошибок в рамках бизнес-логики.
type Const string

func (e Const) Error() string {
	return string(e)
}

const (
	/* Общие ошибки */

	ErrNotFound         = Const("запись не найдена")
	ErrUniqueConstraint = Const("нарушена уникальность")
	ErrMissingUser      = Const("пользователь не указан")
	ErrValidate         = Const("ошибка валидации")

	/* Ошибки авторизации и аутентификации */

	ErrAuthentication   = Const("ошибка аутентификации")
	ErrMissingToken     = Const("не указан токен доступа")
	ErrInvalidToken     = Const("некорректный токен")
	ErrInvalidEmail     = Const("некорректный email")
	ErrPasswordTooEasy  = Const("пароль слишком простой")
	ErrAccessDenied     = Const("доступ запрещен")
	ErrMissingSession   = Const("сессия отсутствует")
	ErrEmailNotVerified = Const("email не подтвержден")
	ErrUserDeleted      = Const("пользователь удален")

	/* Ошибки Id и Uuid */

	ErrIdMissing    = Const("не указан Id")
	ErrUuidMissing  = Const("не указан Uuid")
	ErrIdValidate   = Const("неверно указан Id")
	ErrUuidValidate = Const("неверно указан Uuid")

	/* Ошибки бизнес логики */

	ErrInvalidNumber     = Const("некорректный номер")
	ErrInvalidIndex      = Const("некорректный индекс")
	ErrInvalidAuthType   = Const("некорректный тип авторизации")
	ErrInvalidRole       = Const("некорректная роль")
	ErrInvalidPermission = Const("некорректное право доступа")
	ErrInvalidObjectType = Const("некорректный тип объекта")
	ErrNegativeScore     = Const("отрицательное количество баллов")
)

type PanicError struct {
	Message string
	Detail  string
}

func FromPanic(panicInstance interface{}) *PanicError {
	// Ошибка, полученная из паники
	var panicMsg string

	switch msg := panicInstance.(type) {
	case string:
		panicMsg = msg
	case error:
		panicMsg = msg.Error()

	default:
		panicMsg = fmt.Sprint(panicInstance)
	}

	return &PanicError{
		Message: "внутрення ошибка системы",
		Detail:  panicMsg,
	}
}

func (r *PanicError) Error() string {
	return r.Message
}
