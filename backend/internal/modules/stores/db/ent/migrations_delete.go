// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"waterfall-backend/internal/modules/stores/db/ent/migrations"
	"waterfall-backend/internal/modules/stores/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MigrationsDelete is the builder for deleting a Migrations entity.
type MigrationsDelete struct {
	config
	hooks    []Hook
	mutation *MigrationsMutation
}

// Where appends a list predicates to the MigrationsDelete builder.
func (md *MigrationsDelete) Where(ps ...predicate.Migrations) *MigrationsDelete {
	md.mutation.Where(ps...)
	return md
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (md *MigrationsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, MigrationsMutation](ctx, md.sqlExec, md.mutation, md.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (md *MigrationsDelete) ExecX(ctx context.Context) int {
	n, err := md.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (md *MigrationsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(migrations.Table, sqlgraph.NewFieldSpec(migrations.FieldID, field.TypeInt))
	if ps := md.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, md.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	md.mutation.done = true
	return affected, err
}

// MigrationsDeleteOne is the builder for deleting a single Migrations entity.
type MigrationsDeleteOne struct {
	md *MigrationsDelete
}

// Where appends a list predicates to the MigrationsDelete builder.
func (mdo *MigrationsDeleteOne) Where(ps ...predicate.Migrations) *MigrationsDeleteOne {
	mdo.md.mutation.Where(ps...)
	return mdo
}

// Exec executes the deletion query.
func (mdo *MigrationsDeleteOne) Exec(ctx context.Context) error {
	n, err := mdo.md.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{migrations.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mdo *MigrationsDeleteOne) ExecX(ctx context.Context) {
	if err := mdo.Exec(ctx); err != nil {
		panic(err)
	}
}
