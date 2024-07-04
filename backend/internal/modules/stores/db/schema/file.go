package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"waterfall-backend/internal/models"
	"waterfall-backend/internal/models/files"
	"waterfall-backend/internal/modules/stores/db/ent/privacy"
	"waterfall-backend/internal/modules/stores/db/rules"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

func (File) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		TimeMixin{},
	}
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("filename"),
		field.String("mime_type"),
		field.String("description"),
		field.String("creator_uuid").SchemaType(map[string]string{
			dialect.Postgres: "uuid",
		}).Optional().Nillable(),
		field.String("object_type").GoType(models.ObjectType("")),
		field.String("object_ref").SchemaType(map[string]string{
			dialect.Postgres: "uuid",
		}),
		field.String("type").GoType(files.Type("")),
		field.Bool("temp"),
		field.Uint("sequence_number").Optional().Nillable(),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", User.Type).
			Ref("files").
			Field("creator_uuid").
			Unique(),
	}
}

func (File) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("object_type", "object_ref"),
	}
}

func (File) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rules.AllowIfAdmin(),
			// Проверка доступа при получении реализована в БЛ
		},
		Mutation: privacy.MutationPolicy{
			rules.AllowIfAdmin(),
			rules.FilesFilterByCreatorUuid(),
		},
	}
}
