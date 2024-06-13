// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kebin6/wolflamp-rpc/ent/exchange"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
)

// ExchangeQuery is the builder for querying Exchange entities.
type ExchangeQuery struct {
	config
	ctx        *QueryContext
	order      []exchange.OrderOption
	inters     []Interceptor
	predicates []predicate.Exchange
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ExchangeQuery builder.
func (eq *ExchangeQuery) Where(ps ...predicate.Exchange) *ExchangeQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit the number of records to be returned by this query.
func (eq *ExchangeQuery) Limit(limit int) *ExchangeQuery {
	eq.ctx.Limit = &limit
	return eq
}

// Offset to start from.
func (eq *ExchangeQuery) Offset(offset int) *ExchangeQuery {
	eq.ctx.Offset = &offset
	return eq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eq *ExchangeQuery) Unique(unique bool) *ExchangeQuery {
	eq.ctx.Unique = &unique
	return eq
}

// Order specifies how the records should be ordered.
func (eq *ExchangeQuery) Order(o ...exchange.OrderOption) *ExchangeQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// First returns the first Exchange entity from the query.
// Returns a *NotFoundError when no Exchange was found.
func (eq *ExchangeQuery) First(ctx context.Context) (*Exchange, error) {
	nodes, err := eq.Limit(1).All(setContextOp(ctx, eq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{exchange.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *ExchangeQuery) FirstX(ctx context.Context) *Exchange {
	node, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Exchange ID from the query.
// Returns a *NotFoundError when no Exchange ID was found.
func (eq *ExchangeQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = eq.Limit(1).IDs(setContextOp(ctx, eq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{exchange.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eq *ExchangeQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Exchange entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Exchange entity is found.
// Returns a *NotFoundError when no Exchange entities are found.
func (eq *ExchangeQuery) Only(ctx context.Context) (*Exchange, error) {
	nodes, err := eq.Limit(2).All(setContextOp(ctx, eq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{exchange.Label}
	default:
		return nil, &NotSingularError{exchange.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *ExchangeQuery) OnlyX(ctx context.Context) *Exchange {
	node, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Exchange ID in the query.
// Returns a *NotSingularError when more than one Exchange ID is found.
// Returns a *NotFoundError when no entities are found.
func (eq *ExchangeQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = eq.Limit(2).IDs(setContextOp(ctx, eq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{exchange.Label}
	default:
		err = &NotSingularError{exchange.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *ExchangeQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Exchanges.
func (eq *ExchangeQuery) All(ctx context.Context) ([]*Exchange, error) {
	ctx = setContextOp(ctx, eq.ctx, "All")
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Exchange, *ExchangeQuery]()
	return withInterceptors[[]*Exchange](ctx, eq, qr, eq.inters)
}

// AllX is like All, but panics if an error occurs.
func (eq *ExchangeQuery) AllX(ctx context.Context) []*Exchange {
	nodes, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Exchange IDs.
func (eq *ExchangeQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if eq.ctx.Unique == nil && eq.path != nil {
		eq.Unique(true)
	}
	ctx = setContextOp(ctx, eq.ctx, "IDs")
	if err = eq.Select(exchange.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *ExchangeQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *ExchangeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, eq.ctx, "Count")
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, eq, querierCount[*ExchangeQuery](), eq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (eq *ExchangeQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *ExchangeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, eq.ctx, "Exist")
	switch _, err := eq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *ExchangeQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ExchangeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *ExchangeQuery) Clone() *ExchangeQuery {
	if eq == nil {
		return nil
	}
	return &ExchangeQuery{
		config:     eq.config,
		ctx:        eq.ctx.Clone(),
		order:      append([]exchange.OrderOption{}, eq.order...),
		inters:     append([]Interceptor{}, eq.inters...),
		predicates: append([]predicate.Exchange{}, eq.predicates...),
		// clone intermediate query.
		sql:  eq.sql.Clone(),
		path: eq.path,
	}
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
//	client.Exchange.Query().
//		GroupBy(exchange.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eq *ExchangeQuery) GroupBy(field string, fields ...string) *ExchangeGroupBy {
	eq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ExchangeGroupBy{build: eq}
	grbuild.flds = &eq.ctx.Fields
	grbuild.label = exchange.Label
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
//	client.Exchange.Query().
//		Select(exchange.FieldCreatedAt).
//		Scan(ctx, &v)
func (eq *ExchangeQuery) Select(fields ...string) *ExchangeSelect {
	eq.ctx.Fields = append(eq.ctx.Fields, fields...)
	sbuild := &ExchangeSelect{ExchangeQuery: eq}
	sbuild.label = exchange.Label
	sbuild.flds, sbuild.scan = &eq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ExchangeSelect configured with the given aggregations.
func (eq *ExchangeQuery) Aggregate(fns ...AggregateFunc) *ExchangeSelect {
	return eq.Select().Aggregate(fns...)
}

func (eq *ExchangeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range eq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, eq); err != nil {
				return err
			}
		}
	}
	for _, f := range eq.ctx.Fields {
		if !exchange.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *ExchangeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Exchange, error) {
	var (
		nodes = []*Exchange{}
		_spec = eq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Exchange).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Exchange{config: eq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (eq *ExchangeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	_spec.Node.Columns = eq.ctx.Fields
	if len(eq.ctx.Fields) > 0 {
		_spec.Unique = eq.ctx.Unique != nil && *eq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *ExchangeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(exchange.Table, exchange.Columns, sqlgraph.NewFieldSpec(exchange.FieldID, field.TypeUint64))
	_spec.From = eq.sql
	if unique := eq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if eq.path != nil {
		_spec.Unique = true
	}
	if fields := eq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, exchange.FieldID)
		for i := range fields {
			if fields[i] != exchange.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eq *ExchangeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(exchange.Table)
	columns := eq.ctx.Fields
	if len(columns) == 0 {
		columns = exchange.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eq.ctx.Unique != nil && *eq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector)
	}
	if offset := eq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ExchangeGroupBy is the group-by builder for Exchange entities.
type ExchangeGroupBy struct {
	selector
	build *ExchangeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *ExchangeGroupBy) Aggregate(fns ...AggregateFunc) *ExchangeGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the selector query and scans the result into the given value.
func (egb *ExchangeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, egb.build.ctx, "GroupBy")
	if err := egb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExchangeQuery, *ExchangeGroupBy](ctx, egb.build, egb, egb.build.inters, v)
}

func (egb *ExchangeGroupBy) sqlScan(ctx context.Context, root *ExchangeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(egb.fns))
	for _, fn := range egb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*egb.flds)+len(egb.fns))
		for _, f := range *egb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*egb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := egb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ExchangeSelect is the builder for selecting fields of Exchange entities.
type ExchangeSelect struct {
	*ExchangeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (es *ExchangeSelect) Aggregate(fns ...AggregateFunc) *ExchangeSelect {
	es.fns = append(es.fns, fns...)
	return es
}

// Scan applies the selector query and scans the result into the given value.
func (es *ExchangeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, es.ctx, "Select")
	if err := es.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExchangeQuery, *ExchangeSelect](ctx, es.ExchangeQuery, es, es.inters, v)
}

func (es *ExchangeSelect) sqlScan(ctx context.Context, root *ExchangeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(es.fns))
	for _, fn := range es.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*es.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
