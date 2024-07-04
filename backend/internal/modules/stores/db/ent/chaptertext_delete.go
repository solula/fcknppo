// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"waterfall-backend/internal/modules/stores/db/ent/chaptertext"
	"waterfall-backend/internal/modules/stores/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChapterTextDelete is the builder for deleting a ChapterText entity.
type ChapterTextDelete struct {
	config
	hooks    []Hook
	mutation *ChapterTextMutation
}

// Where appends a list predicates to the ChapterTextDelete builder.
func (ctd *ChapterTextDelete) Where(ps ...predicate.ChapterText) *ChapterTextDelete {
	ctd.mutation.Where(ps...)
	return ctd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ctd *ChapterTextDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ChapterTextMutation](ctx, ctd.sqlExec, ctd.mutation, ctd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ctd *ChapterTextDelete) ExecX(ctx context.Context) int {
	n, err := ctd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ctd *ChapterTextDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(chaptertext.Table, sqlgraph.NewFieldSpec(chaptertext.FieldID, field.TypeString))
	if ps := ctd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ctd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ctd.mutation.done = true
	return affected, err
}

// ChapterTextDeleteOne is the builder for deleting a single ChapterText entity.
type ChapterTextDeleteOne struct {
	ctd *ChapterTextDelete
}

// Where appends a list predicates to the ChapterTextDelete builder.
func (ctdo *ChapterTextDeleteOne) Where(ps ...predicate.ChapterText) *ChapterTextDeleteOne {
	ctdo.ctd.mutation.Where(ps...)
	return ctdo
}

// Exec executes the deletion query.
func (ctdo *ChapterTextDeleteOne) Exec(ctx context.Context) error {
	n, err := ctdo.ctd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{chaptertext.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ctdo *ChapterTextDeleteOne) ExecX(ctx context.Context) {
	if err := ctdo.Exec(ctx); err != nil {
		panic(err)
	}
}
