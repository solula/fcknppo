package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"fmt"
	"net/mail"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/modules/stores/db/ent/privacy"
	"waterfall-backend/internal/modules/stores/db/rules"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "auth_users"},
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		TimeMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique().Validate(func(email string) error {
			_, err := mail.ParseAddress(email)
			if err != nil {
				return fmt.Errorf("%w: %s", err_const.ErrInvalidEmail, err)
			}
			return nil
		}).Optional().Nillable(),
		field.String("fullname"),
		field.String("username"),
		field.String("password_hash").Sensitive().Optional().Nillable(),
		field.Int64("vk_id").Unique().Optional().Nillable(),
		field.Int("score").Default(0).Validate(func(v int) error {
			if v < 0 {
				return err_const.ErrNegativeScore
			}
			return nil
		}),
		field.Bool("email_verified").Default(false),

		field.Uint("serial_number").Unique().SchemaType(map[string]string{
			dialect.Postgres: "serial",
		}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type).Required().StorageKey(
			edge.Table("auth_users_roles"), edge.Columns("user_uuid", "role_id"),
		),
		edge.To("files", File.Type),
		edge.To("comments", Comment.Type),
	}
}

func (User) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rules.AllowIfSystem(),
			rules.AllowIfAdmin(),
			rules.UsersAllowIfProtectionNotNeeded(),
			rules.UsersFilterBySelfUuid(),
		},
		Mutation: privacy.MutationPolicy{
			rules.AllowIfSystem(),
			rules.AllowIfAdmin(),
			rules.UsersAllowMutateSelf(),
			privacy.AlwaysDenyRule(),
		},
	}
}
