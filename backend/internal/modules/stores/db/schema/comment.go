package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"waterfall-backend/internal/modules/stores/db/ent/privacy"
	"waterfall-backend/internal/modules/stores/db/rules"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		TimeMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").MaxLen(4096),

		field.String("author_uuid").SchemaType(map[string]string{
			dialect.Postgres: "uuid",
		}),
		field.String("parent_uuid").SchemaType(map[string]string{
			dialect.Postgres: "uuid",
		}).Optional().Nillable(),
		field.String("chapter_uuid").SchemaType(map[string]string{
			dialect.Postgres: "uuid",
		}).Optional().Nillable(),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("comments").
			Field("author_uuid").
			Unique().
			Required(),
		edge.To("children", Comment.Type).
			From("parent").
			Field("parent_uuid").
			Unique(),
		edge.From("chapter", Chapter.Type).
			Ref("comments").
			Field("chapter_uuid").
			Unique(),
	}
}

func (Comment) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("chapter"),
	}
}

func (Comment) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rules.AllowIfAdmin(),
		},
		Mutation: privacy.MutationPolicy{
			rules.AllowIfAdmin(),
			rules.CommentsFilterByAuthorUuid(),
		},
	}
}
