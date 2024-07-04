package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"waterfall-backend/internal/modules/stores/db/ent/privacy"
	"waterfall-backend/internal/modules/stores/db/rules"
)

// Migrations holds the schema definition for the Migrations entity.
type Migrations struct {
	ent.Schema
}

func (Migrations) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "db_migrations"},
	}
}

// Fields of the Migrations.
func (Migrations) Fields() []ent.Field {
	return []ent.Field{
		field.Int("migrated"),
	}
}

// Edges of the Migrations.
func (Migrations) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Migrations) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rules.AllowIfSystem(),
			rules.AllowIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
	}
}
