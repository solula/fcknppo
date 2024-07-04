// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"waterfall-backend/internal/modules/stores/db/ent/chapter"
	"waterfall-backend/internal/modules/stores/db/ent/chaptertext"
	"waterfall-backend/internal/modules/stores/db/ent/comment"
	"waterfall-backend/internal/modules/stores/db/ent/part"
	"waterfall-backend/internal/modules/stores/db/ent/predicate"
	"waterfall-backend/internal/modules/stores/db/ent/release"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChapterQuery is the builder for querying Chapter entities.
type ChapterQuery struct {
	config
	ctx             *QueryContext
	order           []chapter.Order
	inters          []Interceptor
	predicates      []predicate.Chapter
	withPart        *PartQuery
	withRelease     *ReleaseQuery
	withComments    *CommentQuery
	withChapterText *ChapterTextQuery
	modifiers       []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ChapterQuery builder.
func (cq *ChapterQuery) Where(ps ...predicate.Chapter) *ChapterQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ChapterQuery) Limit(limit int) *ChapterQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *ChapterQuery) Offset(offset int) *ChapterQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ChapterQuery) Unique(unique bool) *ChapterQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ChapterQuery) Order(o ...chapter.Order) *ChapterQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryPart chains the current query on the "part" edge.
func (cq *ChapterQuery) QueryPart() *PartQuery {
	query := (&PartClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(chapter.Table, chapter.FieldID, selector),
			sqlgraph.To(part.Table, part.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, chapter.PartTable, chapter.PartColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRelease chains the current query on the "release" edge.
func (cq *ChapterQuery) QueryRelease() *ReleaseQuery {
	query := (&ReleaseClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(chapter.Table, chapter.FieldID, selector),
			sqlgraph.To(release.Table, release.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, chapter.ReleaseTable, chapter.ReleaseColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryComments chains the current query on the "comments" edge.
func (cq *ChapterQuery) QueryComments() *CommentQuery {
	query := (&CommentClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(chapter.Table, chapter.FieldID, selector),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, chapter.CommentsTable, chapter.CommentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChapterText chains the current query on the "chapter_text" edge.
func (cq *ChapterQuery) QueryChapterText() *ChapterTextQuery {
	query := (&ChapterTextClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(chapter.Table, chapter.FieldID, selector),
			sqlgraph.To(chaptertext.Table, chaptertext.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, chapter.ChapterTextTable, chapter.ChapterTextColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Chapter entity from the query.
// Returns a *NotFoundError when no Chapter was found.
func (cq *ChapterQuery) First(ctx context.Context) (*Chapter, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{chapter.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ChapterQuery) FirstX(ctx context.Context) *Chapter {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Chapter ID from the query.
// Returns a *NotFoundError when no Chapter ID was found.
func (cq *ChapterQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{chapter.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ChapterQuery) FirstIDX(ctx context.Context) string {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Chapter entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Chapter entity is found.
// Returns a *NotFoundError when no Chapter entities are found.
func (cq *ChapterQuery) Only(ctx context.Context) (*Chapter, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{chapter.Label}
	default:
		return nil, &NotSingularError{chapter.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ChapterQuery) OnlyX(ctx context.Context) *Chapter {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Chapter ID in the query.
// Returns a *NotSingularError when more than one Chapter ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ChapterQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{chapter.Label}
	default:
		err = &NotSingularError{chapter.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ChapterQuery) OnlyIDX(ctx context.Context) string {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Chapters.
func (cq *ChapterQuery) All(ctx context.Context) ([]*Chapter, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Chapter, *ChapterQuery]()
	return withInterceptors[[]*Chapter](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ChapterQuery) AllX(ctx context.Context) []*Chapter {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Chapter IDs.
func (cq *ChapterQuery) IDs(ctx context.Context) (ids []string, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(chapter.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ChapterQuery) IDsX(ctx context.Context) []string {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ChapterQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ChapterQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ChapterQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ChapterQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ChapterQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ChapterQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ChapterQuery) Clone() *ChapterQuery {
	if cq == nil {
		return nil
	}
	return &ChapterQuery{
		config:          cq.config,
		ctx:             cq.ctx.Clone(),
		order:           append([]chapter.Order{}, cq.order...),
		inters:          append([]Interceptor{}, cq.inters...),
		predicates:      append([]predicate.Chapter{}, cq.predicates...),
		withPart:        cq.withPart.Clone(),
		withRelease:     cq.withRelease.Clone(),
		withComments:    cq.withComments.Clone(),
		withChapterText: cq.withChapterText.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithPart tells the query-builder to eager-load the nodes that are connected to
// the "part" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ChapterQuery) WithPart(opts ...func(*PartQuery)) *ChapterQuery {
	query := (&PartClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withPart = query
	return cq
}

// WithRelease tells the query-builder to eager-load the nodes that are connected to
// the "release" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ChapterQuery) WithRelease(opts ...func(*ReleaseQuery)) *ChapterQuery {
	query := (&ReleaseClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withRelease = query
	return cq
}

// WithComments tells the query-builder to eager-load the nodes that are connected to
// the "comments" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ChapterQuery) WithComments(opts ...func(*CommentQuery)) *ChapterQuery {
	query := (&CommentClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withComments = query
	return cq
}

// WithChapterText tells the query-builder to eager-load the nodes that are connected to
// the "chapter_text" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ChapterQuery) WithChapterText(opts ...func(*ChapterTextQuery)) *ChapterQuery {
	query := (&ChapterTextClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withChapterText = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Chapter.Query().
//		GroupBy(chapter.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *ChapterQuery) GroupBy(field string, fields ...string) *ChapterGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ChapterGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = chapter.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Chapter.Query().
//		Select(chapter.FieldCreatedAt).
//		Scan(ctx, &v)
func (cq *ChapterQuery) Select(fields ...string) *ChapterSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &ChapterSelect{ChapterQuery: cq}
	sbuild.label = chapter.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ChapterSelect configured with the given aggregations.
func (cq *ChapterQuery) Aggregate(fns ...AggregateFunc) *ChapterSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ChapterQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !chapter.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	if chapter.Policy == nil {
		return errors.New("ent: uninitialized chapter.Policy (forgotten import ent/runtime?)")
	}
	if err := chapter.Policy.EvalQuery(ctx, cq); err != nil {
		return err
	}
	return nil
}

func (cq *ChapterQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Chapter, error) {
	var (
		nodes       = []*Chapter{}
		_spec       = cq.querySpec()
		loadedTypes = [4]bool{
			cq.withPart != nil,
			cq.withRelease != nil,
			cq.withComments != nil,
			cq.withChapterText != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Chapter).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Chapter{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withPart; query != nil {
		if err := cq.loadPart(ctx, query, nodes, nil,
			func(n *Chapter, e *Part) { n.Edges.Part = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withRelease; query != nil {
		if err := cq.loadRelease(ctx, query, nodes, nil,
			func(n *Chapter, e *Release) { n.Edges.Release = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withComments; query != nil {
		if err := cq.loadComments(ctx, query, nodes,
			func(n *Chapter) { n.Edges.Comments = []*Comment{} },
			func(n *Chapter, e *Comment) { n.Edges.Comments = append(n.Edges.Comments, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withChapterText; query != nil {
		if err := cq.loadChapterText(ctx, query, nodes, nil,
			func(n *Chapter, e *ChapterText) { n.Edges.ChapterText = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ChapterQuery) loadPart(ctx context.Context, query *PartQuery, nodes []*Chapter, init func(*Chapter), assign func(*Chapter, *Part)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Chapter)
	for i := range nodes {
		fk := nodes[i].PartUUID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(part.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "part_uuid" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *ChapterQuery) loadRelease(ctx context.Context, query *ReleaseQuery, nodes []*Chapter, init func(*Chapter), assign func(*Chapter, *Release)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Chapter)
	for i := range nodes {
		if nodes[i].ReleaseUUID == nil {
			continue
		}
		fk := *nodes[i].ReleaseUUID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(release.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "release_uuid" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *ChapterQuery) loadComments(ctx context.Context, query *CommentQuery, nodes []*Chapter, init func(*Chapter), assign func(*Chapter, *Comment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Chapter)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Comment(func(s *sql.Selector) {
		s.Where(sql.InValues(chapter.CommentsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ChapterUUID
		if fk == nil {
			return fmt.Errorf(`foreign-key "chapter_uuid" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "chapter_uuid" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *ChapterQuery) loadChapterText(ctx context.Context, query *ChapterTextQuery, nodes []*Chapter, init func(*Chapter), assign func(*Chapter, *ChapterText)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Chapter)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.Where(predicate.ChapterText(func(s *sql.Selector) {
		s.Where(sql.InValues(chapter.ChapterTextColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ChapterUUID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "chapter_uuid" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *ChapterQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ChapterQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(chapter.Table, chapter.Columns, sqlgraph.NewFieldSpec(chapter.FieldID, field.TypeString))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chapter.FieldID)
		for i := range fields {
			if fields[i] != chapter.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if cq.withPart != nil {
			_spec.Node.AddColumnOnce(chapter.FieldPartUUID)
		}
		if cq.withRelease != nil {
			_spec.Node.AddColumnOnce(chapter.FieldReleaseUUID)
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ChapterQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(chapter.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = chapter.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range cq.modifiers {
		m(selector)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cq *ChapterQuery) Modify(modifiers ...func(s *sql.Selector)) *ChapterSelect {
	cq.modifiers = append(cq.modifiers, modifiers...)
	return cq.Select()
}

// ChapterGroupBy is the group-by builder for Chapter entities.
type ChapterGroupBy struct {
	selector
	build *ChapterQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ChapterGroupBy) Aggregate(fns ...AggregateFunc) *ChapterGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ChapterGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChapterQuery, *ChapterGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ChapterGroupBy) sqlScan(ctx context.Context, root *ChapterQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ChapterSelect is the builder for selecting fields of Chapter entities.
type ChapterSelect struct {
	*ChapterQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ChapterSelect) Aggregate(fns ...AggregateFunc) *ChapterSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ChapterSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ChapterQuery, *ChapterSelect](ctx, cs.ChapterQuery, cs, cs.inters, v)
}

func (cs *ChapterSelect) sqlScan(ctx context.Context, root *ChapterQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cs *ChapterSelect) Modify(modifiers ...func(s *sql.Selector)) *ChapterSelect {
	cs.modifiers = append(cs.modifiers, modifiers...)
	return cs
}
