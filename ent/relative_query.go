// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/afv1/sims-tree/ent/family"
	"github.com/afv1/sims-tree/ent/predicate"
	"github.com/afv1/sims-tree/ent/relative"
)

// RelativeQuery is the builder for querying Relative entities.
type RelativeQuery struct {
	config
	limit        *int
	offset       *int
	unique       *bool
	order        []OrderFunc
	fields       []string
	predicates   []predicate.Relative
	withFamilies *FamilyQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RelativeQuery builder.
func (rq *RelativeQuery) Where(ps ...predicate.Relative) *RelativeQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit adds a limit step to the query.
func (rq *RelativeQuery) Limit(limit int) *RelativeQuery {
	rq.limit = &limit
	return rq
}

// Offset adds an offset step to the query.
func (rq *RelativeQuery) Offset(offset int) *RelativeQuery {
	rq.offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RelativeQuery) Unique(unique bool) *RelativeQuery {
	rq.unique = &unique
	return rq
}

// Order adds an order step to the query.
func (rq *RelativeQuery) Order(o ...OrderFunc) *RelativeQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryFamilies chains the current query on the "families" edge.
func (rq *RelativeQuery) QueryFamilies() *FamilyQuery {
	query := &FamilyQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(relative.Table, relative.FieldID, selector),
			sqlgraph.To(family.Table, family.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, relative.FamiliesTable, relative.FamiliesColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Relative entity from the query.
// Returns a *NotFoundError when no Relative was found.
func (rq *RelativeQuery) First(ctx context.Context) (*Relative, error) {
	nodes, err := rq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{relative.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RelativeQuery) FirstX(ctx context.Context) *Relative {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Relative ID from the query.
// Returns a *NotFoundError when no Relative ID was found.
func (rq *RelativeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{relative.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RelativeQuery) FirstIDX(ctx context.Context) int {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Relative entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Relative entity is found.
// Returns a *NotFoundError when no Relative entities are found.
func (rq *RelativeQuery) Only(ctx context.Context) (*Relative, error) {
	nodes, err := rq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{relative.Label}
	default:
		return nil, &NotSingularError{relative.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RelativeQuery) OnlyX(ctx context.Context) *Relative {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Relative ID in the query.
// Returns a *NotSingularError when more than one Relative ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RelativeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{relative.Label}
	default:
		err = &NotSingularError{relative.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RelativeQuery) OnlyIDX(ctx context.Context) int {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Relatives.
func (rq *RelativeQuery) All(ctx context.Context) ([]*Relative, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rq *RelativeQuery) AllX(ctx context.Context) []*Relative {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Relative IDs.
func (rq *RelativeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rq.Select(relative.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RelativeQuery) IDsX(ctx context.Context) []int {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RelativeQuery) Count(ctx context.Context) (int, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RelativeQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RelativeQuery) Exist(ctx context.Context) (bool, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RelativeQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RelativeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RelativeQuery) Clone() *RelativeQuery {
	if rq == nil {
		return nil
	}
	return &RelativeQuery{
		config:       rq.config,
		limit:        rq.limit,
		offset:       rq.offset,
		order:        append([]OrderFunc{}, rq.order...),
		predicates:   append([]predicate.Relative{}, rq.predicates...),
		withFamilies: rq.withFamilies.Clone(),
		// clone intermediate query.
		sql:    rq.sql.Clone(),
		path:   rq.path,
		unique: rq.unique,
	}
}

// WithFamilies tells the query-builder to eager-load the nodes that are connected to
// the "families" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RelativeQuery) WithFamilies(opts ...func(*FamilyQuery)) *RelativeQuery {
	query := &FamilyQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withFamilies = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Relative.Query().
//		GroupBy(relative.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rq *RelativeQuery) GroupBy(field string, fields ...string) *RelativeGroupBy {
	grbuild := &RelativeGroupBy{config: rq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rq.sqlQuery(ctx), nil
	}
	grbuild.label = relative.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Relative.Query().
//		Select(relative.FieldName).
//		Scan(ctx, &v)
//
func (rq *RelativeQuery) Select(fields ...string) *RelativeSelect {
	rq.fields = append(rq.fields, fields...)
	selbuild := &RelativeSelect{RelativeQuery: rq}
	selbuild.label = relative.Label
	selbuild.flds, selbuild.scan = &rq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a RelativeSelect configured with the given aggregations.
func (rq *RelativeQuery) Aggregate(fns ...AggregateFunc) *RelativeSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RelativeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range rq.fields {
		if !relative.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RelativeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Relative, error) {
	var (
		nodes       = []*Relative{}
		withFKs     = rq.withFKs
		_spec       = rq.querySpec()
		loadedTypes = [1]bool{
			rq.withFamilies != nil,
		}
	)
	if rq.withFamilies != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, relative.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Relative).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Relative{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withFamilies; query != nil {
		if err := rq.loadFamilies(ctx, query, nodes, nil,
			func(n *Relative, e *Family) { n.Edges.Families = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *RelativeQuery) loadFamilies(ctx context.Context, query *FamilyQuery, nodes []*Relative, init func(*Relative), assign func(*Relative, *Family)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Relative)
	for i := range nodes {
		if nodes[i].family_relatives == nil {
			continue
		}
		fk := *nodes[i].family_relatives
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(family.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "family_relatives" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rq *RelativeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.fields
	if len(rq.fields) > 0 {
		_spec.Unique = rq.unique != nil && *rq.unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RelativeQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (rq *RelativeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   relative.Table,
			Columns: relative.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: relative.FieldID,
			},
		},
		From:   rq.sql,
		Unique: true,
	}
	if unique := rq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := rq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, relative.FieldID)
		for i := range fields {
			if fields[i] != relative.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RelativeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(relative.Table)
	columns := rq.fields
	if len(columns) == 0 {
		columns = relative.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.unique != nil && *rq.unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RelativeGroupBy is the group-by builder for Relative entities.
type RelativeGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RelativeGroupBy) Aggregate(fns ...AggregateFunc) *RelativeGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the group-by query and scans the result into the given value.
func (rgb *RelativeGroupBy) Scan(ctx context.Context, v any) error {
	query, err := rgb.path(ctx)
	if err != nil {
		return err
	}
	rgb.sql = query
	return rgb.sqlScan(ctx, v)
}

func (rgb *RelativeGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range rgb.fields {
		if !relative.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rgb *RelativeGroupBy) sqlQuery() *sql.Selector {
	selector := rgb.sql.Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(rgb.fields)+len(rgb.fns))
		for _, f := range rgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(rgb.fields...)...)
}

// RelativeSelect is the builder for selecting fields of Relative entities.
type RelativeSelect struct {
	*RelativeQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RelativeSelect) Aggregate(fns ...AggregateFunc) *RelativeSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RelativeSelect) Scan(ctx context.Context, v any) error {
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	rs.sql = rs.RelativeQuery.sqlQuery(ctx)
	return rs.sqlScan(ctx, v)
}

func (rs *RelativeSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(rs.sql))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		rs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		rs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := rs.sql.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
