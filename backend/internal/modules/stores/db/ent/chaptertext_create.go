// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"waterfall-backend/internal/modules/stores/db/ent/chapter"
	"waterfall-backend/internal/modules/stores/db/ent/chaptertext"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChapterTextCreate is the builder for creating a ChapterText entity.
type ChapterTextCreate struct {
	config
	mutation *ChapterTextMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ctc *ChapterTextCreate) SetCreatedAt(t time.Time) *ChapterTextCreate {
	ctc.mutation.SetCreatedAt(t)
	return ctc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ctc *ChapterTextCreate) SetNillableCreatedAt(t *time.Time) *ChapterTextCreate {
	if t != nil {
		ctc.SetCreatedAt(*t)
	}
	return ctc
}

// SetUpdatedAt sets the "updated_at" field.
func (ctc *ChapterTextCreate) SetUpdatedAt(t time.Time) *ChapterTextCreate {
	ctc.mutation.SetUpdatedAt(t)
	return ctc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ctc *ChapterTextCreate) SetNillableUpdatedAt(t *time.Time) *ChapterTextCreate {
	if t != nil {
		ctc.SetUpdatedAt(*t)
	}
	return ctc
}

// SetDeletedAt sets the "deleted_at" field.
func (ctc *ChapterTextCreate) SetDeletedAt(t time.Time) *ChapterTextCreate {
	ctc.mutation.SetDeletedAt(t)
	return ctc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ctc *ChapterTextCreate) SetNillableDeletedAt(t *time.Time) *ChapterTextCreate {
	if t != nil {
		ctc.SetDeletedAt(*t)
	}
	return ctc
}

// SetChapterUUID sets the "chapter_uuid" field.
func (ctc *ChapterTextCreate) SetChapterUUID(s string) *ChapterTextCreate {
	ctc.mutation.SetChapterUUID(s)
	return ctc
}

// SetText sets the "text" field.
func (ctc *ChapterTextCreate) SetText(s string) *ChapterTextCreate {
	ctc.mutation.SetText(s)
	return ctc
}

// SetID sets the "id" field.
func (ctc *ChapterTextCreate) SetID(s string) *ChapterTextCreate {
	ctc.mutation.SetID(s)
	return ctc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ctc *ChapterTextCreate) SetNillableID(s *string) *ChapterTextCreate {
	if s != nil {
		ctc.SetID(*s)
	}
	return ctc
}

// SetChapterID sets the "chapter" edge to the Chapter entity by ID.
func (ctc *ChapterTextCreate) SetChapterID(id string) *ChapterTextCreate {
	ctc.mutation.SetChapterID(id)
	return ctc
}

// SetChapter sets the "chapter" edge to the Chapter entity.
func (ctc *ChapterTextCreate) SetChapter(c *Chapter) *ChapterTextCreate {
	return ctc.SetChapterID(c.ID)
}

// Mutation returns the ChapterTextMutation object of the builder.
func (ctc *ChapterTextCreate) Mutation() *ChapterTextMutation {
	return ctc.mutation
}

// Save creates the ChapterText in the database.
func (ctc *ChapterTextCreate) Save(ctx context.Context) (*ChapterText, error) {
	if err := ctc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*ChapterText, ChapterTextMutation](ctx, ctc.sqlSave, ctc.mutation, ctc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ctc *ChapterTextCreate) SaveX(ctx context.Context) *ChapterText {
	v, err := ctc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctc *ChapterTextCreate) Exec(ctx context.Context) error {
	_, err := ctc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctc *ChapterTextCreate) ExecX(ctx context.Context) {
	if err := ctc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctc *ChapterTextCreate) defaults() error {
	if _, ok := ctc.mutation.CreatedAt(); !ok {
		if chaptertext.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized chaptertext.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := chaptertext.DefaultCreatedAt()
		ctc.mutation.SetCreatedAt(v)
	}
	if _, ok := ctc.mutation.UpdatedAt(); !ok {
		if chaptertext.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized chaptertext.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := chaptertext.DefaultUpdatedAt()
		ctc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ctc.mutation.ID(); !ok {
		if chaptertext.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized chaptertext.DefaultID (forgotten import ent/runtime?)")
		}
		v := chaptertext.DefaultID()
		ctc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ctc *ChapterTextCreate) check() error {
	if _, ok := ctc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ChapterText.created_at"`)}
	}
	if _, ok := ctc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ChapterText.updated_at"`)}
	}
	if _, ok := ctc.mutation.ChapterUUID(); !ok {
		return &ValidationError{Name: "chapter_uuid", err: errors.New(`ent: missing required field "ChapterText.chapter_uuid"`)}
	}
	if _, ok := ctc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "ChapterText.text"`)}
	}
	if _, ok := ctc.mutation.ChapterID(); !ok {
		return &ValidationError{Name: "chapter", err: errors.New(`ent: missing required edge "ChapterText.chapter"`)}
	}
	return nil
}

func (ctc *ChapterTextCreate) sqlSave(ctx context.Context) (*ChapterText, error) {
	if err := ctc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ctc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ctc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected ChapterText.ID type: %T", _spec.ID.Value)
		}
	}
	ctc.mutation.id = &_node.ID
	ctc.mutation.done = true
	return _node, nil
}

func (ctc *ChapterTextCreate) createSpec() (*ChapterText, *sqlgraph.CreateSpec) {
	var (
		_node = &ChapterText{config: ctc.config}
		_spec = sqlgraph.NewCreateSpec(chaptertext.Table, sqlgraph.NewFieldSpec(chaptertext.FieldID, field.TypeString))
	)
	_spec.OnConflict = ctc.conflict
	if id, ok := ctc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ctc.mutation.CreatedAt(); ok {
		_spec.SetField(chaptertext.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ctc.mutation.UpdatedAt(); ok {
		_spec.SetField(chaptertext.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ctc.mutation.DeletedAt(); ok {
		_spec.SetField(chaptertext.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := ctc.mutation.Text(); ok {
		_spec.SetField(chaptertext.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if nodes := ctc.mutation.ChapterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   chaptertext.ChapterTable,
			Columns: []string{chaptertext.ChapterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chapter.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ChapterUUID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ChapterText.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ChapterTextUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ctc *ChapterTextCreate) OnConflict(opts ...sql.ConflictOption) *ChapterTextUpsertOne {
	ctc.conflict = opts
	return &ChapterTextUpsertOne{
		create: ctc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ChapterText.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ctc *ChapterTextCreate) OnConflictColumns(columns ...string) *ChapterTextUpsertOne {
	ctc.conflict = append(ctc.conflict, sql.ConflictColumns(columns...))
	return &ChapterTextUpsertOne{
		create: ctc,
	}
}

type (
	// ChapterTextUpsertOne is the builder for "upsert"-ing
	//  one ChapterText node.
	ChapterTextUpsertOne struct {
		create *ChapterTextCreate
	}

	// ChapterTextUpsert is the "OnConflict" setter.
	ChapterTextUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *ChapterTextUpsert) SetUpdatedAt(v time.Time) *ChapterTextUpsert {
	u.Set(chaptertext.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ChapterTextUpsert) UpdateUpdatedAt() *ChapterTextUpsert {
	u.SetExcluded(chaptertext.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ChapterTextUpsert) SetDeletedAt(v time.Time) *ChapterTextUpsert {
	u.Set(chaptertext.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ChapterTextUpsert) UpdateDeletedAt() *ChapterTextUpsert {
	u.SetExcluded(chaptertext.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ChapterTextUpsert) ClearDeletedAt() *ChapterTextUpsert {
	u.SetNull(chaptertext.FieldDeletedAt)
	return u
}

// SetChapterUUID sets the "chapter_uuid" field.
func (u *ChapterTextUpsert) SetChapterUUID(v string) *ChapterTextUpsert {
	u.Set(chaptertext.FieldChapterUUID, v)
	return u
}

// UpdateChapterUUID sets the "chapter_uuid" field to the value that was provided on create.
func (u *ChapterTextUpsert) UpdateChapterUUID() *ChapterTextUpsert {
	u.SetExcluded(chaptertext.FieldChapterUUID)
	return u
}

// SetText sets the "text" field.
func (u *ChapterTextUpsert) SetText(v string) *ChapterTextUpsert {
	u.Set(chaptertext.FieldText, v)
	return u
}

// UpdateText sets the "text" field to the value that was provided on create.
func (u *ChapterTextUpsert) UpdateText() *ChapterTextUpsert {
	u.SetExcluded(chaptertext.FieldText)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ChapterText.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(chaptertext.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ChapterTextUpsertOne) UpdateNewValues() *ChapterTextUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(chaptertext.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(chaptertext.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ChapterText.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ChapterTextUpsertOne) Ignore() *ChapterTextUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ChapterTextUpsertOne) DoNothing() *ChapterTextUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ChapterTextCreate.OnConflict
// documentation for more info.
func (u *ChapterTextUpsertOne) Update(set func(*ChapterTextUpsert)) *ChapterTextUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ChapterTextUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ChapterTextUpsertOne) SetUpdatedAt(v time.Time) *ChapterTextUpsertOne {
	return u.Update(func(s *ChapterTextUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ChapterTextUpsertOne) UpdateUpdatedAt() *ChapterTextUpsertOne {
	return u.Update(func(s *ChapterTextUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ChapterTextUpsertOne) SetDeletedAt(v time.Time) *ChapterTextUpsertOne {
	return u.Update(func(s *ChapterTextUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ChapterTextUpsertOne) UpdateDeletedAt() *ChapterTextUpsertOne {
	return u.Update(func(s *ChapterTextUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ChapterTextUpsertOne) ClearDeletedAt() *ChapterTextUpsertOne {
	return u.Update(func(s *ChapterTextUpsert) {
		s.ClearDeletedAt()
	})
}

// SetChapterUUID sets the "chapter_uuid" field.
func (u *ChapterTextUpsertOne) SetChapterUUID(v string) *ChapterTextUpsertOne {
	return u.Update(func(s *ChapterTextUpsert) {
		s.SetChapterUUID(v)
	})
}

// UpdateChapterUUID sets the "chapter_uuid" field to the value that was provided on create.
func (u *ChapterTextUpsertOne) UpdateChapterUUID() *ChapterTextUpsertOne {
	return u.Update(func(s *ChapterTextUpsert) {
		s.UpdateChapterUUID()
	})
}

// SetText sets the "text" field.
func (u *ChapterTextUpsertOne) SetText(v string) *ChapterTextUpsertOne {
	return u.Update(func(s *ChapterTextUpsert) {
		s.SetText(v)
	})
}

// UpdateText sets the "text" field to the value that was provided on create.
func (u *ChapterTextUpsertOne) UpdateText() *ChapterTextUpsertOne {
	return u.Update(func(s *ChapterTextUpsert) {
		s.UpdateText()
	})
}

// Exec executes the query.
func (u *ChapterTextUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ChapterTextCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ChapterTextUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ChapterTextUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ChapterTextUpsertOne.ID is not supported by MySQL driver. Use ChapterTextUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ChapterTextUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ChapterTextCreateBulk is the builder for creating many ChapterText entities in bulk.
type ChapterTextCreateBulk struct {
	config
	builders []*ChapterTextCreate
	conflict []sql.ConflictOption
}

// Save creates the ChapterText entities in the database.
func (ctcb *ChapterTextCreateBulk) Save(ctx context.Context) ([]*ChapterText, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ctcb.builders))
	nodes := make([]*ChapterText, len(ctcb.builders))
	mutators := make([]Mutator, len(ctcb.builders))
	for i := range ctcb.builders {
		func(i int, root context.Context) {
			builder := ctcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChapterTextMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ctcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ctcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ctcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ctcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ctcb *ChapterTextCreateBulk) SaveX(ctx context.Context) []*ChapterText {
	v, err := ctcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctcb *ChapterTextCreateBulk) Exec(ctx context.Context) error {
	_, err := ctcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctcb *ChapterTextCreateBulk) ExecX(ctx context.Context) {
	if err := ctcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ChapterText.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ChapterTextUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ctcb *ChapterTextCreateBulk) OnConflict(opts ...sql.ConflictOption) *ChapterTextUpsertBulk {
	ctcb.conflict = opts
	return &ChapterTextUpsertBulk{
		create: ctcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ChapterText.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ctcb *ChapterTextCreateBulk) OnConflictColumns(columns ...string) *ChapterTextUpsertBulk {
	ctcb.conflict = append(ctcb.conflict, sql.ConflictColumns(columns...))
	return &ChapterTextUpsertBulk{
		create: ctcb,
	}
}

// ChapterTextUpsertBulk is the builder for "upsert"-ing
// a bulk of ChapterText nodes.
type ChapterTextUpsertBulk struct {
	create *ChapterTextCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ChapterText.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(chaptertext.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ChapterTextUpsertBulk) UpdateNewValues() *ChapterTextUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(chaptertext.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(chaptertext.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ChapterText.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ChapterTextUpsertBulk) Ignore() *ChapterTextUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ChapterTextUpsertBulk) DoNothing() *ChapterTextUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ChapterTextCreateBulk.OnConflict
// documentation for more info.
func (u *ChapterTextUpsertBulk) Update(set func(*ChapterTextUpsert)) *ChapterTextUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ChapterTextUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ChapterTextUpsertBulk) SetUpdatedAt(v time.Time) *ChapterTextUpsertBulk {
	return u.Update(func(s *ChapterTextUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ChapterTextUpsertBulk) UpdateUpdatedAt() *ChapterTextUpsertBulk {
	return u.Update(func(s *ChapterTextUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ChapterTextUpsertBulk) SetDeletedAt(v time.Time) *ChapterTextUpsertBulk {
	return u.Update(func(s *ChapterTextUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ChapterTextUpsertBulk) UpdateDeletedAt() *ChapterTextUpsertBulk {
	return u.Update(func(s *ChapterTextUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ChapterTextUpsertBulk) ClearDeletedAt() *ChapterTextUpsertBulk {
	return u.Update(func(s *ChapterTextUpsert) {
		s.ClearDeletedAt()
	})
}

// SetChapterUUID sets the "chapter_uuid" field.
func (u *ChapterTextUpsertBulk) SetChapterUUID(v string) *ChapterTextUpsertBulk {
	return u.Update(func(s *ChapterTextUpsert) {
		s.SetChapterUUID(v)
	})
}

// UpdateChapterUUID sets the "chapter_uuid" field to the value that was provided on create.
func (u *ChapterTextUpsertBulk) UpdateChapterUUID() *ChapterTextUpsertBulk {
	return u.Update(func(s *ChapterTextUpsert) {
		s.UpdateChapterUUID()
	})
}

// SetText sets the "text" field.
func (u *ChapterTextUpsertBulk) SetText(v string) *ChapterTextUpsertBulk {
	return u.Update(func(s *ChapterTextUpsert) {
		s.SetText(v)
	})
}

// UpdateText sets the "text" field to the value that was provided on create.
func (u *ChapterTextUpsertBulk) UpdateText() *ChapterTextUpsertBulk {
	return u.Update(func(s *ChapterTextUpsert) {
		s.UpdateText()
	})
}

// Exec executes the query.
func (u *ChapterTextUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ChapterTextCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ChapterTextCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ChapterTextUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}