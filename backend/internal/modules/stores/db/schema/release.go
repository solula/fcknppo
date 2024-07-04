package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"waterfall-backend/internal/modules/stores/db/ent/privacy"
	"waterfall-backend/internal/modules/stores/db/rules"
)

// Release holds the schema definition for the Release entity.
type Release struct {
	ent.Schema
}

func (Release) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		TimeMixin{},
	}
}

// Fields of the Release.
func (Release) Fields() []ent.Field {
	return []ent.Field{
		field.Time("release_date"),
		field.String("description"),
	}
}

// Edges of the Release.
func (Release) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("chapters", Chapter.Type),
		edge.To("parts", Part.Type),
	}
}

func (Release) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("release_date"),
	}
}

func (Release) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rules.AllowIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
	}
}
