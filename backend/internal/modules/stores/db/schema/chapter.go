package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/modules/stores/db/ent/chapter"
	"waterfall-backend/internal/modules/stores/db/ent/privacy"
	"waterfall-backend/internal/modules/stores/db/rules"
)

// Chapter holds the schema definition for the Chapter entity.
type Chapter struct {
	ent.Schema
}

func (Chapter) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		TimeMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the Chapter.
func (Chapter) Fields() []ent.Field {
	return []ent.Field{
		field.Int("number").Optional().Validate(func(v int) error {
			if v < 1 {
				return err_const.ErrInvalidNumber
			}
			return nil
		}),
		field.String("title"),
		field.String("part_uuid").SchemaType(map[string]string{
			dialect.Postgres: "uuid",
		}),
		field.String("release_uuid").SchemaType(map[string]string{
			dialect.Postgres: "uuid",
		}).Optional().Nillable(),
	}
}

// Edges of the Chapter.
func (Chapter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("part", Part.Type).
			Ref("chapters").
			Field("part_uuid").
			Unique().
			Required(),
		edge.From("release", Release.Type).
			Ref("chapters").
			Field("release_uuid").
			Unique(),
		edge.To("comments", Comment.Type),
		edge.To("chapter_text", ChapterText.Type).
			Unique(),
	}
}

func (Chapter) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("number", "part_uuid").Unique(),
		index.Edges("part"),
		index.Edges("release"),
	}
}

func (Chapter) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rules.AllowIfAdmin(),
			rules.ChaptersHasRelease(),
			rules.ChaptersFilterByReleaseDateOrSetSelection(
				chapter.FieldID,
				chapter.FieldNumber,
				chapter.FieldTitle,
				chapter.FieldReleaseUUID,
			),
		},
		Mutation: privacy.MutationPolicy{
			rules.AllowIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
	}
}
