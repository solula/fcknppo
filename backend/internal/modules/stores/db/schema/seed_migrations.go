package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"waterfall-backend/internal/modules/stores/db/ent/privacy"
	"waterfall-backend/internal/modules/stores/db/rules"
)

// SeedMigrations holds the schema definition for the SeedMigrations entity.
type SeedMigrations struct {
	ent.Schema
}

func (SeedMigrations) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "db_seed_migrations"},
	}
}

// Fields of the SeedMigrations.
func (SeedMigrations) Fields() []ent.Field {
	return []ent.Field{
		field.Int("migrated"),
	}
}

// Edges of the SeedMigrations.
func (SeedMigrations) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (SeedMigrations) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rules.AllowIfSystem(),
			rules.AllowIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
	}
}
