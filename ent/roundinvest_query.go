// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
	"github.com/kebin6/wolflamp-rpc/ent/round"
	"github.com/kebin6/wolflamp-rpc/ent/roundinvest"
)

// RoundInvestQuery is the builder for querying RoundInvest entities.
type RoundInvestQuery struct {
	config
	ctx        *QueryContext
	order      []roundinvest.OrderOption
	inters     []Interceptor
	predicates []predicate.RoundInvest
	withRound  *RoundQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RoundInvestQuery builder.
func (riq *RoundInvestQuery) Where(ps ...predicate.RoundInvest) *RoundInvestQuery {
	riq.predicates = append(riq.predicates, ps...)
	return riq
}

// Limit the number of records to be returned by this query.
func (riq *RoundInvestQuery) Limit(limit int) *RoundInvestQuery {
	riq.ctx.Limit = &limit
	return riq
}

// Offset to start from.
func (riq *RoundInvestQuery) Offset(offset int) *RoundInvestQuery {
	riq.ctx.Offset = &offset
	return riq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (riq *RoundInvestQuery) Unique(unique bool) *RoundInvestQuery {
	riq.ctx.Unique = &unique
	return riq
}

// Order specifies how the records should be ordered.
func (riq *RoundInvestQuery) Order(o ...roundinvest.OrderOption) *RoundInvestQuery {
	riq.order = append(riq.order, o...)
	return riq
}

// QueryRound chains the current query on the "round" edge.
func (riq *RoundInvestQuery) QueryRound() *RoundQuery {
	query := (&RoundClient{config: riq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := riq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := riq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(roundinvest.Table, roundinvest.FieldID, selector),
			sqlgraph.To(round.Table, round.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, roundinvest.RoundTable, roundinvest.RoundColumn),
		)
		fromU = sqlgraph.SetNeighbors(riq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RoundInvest entity from the query.
// Returns a *NotFoundError when no RoundInvest was found.
func (riq *RoundInvestQuery) First(ctx context.Context) (*RoundInvest, error) {
	nodes, err := riq.Limit(1).All(setContextOp(ctx, riq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{roundinvest.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (riq *RoundInvestQuery) FirstX(ctx context.Context) *RoundInvest {
	node, err := riq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RoundInvest ID from the query.
// Returns a *NotFoundError when no RoundInvest ID was found.
func (riq *RoundInvestQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = riq.Limit(1).IDs(setContextOp(ctx, riq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{roundinvest.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (riq *RoundInvestQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := riq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RoundInvest entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RoundInvest entity is found.
// Returns a *NotFoundError when no RoundInvest entities are found.
func (riq *RoundInvestQuery) Only(ctx context.Context) (*RoundInvest, error) {
	nodes, err := riq.Limit(2).All(setContextOp(ctx, riq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{roundinvest.Label}
	default:
		return nil, &NotSingularError{roundinvest.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (riq *RoundInvestQuery) OnlyX(ctx context.Context) *RoundInvest {
	node, err := riq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RoundInvest ID in the query.
// Returns a *NotSingularError when more than one RoundInvest ID is found.
// Returns a *NotFoundError when no entities are found.
func (riq *RoundInvestQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = riq.Limit(2).IDs(setContextOp(ctx, riq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{roundinvest.Label}
	default:
		err = &NotSingularError{roundinvest.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (riq *RoundInvestQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := riq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RoundInvests.
func (riq *RoundInvestQuery) All(ctx context.Context) ([]*RoundInvest, error) {
	ctx = setContextOp(ctx, riq.ctx, "All")
	if err := riq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RoundInvest, *RoundInvestQuery]()
	return withInterceptors[[]*RoundInvest](ctx, riq, qr, riq.inters)
}

// AllX is like All, but panics if an error occurs.
func (riq *RoundInvestQuery) AllX(ctx context.Context) []*RoundInvest {
	nodes, err := riq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RoundInvest IDs.
func (riq *RoundInvestQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if riq.ctx.Unique == nil && riq.path != nil {
		riq.Unique(true)
	}
	ctx = setContextOp(ctx, riq.ctx, "IDs")
	if err = riq.Select(roundinvest.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (riq *RoundInvestQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := riq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (riq *RoundInvestQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, riq.ctx, "Count")
	if err := riq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, riq, querierCount[*RoundInvestQuery](), riq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (riq *RoundInvestQuery) CountX(ctx context.Context) int {
	count, err := riq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (riq *RoundInvestQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, riq.ctx, "Exist")
	switch _, err := riq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (riq *RoundInvestQuery) ExistX(ctx context.Context) bool {
	exist, err := riq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RoundInvestQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (riq *RoundInvestQuery) Clone() *RoundInvestQuery {
	if riq == nil {
		return nil
	}
	return &RoundInvestQuery{
		config:     riq.config,
		ctx:        riq.ctx.Clone(),
		order:      append([]roundinvest.OrderOption{}, riq.order...),
		inters:     append([]Interceptor{}, riq.inters...),
		predicates: append([]predicate.RoundInvest{}, riq.predicates...),
		withRound:  riq.withRound.Clone(),
		// clone intermediate query.
		sql:  riq.sql.Clone(),
		path: riq.path,
	}
}

// WithRound tells the query-builder to eager-load the nodes that are connected to
// the "round" edge. The optional arguments are used to configure the query builder of the edge.
func (riq *RoundInvestQuery) WithRound(opts ...func(*RoundQuery)) *RoundInvestQuery {
	query := (&RoundClient{config: riq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	riq.withRound = query
	return riq
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
//	client.RoundInvest.Query().
//		GroupBy(roundinvest.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (riq *RoundInvestQuery) GroupBy(field string, fields ...string) *RoundInvestGroupBy {
	riq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RoundInvestGroupBy{build: riq}
	grbuild.flds = &riq.ctx.Fields
	grbuild.label = roundinvest.Label
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
//	client.RoundInvest.Query().
//		Select(roundinvest.FieldCreatedAt).
//		Scan(ctx, &v)
func (riq *RoundInvestQuery) Select(fields ...string) *RoundInvestSelect {
	riq.ctx.Fields = append(riq.ctx.Fields, fields...)
	sbuild := &RoundInvestSelect{RoundInvestQuery: riq}
	sbuild.label = roundinvest.Label
	sbuild.flds, sbuild.scan = &riq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RoundInvestSelect configured with the given aggregations.
func (riq *RoundInvestQuery) Aggregate(fns ...AggregateFunc) *RoundInvestSelect {
	return riq.Select().Aggregate(fns...)
}

func (riq *RoundInvestQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range riq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, riq); err != nil {
				return err
			}
		}
	}
	for _, f := range riq.ctx.Fields {
		if !roundinvest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if riq.path != nil {
		prev, err := riq.path(ctx)
		if err != nil {
			return err
		}
		riq.sql = prev
	}
	return nil
}

func (riq *RoundInvestQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RoundInvest, error) {
	var (
		nodes       = []*RoundInvest{}
		_spec       = riq.querySpec()
		loadedTypes = [1]bool{
			riq.withRound != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RoundInvest).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RoundInvest{config: riq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, riq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := riq.withRound; query != nil {
		if err := riq.loadRound(ctx, query, nodes, nil,
			func(n *RoundInvest, e *Round) { n.Edges.Round = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (riq *RoundInvestQuery) loadRound(ctx context.Context, query *RoundQuery, nodes []*RoundInvest, init func(*RoundInvest), assign func(*RoundInvest, *Round)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*RoundInvest)
	for i := range nodes {
		fk := nodes[i].RoundID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(round.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "round_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (riq *RoundInvestQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := riq.querySpec()
	_spec.Node.Columns = riq.ctx.Fields
	if len(riq.ctx.Fields) > 0 {
		_spec.Unique = riq.ctx.Unique != nil && *riq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, riq.driver, _spec)
}

func (riq *RoundInvestQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(roundinvest.Table, roundinvest.Columns, sqlgraph.NewFieldSpec(roundinvest.FieldID, field.TypeUint64))
	_spec.From = riq.sql
	if unique := riq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if riq.path != nil {
		_spec.Unique = true
	}
	if fields := riq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, roundinvest.FieldID)
		for i := range fields {
			if fields[i] != roundinvest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if riq.withRound != nil {
			_spec.Node.AddColumnOnce(roundinvest.FieldRoundID)
		}
	}
	if ps := riq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := riq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := riq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := riq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (riq *RoundInvestQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(riq.driver.Dialect())
	t1 := builder.Table(roundinvest.Table)
	columns := riq.ctx.Fields
	if len(columns) == 0 {
		columns = roundinvest.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if riq.sql != nil {
		selector = riq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if riq.ctx.Unique != nil && *riq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range riq.predicates {
		p(selector)
	}
	for _, p := range riq.order {
		p(selector)
	}
	if offset := riq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := riq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RoundInvestGroupBy is the group-by builder for RoundInvest entities.
type RoundInvestGroupBy struct {
	selector
	build *RoundInvestQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rigb *RoundInvestGroupBy) Aggregate(fns ...AggregateFunc) *RoundInvestGroupBy {
	rigb.fns = append(rigb.fns, fns...)
	return rigb
}

// Scan applies the selector query and scans the result into the given value.
func (rigb *RoundInvestGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rigb.build.ctx, "GroupBy")
	if err := rigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoundInvestQuery, *RoundInvestGroupBy](ctx, rigb.build, rigb, rigb.build.inters, v)
}

func (rigb *RoundInvestGroupBy) sqlScan(ctx context.Context, root *RoundInvestQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rigb.fns))
	for _, fn := range rigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rigb.flds)+len(rigb.fns))
		for _, f := range *rigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RoundInvestSelect is the builder for selecting fields of RoundInvest entities.
type RoundInvestSelect struct {
	*RoundInvestQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ris *RoundInvestSelect) Aggregate(fns ...AggregateFunc) *RoundInvestSelect {
	ris.fns = append(ris.fns, fns...)
	return ris
}

// Scan applies the selector query and scans the result into the given value.
func (ris *RoundInvestSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ris.ctx, "Select")
	if err := ris.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoundInvestQuery, *RoundInvestSelect](ctx, ris.RoundInvestQuery, ris, ris.inters, v)
}

func (ris *RoundInvestSelect) sqlScan(ctx context.Context, root *RoundInvestQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ris.fns))
	for _, fn := range ris.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ris.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ris.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
