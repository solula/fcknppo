// Code generated by ent, DO NOT EDIT.

package file

import (
	"time"
	"waterfall-backend/internal/models"
	"waterfall-backend/internal/models/files"
	"waterfall-backend/internal/modules/stores/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.File {
	return predicate.File(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.File {
	return predicate.File(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.File {
	return predicate.File(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.File {
	return predicate.File(sql.FieldEQ(FieldUpdatedAt, v))
}

// Filename applies equality check predicate on the "filename" field. It's identical to FilenameEQ.
func Filename(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldFilename, v))
}

// MimeType applies equality check predicate on the "mime_type" field. It's identical to MimeTypeEQ.
func MimeType(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldMimeType, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldDescription, v))
}

// CreatorUUID applies equality check predicate on the "creator_uuid" field. It's identical to CreatorUUIDEQ.
func CreatorUUID(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldCreatorUUID, v))
}

// ObjectType applies equality check predicate on the "object_type" field. It's identical to ObjectTypeEQ.
func ObjectType(v models.ObjectType) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldEQ(FieldObjectType, vc))
}

// ObjectRef applies equality check predicate on the "object_ref" field. It's identical to ObjectRefEQ.
func ObjectRef(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldObjectRef, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v files.Type) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldEQ(FieldType, vc))
}

// Temp applies equality check predicate on the "temp" field. It's identical to TempEQ.
func Temp(v bool) predicate.File {
	return predicate.File(sql.FieldEQ(FieldTemp, v))
}

// SequenceNumber applies equality check predicate on the "sequence_number" field. It's identical to SequenceNumberEQ.
func SequenceNumber(v uint) predicate.File {
	return predicate.File(sql.FieldEQ(FieldSequenceNumber, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.File {
	return predicate.File(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.File {
	return predicate.File(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.File {
	return predicate.File(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.File {
	return predicate.File(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.File {
	return predicate.File(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.File {
	return predicate.File(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.File {
	return predicate.File(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.File {
	return predicate.File(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.File {
	return predicate.File(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.File {
	return predicate.File(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.File {
	return predicate.File(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.File {
	return predicate.File(sql.FieldLTE(FieldUpdatedAt, v))
}

// FilenameEQ applies the EQ predicate on the "filename" field.
func FilenameEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldFilename, v))
}

// FilenameNEQ applies the NEQ predicate on the "filename" field.
func FilenameNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldFilename, v))
}

// FilenameIn applies the In predicate on the "filename" field.
func FilenameIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldFilename, vs...))
}

// FilenameNotIn applies the NotIn predicate on the "filename" field.
func FilenameNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldFilename, vs...))
}

// FilenameGT applies the GT predicate on the "filename" field.
func FilenameGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldFilename, v))
}

// FilenameGTE applies the GTE predicate on the "filename" field.
func FilenameGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldFilename, v))
}

// FilenameLT applies the LT predicate on the "filename" field.
func FilenameLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldFilename, v))
}

// FilenameLTE applies the LTE predicate on the "filename" field.
func FilenameLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldFilename, v))
}

// FilenameContains applies the Contains predicate on the "filename" field.
func FilenameContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldFilename, v))
}

// FilenameHasPrefix applies the HasPrefix predicate on the "filename" field.
func FilenameHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldFilename, v))
}

// FilenameHasSuffix applies the HasSuffix predicate on the "filename" field.
func FilenameHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldFilename, v))
}

// FilenameEqualFold applies the EqualFold predicate on the "filename" field.
func FilenameEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldFilename, v))
}

// FilenameContainsFold applies the ContainsFold predicate on the "filename" field.
func FilenameContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldFilename, v))
}

// MimeTypeEQ applies the EQ predicate on the "mime_type" field.
func MimeTypeEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldMimeType, v))
}

// MimeTypeNEQ applies the NEQ predicate on the "mime_type" field.
func MimeTypeNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldMimeType, v))
}

// MimeTypeIn applies the In predicate on the "mime_type" field.
func MimeTypeIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldMimeType, vs...))
}

// MimeTypeNotIn applies the NotIn predicate on the "mime_type" field.
func MimeTypeNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldMimeType, vs...))
}

// MimeTypeGT applies the GT predicate on the "mime_type" field.
func MimeTypeGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldMimeType, v))
}

// MimeTypeGTE applies the GTE predicate on the "mime_type" field.
func MimeTypeGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldMimeType, v))
}

// MimeTypeLT applies the LT predicate on the "mime_type" field.
func MimeTypeLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldMimeType, v))
}

// MimeTypeLTE applies the LTE predicate on the "mime_type" field.
func MimeTypeLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldMimeType, v))
}

// MimeTypeContains applies the Contains predicate on the "mime_type" field.
func MimeTypeContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldMimeType, v))
}

// MimeTypeHasPrefix applies the HasPrefix predicate on the "mime_type" field.
func MimeTypeHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldMimeType, v))
}

// MimeTypeHasSuffix applies the HasSuffix predicate on the "mime_type" field.
func MimeTypeHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldMimeType, v))
}

// MimeTypeEqualFold applies the EqualFold predicate on the "mime_type" field.
func MimeTypeEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldMimeType, v))
}

// MimeTypeContainsFold applies the ContainsFold predicate on the "mime_type" field.
func MimeTypeContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldMimeType, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldDescription, v))
}

// CreatorUUIDEQ applies the EQ predicate on the "creator_uuid" field.
func CreatorUUIDEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldCreatorUUID, v))
}

// CreatorUUIDNEQ applies the NEQ predicate on the "creator_uuid" field.
func CreatorUUIDNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldCreatorUUID, v))
}

// CreatorUUIDIn applies the In predicate on the "creator_uuid" field.
func CreatorUUIDIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldCreatorUUID, vs...))
}

// CreatorUUIDNotIn applies the NotIn predicate on the "creator_uuid" field.
func CreatorUUIDNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldCreatorUUID, vs...))
}

// CreatorUUIDGT applies the GT predicate on the "creator_uuid" field.
func CreatorUUIDGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldCreatorUUID, v))
}

// CreatorUUIDGTE applies the GTE predicate on the "creator_uuid" field.
func CreatorUUIDGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldCreatorUUID, v))
}

// CreatorUUIDLT applies the LT predicate on the "creator_uuid" field.
func CreatorUUIDLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldCreatorUUID, v))
}

// CreatorUUIDLTE applies the LTE predicate on the "creator_uuid" field.
func CreatorUUIDLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldCreatorUUID, v))
}

// CreatorUUIDContains applies the Contains predicate on the "creator_uuid" field.
func CreatorUUIDContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldCreatorUUID, v))
}

// CreatorUUIDHasPrefix applies the HasPrefix predicate on the "creator_uuid" field.
func CreatorUUIDHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldCreatorUUID, v))
}

// CreatorUUIDHasSuffix applies the HasSuffix predicate on the "creator_uuid" field.
func CreatorUUIDHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldCreatorUUID, v))
}

// CreatorUUIDIsNil applies the IsNil predicate on the "creator_uuid" field.
func CreatorUUIDIsNil() predicate.File {
	return predicate.File(sql.FieldIsNull(FieldCreatorUUID))
}

// CreatorUUIDNotNil applies the NotNil predicate on the "creator_uuid" field.
func CreatorUUIDNotNil() predicate.File {
	return predicate.File(sql.FieldNotNull(FieldCreatorUUID))
}

// CreatorUUIDEqualFold applies the EqualFold predicate on the "creator_uuid" field.
func CreatorUUIDEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldCreatorUUID, v))
}

// CreatorUUIDContainsFold applies the ContainsFold predicate on the "creator_uuid" field.
func CreatorUUIDContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldCreatorUUID, v))
}

// ObjectTypeEQ applies the EQ predicate on the "object_type" field.
func ObjectTypeEQ(v models.ObjectType) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldEQ(FieldObjectType, vc))
}

// ObjectTypeNEQ applies the NEQ predicate on the "object_type" field.
func ObjectTypeNEQ(v models.ObjectType) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldNEQ(FieldObjectType, vc))
}

// ObjectTypeIn applies the In predicate on the "object_type" field.
func ObjectTypeIn(vs ...models.ObjectType) predicate.File {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.File(sql.FieldIn(FieldObjectType, v...))
}

// ObjectTypeNotIn applies the NotIn predicate on the "object_type" field.
func ObjectTypeNotIn(vs ...models.ObjectType) predicate.File {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.File(sql.FieldNotIn(FieldObjectType, v...))
}

// ObjectTypeGT applies the GT predicate on the "object_type" field.
func ObjectTypeGT(v models.ObjectType) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldGT(FieldObjectType, vc))
}

// ObjectTypeGTE applies the GTE predicate on the "object_type" field.
func ObjectTypeGTE(v models.ObjectType) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldGTE(FieldObjectType, vc))
}

// ObjectTypeLT applies the LT predicate on the "object_type" field.
func ObjectTypeLT(v models.ObjectType) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldLT(FieldObjectType, vc))
}

// ObjectTypeLTE applies the LTE predicate on the "object_type" field.
func ObjectTypeLTE(v models.ObjectType) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldLTE(FieldObjectType, vc))
}

// ObjectTypeContains applies the Contains predicate on the "object_type" field.
func ObjectTypeContains(v models.ObjectType) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldContains(FieldObjectType, vc))
}

// ObjectTypeHasPrefix applies the HasPrefix predicate on the "object_type" field.
func ObjectTypeHasPrefix(v models.ObjectType) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldHasPrefix(FieldObjectType, vc))
}

// ObjectTypeHasSuffix applies the HasSuffix predicate on the "object_type" field.
func ObjectTypeHasSuffix(v models.ObjectType) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldHasSuffix(FieldObjectType, vc))
}

// ObjectTypeEqualFold applies the EqualFold predicate on the "object_type" field.
func ObjectTypeEqualFold(v models.ObjectType) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldEqualFold(FieldObjectType, vc))
}

// ObjectTypeContainsFold applies the ContainsFold predicate on the "object_type" field.
func ObjectTypeContainsFold(v models.ObjectType) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldContainsFold(FieldObjectType, vc))
}

// ObjectRefEQ applies the EQ predicate on the "object_ref" field.
func ObjectRefEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldObjectRef, v))
}

// ObjectRefNEQ applies the NEQ predicate on the "object_ref" field.
func ObjectRefNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldObjectRef, v))
}

// ObjectRefIn applies the In predicate on the "object_ref" field.
func ObjectRefIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldObjectRef, vs...))
}

// ObjectRefNotIn applies the NotIn predicate on the "object_ref" field.
func ObjectRefNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldObjectRef, vs...))
}

// ObjectRefGT applies the GT predicate on the "object_ref" field.
func ObjectRefGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldObjectRef, v))
}

// ObjectRefGTE applies the GTE predicate on the "object_ref" field.
func ObjectRefGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldObjectRef, v))
}

// ObjectRefLT applies the LT predicate on the "object_ref" field.
func ObjectRefLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldObjectRef, v))
}

// ObjectRefLTE applies the LTE predicate on the "object_ref" field.
func ObjectRefLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldObjectRef, v))
}

// ObjectRefContains applies the Contains predicate on the "object_ref" field.
func ObjectRefContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldObjectRef, v))
}

// ObjectRefHasPrefix applies the HasPrefix predicate on the "object_ref" field.
func ObjectRefHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldObjectRef, v))
}

// ObjectRefHasSuffix applies the HasSuffix predicate on the "object_ref" field.
func ObjectRefHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldObjectRef, v))
}

// ObjectRefEqualFold applies the EqualFold predicate on the "object_ref" field.
func ObjectRefEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldObjectRef, v))
}

// ObjectRefContainsFold applies the ContainsFold predicate on the "object_ref" field.
func ObjectRefContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldObjectRef, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v files.Type) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v files.Type) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...files.Type) predicate.File {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.File(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...files.Type) predicate.File {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.File(sql.FieldNotIn(FieldType, v...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v files.Type) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldGT(FieldType, vc))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v files.Type) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldGTE(FieldType, vc))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v files.Type) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldLT(FieldType, vc))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v files.Type) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldLTE(FieldType, vc))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v files.Type) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldContains(FieldType, vc))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v files.Type) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldHasPrefix(FieldType, vc))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v files.Type) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldHasSuffix(FieldType, vc))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v files.Type) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldEqualFold(FieldType, vc))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v files.Type) predicate.File {
	vc := string(v)
	return predicate.File(sql.FieldContainsFold(FieldType, vc))
}

// TempEQ applies the EQ predicate on the "temp" field.
func TempEQ(v bool) predicate.File {
	return predicate.File(sql.FieldEQ(FieldTemp, v))
}

// TempNEQ applies the NEQ predicate on the "temp" field.
func TempNEQ(v bool) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldTemp, v))
}

// SequenceNumberEQ applies the EQ predicate on the "sequence_number" field.
func SequenceNumberEQ(v uint) predicate.File {
	return predicate.File(sql.FieldEQ(FieldSequenceNumber, v))
}

// SequenceNumberNEQ applies the NEQ predicate on the "sequence_number" field.
func SequenceNumberNEQ(v uint) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldSequenceNumber, v))
}

// SequenceNumberIn applies the In predicate on the "sequence_number" field.
func SequenceNumberIn(vs ...uint) predicate.File {
	return predicate.File(sql.FieldIn(FieldSequenceNumber, vs...))
}

// SequenceNumberNotIn applies the NotIn predicate on the "sequence_number" field.
func SequenceNumberNotIn(vs ...uint) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldSequenceNumber, vs...))
}

// SequenceNumberGT applies the GT predicate on the "sequence_number" field.
func SequenceNumberGT(v uint) predicate.File {
	return predicate.File(sql.FieldGT(FieldSequenceNumber, v))
}

// SequenceNumberGTE applies the GTE predicate on the "sequence_number" field.
func SequenceNumberGTE(v uint) predicate.File {
	return predicate.File(sql.FieldGTE(FieldSequenceNumber, v))
}

// SequenceNumberLT applies the LT predicate on the "sequence_number" field.
func SequenceNumberLT(v uint) predicate.File {
	return predicate.File(sql.FieldLT(FieldSequenceNumber, v))
}

// SequenceNumberLTE applies the LTE predicate on the "sequence_number" field.
func SequenceNumberLTE(v uint) predicate.File {
	return predicate.File(sql.FieldLTE(FieldSequenceNumber, v))
}

// SequenceNumberIsNil applies the IsNil predicate on the "sequence_number" field.
func SequenceNumberIsNil() predicate.File {
	return predicate.File(sql.FieldIsNull(FieldSequenceNumber))
}

// SequenceNumberNotNil applies the NotNil predicate on the "sequence_number" field.
func SequenceNumberNotNil() predicate.File {
	return predicate.File(sql.FieldNotNull(FieldSequenceNumber))
}

// HasCreator applies the HasEdge predicate on the "creator" edge.
func HasCreator() predicate.File {
	return predicate.File(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatorTable, CreatorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatorWith applies the HasEdge predicate on the "creator" edge with a given conditions (other predicates).
func HasCreatorWith(preds ...predicate.User) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		step := newCreatorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		p(s.Not())
	})
}