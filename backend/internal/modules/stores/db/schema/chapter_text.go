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

// ChapterText holds the schema definition for the ChapterText entity.
type ChapterText struct {
	ent.Schema
}

func (ChapterText) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		TimeMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the ChapterText.
func (ChapterText) Fields() []ent.Field {
	return []ent.Field{
		field.String("chapter_uuid").SchemaType(map[string]string{
			dialect.Postgres: "uuid",
		}),
		field.String("text"),
	}
}

// Edges of the ChapterText.
func (ChapterText) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chapter", Chapter.Type).
			Ref("chapter_text").
			Field("chapter_uuid").
			Unique().
			Required(),
	}
}

func (ChapterText) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("chapter"),
	}
}

func (ChapterText) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rules.AllowIfAdmin(),
			rules.ChapterTextsHasRelease(),
			rules.ChapterTextsFilterByReleaseDate(),
		},
		Mutation: privacy.MutationPolicy{
			rules.AllowIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
	}
}
