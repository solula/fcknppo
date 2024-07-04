package optional

// Variable необязательная переменная типа t.
// Если IsSet() == false, то значение стоит трактовать как неустановленное.
// Если IsSet() == true, то значение равно Value().
// Если при этом Value() == nil, это означает, что переменная имеет явно установленное нулевое значение.
type Variable[T any] struct {
	value *T
	isSet bool
}

func (v *Variable[T]) Set(value *T) {
	v.value = value
	v.isSet = true
}

func (v *Variable[T]) IsSet() bool {
	return v.isSet
}

func (v *Variable[T]) Value() *T {
	return v.value
}

type StringVariable = Variable[string]

type StringSliceVariable = Variable[[]string]

type IntVariable = Variable[int]
