package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"waterfall-backend/internal/models/permissions"
	"waterfall-backend/internal/modules/stores/db/ent/privacy"
	"waterfall-backend/internal/modules/stores/db/rules"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

func (Permission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "auth_permissions"},
	}
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(permissions.Type("")).
			Immutable(),
		field.String("description"),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).
			Ref("permissions"),
	}
}

func (Permission) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rules.AllowIfSystem(),
			rules.AllowIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
	}
}
