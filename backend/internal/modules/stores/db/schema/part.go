package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/modules/stores/db/ent/part"
	"waterfall-backend/internal/modules/stores/db/ent/privacy"
	rules2 "waterfall-backend/internal/modules/stores/db/rules"
)

// Part holds the schema definition for the Part entity.
type Part struct {
	ent.Schema
}

func (Part) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		TimeMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the Part.
func (Part) Fields() []ent.Field {
	return []ent.Field{
		field.Int("number").Unique().Optional().Validate(func(v int) error {
			if v < 1 {
				return err_const.ErrInvalidNumber
			}
			return nil
		}),
		field.String("title"),
		field.String("annotation").Optional().Nillable(),
		field.String("release_uuid").SchemaType(map[string]string{
			dialect.Postgres: "uuid",
		}).Optional().Nillable(),
	}
}

// Edges of the Part.
func (Part) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chapters", Chapter.Type),
		edge.From("release", Release.Type).
			Ref("parts").
			Field("release_uuid").
			Unique(),
	}
}

func (Part) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("release"),
	}
}

func (Part) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rules2.AllowIfAdmin(),
			rules2.PartsHasRelease(),
			rules2.PartsFilterByReleaseDateOrSetSelection(
				part.FieldID,
				part.FieldNumber,
				part.FieldTitle,
				part.FieldReleaseUUID,
			),
		},
		Mutation: privacy.MutationPolicy{
			rules2.AllowIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
	}
}
