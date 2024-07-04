// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"waterfall-backend/internal/modules/stores/db/ent/migrations"
	"waterfall-backend/internal/modules/stores/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MigrationsQuery is the builder for querying Migrations entities.
type MigrationsQuery struct {
	config
	ctx        *QueryContext
	order      []migrations.Order
	inters     []Interceptor
	predicates []predicate.Migrations
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MigrationsQuery builder.
func (mq *MigrationsQuery) Where(ps ...predicate.Migrations) *MigrationsQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit the number of records to be returned by this query.
func (mq *MigrationsQuery) Limit(limit int) *MigrationsQuery {
	mq.ctx.Limit = &limit
	return mq
}

// Offset to start from.
func (mq *MigrationsQuery) Offset(offset int) *MigrationsQuery {
	mq.ctx.Offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MigrationsQuery) Unique(unique bool) *MigrationsQuery {
	mq.ctx.Unique = &unique
	return mq
}

// Order specifies how the records should be ordered.
func (mq *MigrationsQuery) Order(o ...migrations.Order) *MigrationsQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// First returns the first Migrations entity from the query.
// Returns a *NotFoundError when no Migrations was found.
func (mq *MigrationsQuery) First(ctx context.Context) (*Migrations, error) {
	nodes, err := mq.Limit(1).All(setContextOp(ctx, mq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{migrations.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MigrationsQuery) FirstX(ctx context.Context) *Migrations {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Migrations ID from the query.
// Returns a *NotFoundError when no Migrations ID was found.
func (mq *MigrationsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(1).IDs(setContextOp(ctx, mq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{migrations.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MigrationsQuery) FirstIDX(ctx context.Context) int {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Migrations entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Migrations entity is found.
// Returns a *NotFoundError when no Migrations entities are found.
func (mq *MigrationsQuery) Only(ctx context.Context) (*Migrations, error) {
	nodes, err := mq.Limit(2).All(setContextOp(ctx, mq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{migrations.Label}
	default:
		return nil, &NotSingularError{migrations.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MigrationsQuery) OnlyX(ctx context.Context) *Migrations {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Migrations ID in the query.
// Returns a *NotSingularError when more than one Migrations ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MigrationsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(2).IDs(setContextOp(ctx, mq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{migrations.Label}
	default:
		err = &NotSingularError{migrations.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MigrationsQuery) OnlyIDX(ctx context.Context) int {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MigrationsSlice.
func (mq *MigrationsQuery) All(ctx context.Context) ([]*Migrations, error) {
	ctx = setContextOp(ctx, mq.ctx, "All")
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Migrations, *MigrationsQuery]()
	return withInterceptors[[]*Migrations](ctx, mq, qr, mq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mq *MigrationsQuery) AllX(ctx context.Context) []*Migrations {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Migrations IDs.
func (mq *MigrationsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if mq.ctx.Unique == nil && mq.path != nil {
		mq.Unique(true)
	}
	ctx = setContextOp(ctx, mq.ctx, "IDs")
	if err = mq.Select(migrations.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MigrationsQuery) IDsX(ctx context.Context) []int {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MigrationsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mq.ctx, "Count")
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mq, querierCount[*MigrationsQuery](), mq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MigrationsQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MigrationsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mq.ctx, "Exist")
	switch _, err := mq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MigrationsQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MigrationsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MigrationsQuery) Clone() *MigrationsQuery {
	if mq == nil {
		return nil
	}
	return &MigrationsQuery{
		config:     mq.config,
		ctx:        mq.ctx.Clone(),
		order:      append([]migrations.Order{}, mq.order...),
		inters:     append([]Interceptor{}, mq.inters...),
		predicates: append([]predicate.Migrations{}, mq.predicates...),
		// clone intermediate query.
		sql:  mq.sql.Clone(),
		path: mq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Migrated int `json:"migrated,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Migrations.Query().
//		GroupBy(migrations.FieldMigrated).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mq *MigrationsQuery) GroupBy(field string, fields ...string) *MigrationsGroupBy {
	mq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MigrationsGroupBy{build: mq}
	grbuild.flds = &mq.ctx.Fields
	grbuild.label = migrations.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Migrated int `json:"migrated,omitempty"`
//	}
//
//	client.Migrations.Query().
//		Select(migrations.FieldMigrated).
//		Scan(ctx, &v)
func (mq *MigrationsQuery) Select(fields ...string) *MigrationsSelect {
	mq.ctx.Fields = append(mq.ctx.Fields, fields...)
	sbuild := &MigrationsSelect{MigrationsQuery: mq}
	sbuild.label = migrations.Label
	sbuild.flds, sbuild.scan = &mq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MigrationsSelect configured with the given aggregations.
func (mq *MigrationsQuery) Aggregate(fns ...AggregateFunc) *MigrationsSelect {
	return mq.Select().Aggregate(fns...)
}

func (mq *MigrationsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mq); err != nil {
				return err
			}
		}
	}
	for _, f := range mq.ctx.Fields {
		if !migrations.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	if migrations.Policy == nil {
		return errors.New("ent: uninitialized migrations.Policy (forgotten import ent/runtime?)")
	}
	if err := migrations.Policy.EvalQuery(ctx, mq); err != nil {
		return err
	}
	return nil
}

func (mq *MigrationsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Migrations, error) {
	var (
		nodes = []*Migrations{}
		_spec = mq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Migrations).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Migrations{config: mq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(mq.modifiers) > 0 {
		_spec.Modifiers = mq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (mq *MigrationsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	if len(mq.modifiers) > 0 {
		_spec.Modifiers = mq.modifiers
	}
	_spec.Node.Columns = mq.ctx.Fields
	if len(mq.ctx.Fields) > 0 {
		_spec.Unique = mq.ctx.Unique != nil && *mq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MigrationsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(migrations.Table, migrations.Columns, sqlgraph.NewFieldSpec(migrations.FieldID, field.TypeInt))
	_spec.From = mq.sql
	if unique := mq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mq.path != nil {
		_spec.Unique = true
	}
	if fields := mq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, migrations.FieldID)
		for i := range fields {
			if fields[i] != migrations.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MigrationsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(migrations.Table)
	columns := mq.ctx.Fields
	if len(columns) == 0 {
		columns = migrations.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.ctx.Unique != nil && *mq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range mq.modifiers {
		m(selector)
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mq *MigrationsQuery) Modify(modifiers ...func(s *sql.Selector)) *MigrationsSelect {
	mq.modifiers = append(mq.modifiers, modifiers...)
	return mq.Select()
}

// MigrationsGroupBy is the group-by builder for Migrations entities.
type MigrationsGroupBy struct {
	selector
	build *MigrationsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MigrationsGroupBy) Aggregate(fns ...AggregateFunc) *MigrationsGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the selector query and scans the result into the given value.
func (mgb *MigrationsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mgb.build.ctx, "GroupBy")
	if err := mgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MigrationsQuery, *MigrationsGroupBy](ctx, mgb.build, mgb, mgb.build.inters, v)
}

func (mgb *MigrationsGroupBy) sqlScan(ctx context.Context, root *MigrationsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mgb.flds)+len(mgb.fns))
		for _, f := range *mgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MigrationsSelect is the builder for selecting fields of Migrations entities.
type MigrationsSelect struct {
	*MigrationsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ms *MigrationsSelect) Aggregate(fns ...AggregateFunc) *MigrationsSelect {
	ms.fns = append(ms.fns, fns...)
	return ms
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MigrationsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ms.ctx, "Select")
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MigrationsQuery, *MigrationsSelect](ctx, ms.MigrationsQuery, ms, ms.inters, v)
}

func (ms *MigrationsSelect) sqlScan(ctx context.Context, root *MigrationsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ms.fns))
	for _, fn := range ms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ms *MigrationsSelect) Modify(modifiers ...func(s *sql.Selector)) *MigrationsSelect {
	ms.modifiers = append(ms.modifiers, modifiers...)
	return ms
}