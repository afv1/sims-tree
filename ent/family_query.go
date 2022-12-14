// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/afv1/sims-tree/ent/family"
	"github.com/afv1/sims-tree/ent/predicate"
	"github.com/afv1/sims-tree/ent/relative"
	"github.com/afv1/sims-tree/ent/user"
)

// FamilyQuery is the builder for querying Family entities.
type FamilyQuery struct {
	config
	limit         *int
	offset        *int
	unique        *bool
	order         []OrderFunc
	fields        []string
	predicates    []predicate.Family
	withOwner     *UserQuery
	withRelatives *RelativeQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FamilyQuery builder.
func (fq *FamilyQuery) Where(ps ...predicate.Family) *FamilyQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit adds a limit step to the query.
func (fq *FamilyQuery) Limit(limit int) *FamilyQuery {
	fq.limit = &limit
	return fq
}

// Offset adds an offset step to the query.
func (fq *FamilyQuery) Offset(offset int) *FamilyQuery {
	fq.offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FamilyQuery) Unique(unique bool) *FamilyQuery {
	fq.unique = &unique
	return fq
}

// Order adds an order step to the query.
func (fq *FamilyQuery) Order(o ...OrderFunc) *FamilyQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryOwner chains the current query on the "owner" edge.
func (fq *FamilyQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(family.Table, family.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, family.OwnerTable, family.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRelatives chains the current query on the "relatives" edge.
func (fq *FamilyQuery) QueryRelatives() *RelativeQuery {
	query := &RelativeQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(family.Table, family.FieldID, selector),
			sqlgraph.To(relative.Table, relative.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, family.RelativesTable, family.RelativesColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Family entity from the query.
// Returns a *NotFoundError when no Family was found.
func (fq *FamilyQuery) First(ctx context.Context) (*Family, error) {
	nodes, err := fq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{family.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FamilyQuery) FirstX(ctx context.Context) *Family {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Family ID from the query.
// Returns a *NotFoundError when no Family ID was found.
func (fq *FamilyQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{family.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FamilyQuery) FirstIDX(ctx context.Context) int {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Family entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Family entity is found.
// Returns a *NotFoundError when no Family entities are found.
func (fq *FamilyQuery) Only(ctx context.Context) (*Family, error) {
	nodes, err := fq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{family.Label}
	default:
		return nil, &NotSingularError{family.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FamilyQuery) OnlyX(ctx context.Context) *Family {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Family ID in the query.
// Returns a *NotSingularError when more than one Family ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FamilyQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{family.Label}
	default:
		err = &NotSingularError{family.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FamilyQuery) OnlyIDX(ctx context.Context) int {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Families.
func (fq *FamilyQuery) All(ctx context.Context) ([]*Family, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fq *FamilyQuery) AllX(ctx context.Context) []*Family {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Family IDs.
func (fq *FamilyQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := fq.Select(family.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FamilyQuery) IDsX(ctx context.Context) []int {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FamilyQuery) Count(ctx context.Context) (int, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FamilyQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FamilyQuery) Exist(ctx context.Context) (bool, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FamilyQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FamilyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FamilyQuery) Clone() *FamilyQuery {
	if fq == nil {
		return nil
	}
	return &FamilyQuery{
		config:        fq.config,
		limit:         fq.limit,
		offset:        fq.offset,
		order:         append([]OrderFunc{}, fq.order...),
		predicates:    append([]predicate.Family{}, fq.predicates...),
		withOwner:     fq.withOwner.Clone(),
		withRelatives: fq.withRelatives.Clone(),
		// clone intermediate query.
		sql:    fq.sql.Clone(),
		path:   fq.path,
		unique: fq.unique,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FamilyQuery) WithOwner(opts ...func(*UserQuery)) *FamilyQuery {
	query := &UserQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withOwner = query
	return fq
}

// WithRelatives tells the query-builder to eager-load the nodes that are connected to
// the "relatives" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FamilyQuery) WithRelatives(opts ...func(*RelativeQuery)) *FamilyQuery {
	query := &RelativeQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withRelatives = query
	return fq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Surname string `json:"surname,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Family.Query().
//		GroupBy(family.FieldSurname).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fq *FamilyQuery) GroupBy(field string, fields ...string) *FamilyGroupBy {
	grbuild := &FamilyGroupBy{config: fq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fq.sqlQuery(ctx), nil
	}
	grbuild.label = family.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Surname string `json:"surname,omitempty"`
//	}
//
//	client.Family.Query().
//		Select(family.FieldSurname).
//		Scan(ctx, &v)
//
func (fq *FamilyQuery) Select(fields ...string) *FamilySelect {
	fq.fields = append(fq.fields, fields...)
	selbuild := &FamilySelect{FamilyQuery: fq}
	selbuild.label = family.Label
	selbuild.flds, selbuild.scan = &fq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a FamilySelect configured with the given aggregations.
func (fq *FamilyQuery) Aggregate(fns ...AggregateFunc) *FamilySelect {
	return fq.Select().Aggregate(fns...)
}

func (fq *FamilyQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fq.fields {
		if !family.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FamilyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Family, error) {
	var (
		nodes       = []*Family{}
		withFKs     = fq.withFKs
		_spec       = fq.querySpec()
		loadedTypes = [2]bool{
			fq.withOwner != nil,
			fq.withRelatives != nil,
		}
	)
	if fq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, family.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Family).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Family{config: fq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fq.withOwner; query != nil {
		if err := fq.loadOwner(ctx, query, nodes, nil,
			func(n *Family, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	if query := fq.withRelatives; query != nil {
		if err := fq.loadRelatives(ctx, query, nodes,
			func(n *Family) { n.Edges.Relatives = []*Relative{} },
			func(n *Family, e *Relative) { n.Edges.Relatives = append(n.Edges.Relatives, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fq *FamilyQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*Family, init func(*Family), assign func(*Family, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Family)
	for i := range nodes {
		if nodes[i].user_families == nil {
			continue
		}
		fk := *nodes[i].user_families
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_families" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (fq *FamilyQuery) loadRelatives(ctx context.Context, query *RelativeQuery, nodes []*Family, init func(*Family), assign func(*Family, *Relative)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Family)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Relative(func(s *sql.Selector) {
		s.Where(sql.InValues(family.RelativesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.family_relatives
		if fk == nil {
			return fmt.Errorf(`foreign-key "family_relatives" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "family_relatives" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (fq *FamilyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	_spec.Node.Columns = fq.fields
	if len(fq.fields) > 0 {
		_spec.Unique = fq.unique != nil && *fq.unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FamilyQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := fq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (fq *FamilyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   family.Table,
			Columns: family.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: family.FieldID,
			},
		},
		From:   fq.sql,
		Unique: true,
	}
	if unique := fq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, family.FieldID)
		for i := range fields {
			if fields[i] != family.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FamilyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(family.Table)
	columns := fq.fields
	if len(columns) == 0 {
		columns = family.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.unique != nil && *fq.unique {
		selector.Distinct()
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FamilyGroupBy is the group-by builder for Family entities.
type FamilyGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FamilyGroupBy) Aggregate(fns ...AggregateFunc) *FamilyGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fgb *FamilyGroupBy) Scan(ctx context.Context, v any) error {
	query, err := fgb.path(ctx)
	if err != nil {
		return err
	}
	fgb.sql = query
	return fgb.sqlScan(ctx, v)
}

func (fgb *FamilyGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range fgb.fields {
		if !family.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fgb *FamilyGroupBy) sqlQuery() *sql.Selector {
	selector := fgb.sql.Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fgb.fields)+len(fgb.fns))
		for _, f := range fgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fgb.fields...)...)
}

// FamilySelect is the builder for selecting fields of Family entities.
type FamilySelect struct {
	*FamilyQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fs *FamilySelect) Aggregate(fns ...AggregateFunc) *FamilySelect {
	fs.fns = append(fs.fns, fns...)
	return fs
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FamilySelect) Scan(ctx context.Context, v any) error {
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	fs.sql = fs.FamilyQuery.sqlQuery(ctx)
	return fs.sqlScan(ctx, v)
}

func (fs *FamilySelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(fs.fns))
	for _, fn := range fs.fns {
		aggregation = append(aggregation, fn(fs.sql))
	}
	switch n := len(*fs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		fs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		fs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := fs.sql.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
