package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"waterfall-backend/internal/models/roles"
	"waterfall-backend/internal/modules/stores/db/ent/privacy"
	"waterfall-backend/internal/modules/stores/db/rules"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "auth_roles"},
	}
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(roles.Type("")).
			Immutable(),
		field.String("description"),
		field.Float("release_delay").SchemaType(map[string]string{
			dialect.Postgres: "bigint",
		}),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("roles"),
		edge.To("permissions", Permission.Type).StorageKey(
			edge.Table("auth_roles_permissions"), edge.Columns("role_id", "permission_id"),
		),
	}
}

func (Role) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rules.AllowIfSystem(),
			rules.AllowIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
	}
}
