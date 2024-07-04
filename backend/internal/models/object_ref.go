package models

import (
	"fmt"
	"slices"
	"waterfall-backend/internal/constants/err_const"
)

type ObjectType string

const (
	ObjectTypeChapters ObjectType = "chapters"
	ObjectTypeParts    ObjectType = "parts"
	ObjectTypeUsers    ObjectType = "users"
	ObjectTypeComments ObjectType = "comments"

	ObjectTypePublic ObjectType = "public"
)

const (
	MockUuid = "ffffffff-ffff-ffff-ffff-ffffffffffff"
)

var (
	ObjectRefPublic = &ObjectRef{
		Type: ObjectTypePublic,
		Ref:  MockUuid,
	}
)

func (t ObjectType) Values() []string {
	return []string{
		string(ObjectTypePublic),
		string(ObjectTypeUsers),
		string(ObjectTypeChapters),
		string(ObjectTypeParts),
		string(ObjectTypeComments),
	}
}

func (t ObjectType) Validate() error {
	if !slices.Contains(t.Values(), string(t)) {
		return err_const.ErrInvalidObjectType
	}
	return nil
}

// ObjectRef ссылка на сущность
type ObjectRef struct {
	Type ObjectType // Тип (название) сущности
	Ref  string     // Ссылка на сущность (идентификатор)
}

func (o ObjectRef) String() string {
	return fmt.Sprintf("%s.%s", o.Type, o.Ref)
}
