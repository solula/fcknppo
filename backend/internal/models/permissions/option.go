package permissions

type Ident struct{}

func (t Type) Ident() interface{} {
	return Ident{}
}

func (t Type) Value() interface{} {
	return t
}
