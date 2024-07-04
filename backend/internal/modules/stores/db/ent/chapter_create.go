// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"waterfall-backend/internal/modules/stores/db/ent/chapter"
	"waterfall-backend/internal/modules/stores/db/ent/chaptertext"
	"waterfall-backend/internal/modules/stores/db/ent/comment"
	"waterfall-backend/internal/modules/stores/db/ent/part"
	"waterfall-backend/internal/modules/stores/db/ent/release"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChapterCreate is the builder for creating a Chapter entity.
type ChapterCreate struct {
	config
	mutation *ChapterMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cc *ChapterCreate) SetCreatedAt(t time.Time) *ChapterCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *ChapterCreate) SetNillableCreatedAt(t *time.Time) *ChapterCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *ChapterCreate) SetUpdatedAt(t time.Time) *ChapterCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *ChapterCreate) SetNillableUpdatedAt(t *time.Time) *ChapterCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *ChapterCreate) SetDeletedAt(t time.Time) *ChapterCreate {
	cc.mutation.SetDeletedAt(t)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *ChapterCreate) SetNillableDeletedAt(t *time.Time) *ChapterCreate {
	if t != nil {
		cc.SetDeletedAt(*t)
	}
	return cc
}

// SetNumber sets the "number" field.
func (cc *ChapterCreate) SetNumber(i int) *ChapterCreate {
	cc.mutation.SetNumber(i)
	return cc
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (cc *ChapterCreate) SetNillableNumber(i *int) *ChapterCreate {
	if i != nil {
		cc.SetNumber(*i)
	}
	return cc
}

// SetTitle sets the "title" field.
func (cc *ChapterCreate) SetTitle(s string) *ChapterCreate {
	cc.mutation.SetTitle(s)
	return cc
}

// SetPartUUID sets the "part_uuid" field.
func (cc *ChapterCreate) SetPartUUID(s string) *ChapterCreate {
	cc.mutation.SetPartUUID(s)
	return cc
}

// SetReleaseUUID sets the "release_uuid" field.
func (cc *ChapterCreate) SetReleaseUUID(s string) *ChapterCreate {
	cc.mutation.SetReleaseUUID(s)
	return cc
}

// SetNillableReleaseUUID sets the "release_uuid" field if the given value is not nil.
func (cc *ChapterCreate) SetNillableReleaseUUID(s *string) *ChapterCreate {
	if s != nil {
		cc.SetReleaseUUID(*s)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ChapterCreate) SetID(s string) *ChapterCreate {
	cc.mutation.SetID(s)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *ChapterCreate) SetNillableID(s *string) *ChapterCreate {
	if s != nil {
		cc.SetID(*s)
	}
	return cc
}

// SetPartID sets the "part" edge to the Part entity by ID.
func (cc *ChapterCreate) SetPartID(id string) *ChapterCreate {
	cc.mutation.SetPartID(id)
	return cc
}

// SetPart sets the "part" edge to the Part entity.
func (cc *ChapterCreate) SetPart(p *Part) *ChapterCreate {
	return cc.SetPartID(p.ID)
}

// SetReleaseID sets the "release" edge to the Release entity by ID.
func (cc *ChapterCreate) SetReleaseID(id string) *ChapterCreate {
	cc.mutation.SetReleaseID(id)
	return cc
}

// SetNillableReleaseID sets the "release" edge to the Release entity by ID if the given value is not nil.
func (cc *ChapterCreate) SetNillableReleaseID(id *string) *ChapterCreate {
	if id != nil {
		cc = cc.SetReleaseID(*id)
	}
	return cc
}

// SetRelease sets the "release" edge to the Release entity.
func (cc *ChapterCreate) SetRelease(r *Release) *ChapterCreate {
	return cc.SetReleaseID(r.ID)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (cc *ChapterCreate) AddCommentIDs(ids ...string) *ChapterCreate {
	cc.mutation.AddCommentIDs(ids...)
	return cc
}

// AddComments adds the "comments" edges to the Comment entity.
func (cc *ChapterCreate) AddComments(c ...*Comment) *ChapterCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddCommentIDs(ids...)
}

// SetChapterTextID sets the "chapter_text" edge to the ChapterText entity by ID.
func (cc *ChapterCreate) SetChapterTextID(id string) *ChapterCreate {
	cc.mutation.SetChapterTextID(id)
	return cc
}

// SetNillableChapterTextID sets the "chapter_text" edge to the ChapterText entity by ID if the given value is not nil.
func (cc *ChapterCreate) SetNillableChapterTextID(id *string) *ChapterCreate {
	if id != nil {
		cc = cc.SetChapterTextID(*id)
	}
	return cc
}

// SetChapterText sets the "chapter_text" edge to the ChapterText entity.
func (cc *ChapterCreate) SetChapterText(c *ChapterText) *ChapterCreate {
	return cc.SetChapterTextID(c.ID)
}

// Mutation returns the ChapterMutation object of the builder.
func (cc *ChapterCreate) Mutation() *ChapterMutation {
	return cc.mutation
}

// Save creates the Chapter in the database.
func (cc *ChapterCreate) Save(ctx context.Context) (*Chapter, error) {
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*Chapter, ChapterMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChapterCreate) SaveX(ctx context.Context) *Chapter {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ChapterCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ChapterCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ChapterCreate) defaults() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		if chapter.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized chapter.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := chapter.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		if chapter.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized chapter.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := chapter.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		if chapter.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized chapter.DefaultID (forgotten import ent/runtime?)")
		}
		v := chapter.DefaultID()
		cc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *ChapterCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Chapter.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Chapter.updated_at"`)}
	}
	if v, ok := cc.mutation.Number(); ok {
		if err := chapter.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "Chapter.number": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Chapter.title"`)}
	}
	if _, ok := cc.mutation.PartUUID(); !ok {
		return &ValidationError{Name: "part_uuid", err: errors.New(`ent: missing required field "Chapter.part_uuid"`)}
	}
	if _, ok := cc.mutation.PartID(); !ok {
		return &ValidationError{Name: "part", err: errors.New(`ent: missing required edge "Chapter.part"`)}
	}
	return nil
}

func (cc *ChapterCreate) sqlSave(ctx context.Context) (*Chapter, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Chapter.ID type: %T", _spec.ID.Value)
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ChapterCreate) createSpec() (*Chapter, *sqlgraph.CreateSpec) {
	var (
		_node = &Chapter{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(chapter.Table, sqlgraph.NewFieldSpec(chapter.FieldID, field.TypeString))
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(chapter.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(chapter.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.SetField(chapter.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := cc.mutation.Number(); ok {
		_spec.SetField(chapter.FieldNumber, field.TypeInt, value)
		_node.Number = value
	}
	if value, ok := cc.mutation.Title(); ok {
		_spec.SetField(chapter.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if nodes := cc.mutation.PartIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chapter.PartTable,
			Columns: []string{chapter.PartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(part.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PartUUID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ReleaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chapter.ReleaseTable,
			Columns: []string{chapter.ReleaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(release.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ReleaseUUID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chapter.CommentsTable,
			Columns: []string{chapter.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ChapterTextIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   chapter.ChapterTextTable,
			Columns: []string{chapter.ChapterTextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(chaptertext.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Chapter.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ChapterUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (cc *ChapterCreate) OnConflict(opts ...sql.ConflictOption) *ChapterUpsertOne {
	cc.conflict = opts
	return &ChapterUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Chapter.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *ChapterCreate) OnConflictColumns(columns ...string) *ChapterUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &ChapterUpsertOne{
		create: cc,
	}
}

type (
	// ChapterUpsertOne is the builder for "upsert"-ing
	//  one Chapter node.
	ChapterUpsertOne struct {
		create *ChapterCreate
	}

	// ChapterUpsert is the "OnConflict" setter.
	ChapterUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *ChapterUpsert) SetUpdatedAt(v time.Time) *ChapterUpsert {
	u.Set(chapter.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ChapterUpsert) UpdateUpdatedAt() *ChapterUpsert {
	u.SetExcluded(chapter.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ChapterUpsert) SetDeletedAt(v time.Time) *ChapterUpsert {
	u.Set(chapter.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ChapterUpsert) UpdateDeletedAt() *ChapterUpsert {
	u.SetExcluded(chapter.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ChapterUpsert) ClearDeletedAt() *ChapterUpsert {
	u.SetNull(chapter.FieldDeletedAt)
	return u
}

// SetNumber sets the "number" field.
func (u *ChapterUpsert) SetNumber(v int) *ChapterUpsert {
	u.Set(chapter.FieldNumber, v)
	return u
}

// UpdateNumber sets the "number" field to the value that was provided on create.
func (u *ChapterUpsert) UpdateNumber() *ChapterUpsert {
	u.SetExcluded(chapter.FieldNumber)
	return u
}

// AddNumber adds v to the "number" field.
func (u *ChapterUpsert) AddNumber(v int) *ChapterUpsert {
	u.Add(chapter.FieldNumber, v)
	return u
}

// ClearNumber clears the value of the "number" field.
func (u *ChapterUpsert) ClearNumber() *ChapterUpsert {
	u.SetNull(chapter.FieldNumber)
	return u
}

// SetTitle sets the "title" field.
func (u *ChapterUpsert) SetTitle(v string) *ChapterUpsert {
	u.Set(chapter.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ChapterUpsert) UpdateTitle() *ChapterUpsert {
	u.SetExcluded(chapter.FieldTitle)
	return u
}

// SetPartUUID sets the "part_uuid" field.
func (u *ChapterUpsert) SetPartUUID(v string) *ChapterUpsert {
	u.Set(chapter.FieldPartUUID, v)
	return u
}

// UpdatePartUUID sets the "part_uuid" field to the value that was provided on create.
func (u *ChapterUpsert) UpdatePartUUID() *ChapterUpsert {
	u.SetExcluded(chapter.FieldPartUUID)
	return u
}

// SetReleaseUUID sets the "release_uuid" field.
func (u *ChapterUpsert) SetReleaseUUID(v string) *ChapterUpsert {
	u.Set(chapter.FieldReleaseUUID, v)
	return u
}

// UpdateReleaseUUID sets the "release_uuid" field to the value that was provided on create.
func (u *ChapterUpsert) UpdateReleaseUUID() *ChapterUpsert {
	u.SetExcluded(chapter.FieldReleaseUUID)
	return u
}

// ClearReleaseUUID clears the value of the "release_uuid" field.
func (u *ChapterUpsert) ClearReleaseUUID() *ChapterUpsert {
	u.SetNull(chapter.FieldReleaseUUID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Chapter.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(chapter.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ChapterUpsertOne) UpdateNewValues() *ChapterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(chapter.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(chapter.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Chapter.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ChapterUpsertOne) Ignore() *ChapterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ChapterUpsertOne) DoNothing() *ChapterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ChapterCreate.OnConflict
// documentation for more info.
func (u *ChapterUpsertOne) Update(set func(*ChapterUpsert)) *ChapterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ChapterUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ChapterUpsertOne) SetUpdatedAt(v time.Time) *ChapterUpsertOne {
	return u.Update(func(s *ChapterUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ChapterUpsertOne) UpdateUpdatedAt() *ChapterUpsertOne {
	return u.Update(func(s *ChapterUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ChapterUpsertOne) SetDeletedAt(v time.Time) *ChapterUpsertOne {
	return u.Update(func(s *ChapterUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ChapterUpsertOne) UpdateDeletedAt() *ChapterUpsertOne {
	return u.Update(func(s *ChapterUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ChapterUpsertOne) ClearDeletedAt() *ChapterUpsertOne {
	return u.Update(func(s *ChapterUpsert) {
		s.ClearDeletedAt()
	})
}

// SetNumber sets the "number" field.
func (u *ChapterUpsertOne) SetNumber(v int) *ChapterUpsertOne {
	return u.Update(func(s *ChapterUpsert) {
		s.SetNumber(v)
	})
}

// AddNumber adds v to the "number" field.
func (u *ChapterUpsertOne) AddNumber(v int) *ChapterUpsertOne {
	return u.Update(func(s *ChapterUpsert) {
		s.AddNumber(v)
	})
}

// UpdateNumber sets the "number" field to the value that was provided on create.
func (u *ChapterUpsertOne) UpdateNumber() *ChapterUpsertOne {
	return u.Update(func(s *ChapterUpsert) {
		s.UpdateNumber()
	})
}

// ClearNumber clears the value of the "number" field.
func (u *ChapterUpsertOne) ClearNumber() *ChapterUpsertOne {
	return u.Update(func(s *ChapterUpsert) {
		s.ClearNumber()
	})
}

// SetTitle sets the "title" field.
func (u *ChapterUpsertOne) SetTitle(v string) *ChapterUpsertOne {
	return u.Update(func(s *ChapterUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ChapterUpsertOne) UpdateTitle() *ChapterUpsertOne {
	return u.Update(func(s *ChapterUpsert) {
		s.UpdateTitle()
	})
}

// SetPartUUID sets the "part_uuid" field.
func (u *ChapterUpsertOne) SetPartUUID(v string) *ChapterUpsertOne {
	return u.Update(func(s *ChapterUpsert) {
		s.SetPartUUID(v)
	})
}

// UpdatePartUUID sets the "part_uuid" field to the value that was provided on create.
func (u *ChapterUpsertOne) UpdatePartUUID() *ChapterUpsertOne {
	return u.Update(func(s *ChapterUpsert) {
		s.UpdatePartUUID()
	})
}

// SetReleaseUUID sets the "release_uuid" field.
func (u *ChapterUpsertOne) SetReleaseUUID(v string) *ChapterUpsertOne {
	return u.Update(func(s *ChapterUpsert) {
		s.SetReleaseUUID(v)
	})
}

// UpdateReleaseUUID sets the "release_uuid" field to the value that was provided on create.
func (u *ChapterUpsertOne) UpdateReleaseUUID() *ChapterUpsertOne {
	return u.Update(func(s *ChapterUpsert) {
		s.UpdateReleaseUUID()
	})
}

// ClearReleaseUUID clears the value of the "release_uuid" field.
func (u *ChapterUpsertOne) ClearReleaseUUID() *ChapterUpsertOne {
	return u.Update(func(s *ChapterUpsert) {
		s.ClearReleaseUUID()
	})
}

// Exec executes the query.
func (u *ChapterUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ChapterCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ChapterUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ChapterUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ChapterUpsertOne.ID is not supported by MySQL driver. Use ChapterUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ChapterUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ChapterCreateBulk is the builder for creating many Chapter entities in bulk.
type ChapterCreateBulk struct {
	config
	builders []*ChapterCreate
	conflict []sql.ConflictOption
}

// Save creates the Chapter entities in the database.
func (ccb *ChapterCreateBulk) Save(ctx context.Context) ([]*Chapter, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Chapter, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChapterMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ChapterCreateBulk) SaveX(ctx context.Context) []*Chapter {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ChapterCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ChapterCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Chapter.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ChapterUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ccb *ChapterCreateBulk) OnConflict(opts ...sql.ConflictOption) *ChapterUpsertBulk {
	ccb.conflict = opts
	return &ChapterUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Chapter.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *ChapterCreateBulk) OnConflictColumns(columns ...string) *ChapterUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &ChapterUpsertBulk{
		create: ccb,
	}
}

// ChapterUpsertBulk is the builder for "upsert"-ing
// a bulk of Chapter nodes.
type ChapterUpsertBulk struct {
	create *ChapterCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Chapter.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(chapter.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ChapterUpsertBulk) UpdateNewValues() *ChapterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(chapter.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(chapter.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Chapter.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ChapterUpsertBulk) Ignore() *ChapterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ChapterUpsertBulk) DoNothing() *ChapterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ChapterCreateBulk.OnConflict
// documentation for more info.
func (u *ChapterUpsertBulk) Update(set func(*ChapterUpsert)) *ChapterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ChapterUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ChapterUpsertBulk) SetUpdatedAt(v time.Time) *ChapterUpsertBulk {
	return u.Update(func(s *ChapterUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ChapterUpsertBulk) UpdateUpdatedAt() *ChapterUpsertBulk {
	return u.Update(func(s *ChapterUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ChapterUpsertBulk) SetDeletedAt(v time.Time) *ChapterUpsertBulk {
	return u.Update(func(s *ChapterUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ChapterUpsertBulk) UpdateDeletedAt() *ChapterUpsertBulk {
	return u.Update(func(s *ChapterUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ChapterUpsertBulk) ClearDeletedAt() *ChapterUpsertBulk {
	return u.Update(func(s *ChapterUpsert) {
		s.ClearDeletedAt()
	})
}

// SetNumber sets the "number" field.
func (u *ChapterUpsertBulk) SetNumber(v int) *ChapterUpsertBulk {
	return u.Update(func(s *ChapterUpsert) {
		s.SetNumber(v)
	})
}

// AddNumber adds v to the "number" field.
func (u *ChapterUpsertBulk) AddNumber(v int) *ChapterUpsertBulk {
	return u.Update(func(s *ChapterUpsert) {
		s.AddNumber(v)
	})
}

// UpdateNumber sets the "number" field to the value that was provided on create.
func (u *ChapterUpsertBulk) UpdateNumber() *ChapterUpsertBulk {
	return u.Update(func(s *ChapterUpsert) {
		s.UpdateNumber()
	})
}

// ClearNumber clears the value of the "number" field.
func (u *ChapterUpsertBulk) ClearNumber() *ChapterUpsertBulk {
	return u.Update(func(s *ChapterUpsert) {
		s.ClearNumber()
	})
}

// SetTitle sets the "title" field.
func (u *ChapterUpsertBulk) SetTitle(v string) *ChapterUpsertBulk {
	return u.Update(func(s *ChapterUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ChapterUpsertBulk) UpdateTitle() *ChapterUpsertBulk {
	return u.Update(func(s *ChapterUpsert) {
		s.UpdateTitle()
	})
}

// SetPartUUID sets the "part_uuid" field.
func (u *ChapterUpsertBulk) SetPartUUID(v string) *ChapterUpsertBulk {
	return u.Update(func(s *ChapterUpsert) {
		s.SetPartUUID(v)
	})
}

// UpdatePartUUID sets the "part_uuid" field to the value that was provided on create.
func (u *ChapterUpsertBulk) UpdatePartUUID() *ChapterUpsertBulk {
	return u.Update(func(s *ChapterUpsert) {
		s.UpdatePartUUID()
	})
}

// SetReleaseUUID sets the "release_uuid" field.
func (u *ChapterUpsertBulk) SetReleaseUUID(v string) *ChapterUpsertBulk {
	return u.Update(func(s *ChapterUpsert) {
		s.SetReleaseUUID(v)
	})
}

// UpdateReleaseUUID sets the "release_uuid" field to the value that was provided on create.
func (u *ChapterUpsertBulk) UpdateReleaseUUID() *ChapterUpsertBulk {
	return u.Update(func(s *ChapterUpsert) {
		s.UpdateReleaseUUID()
	})
}

// ClearReleaseUUID clears the value of the "release_uuid" field.
func (u *ChapterUpsertBulk) ClearReleaseUUID() *ChapterUpsertBulk {
	return u.Update(func(s *ChapterUpsert) {
		s.ClearReleaseUUID()
	})
}

// Exec executes the query.
func (u *ChapterUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ChapterCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ChapterCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ChapterUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}