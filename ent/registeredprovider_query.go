// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/clpombo/search/ent/predicate"
	"github.com/clpombo/search/ent/registeredcontract"
	"github.com/clpombo/search/ent/registeredprovider"
	"github.com/google/uuid"
)

// RegisteredProviderQuery is the builder for querying RegisteredProvider entities.
type RegisteredProviderQuery struct {
	config
	ctx          *QueryContext
	order        []registeredprovider.OrderOption
	inters       []Interceptor
	predicates   []predicate.RegisteredProvider
	withContract *RegisteredContractQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RegisteredProviderQuery builder.
func (rpq *RegisteredProviderQuery) Where(ps ...predicate.RegisteredProvider) *RegisteredProviderQuery {
	rpq.predicates = append(rpq.predicates, ps...)
	return rpq
}

// Limit the number of records to be returned by this query.
func (rpq *RegisteredProviderQuery) Limit(limit int) *RegisteredProviderQuery {
	rpq.ctx.Limit = &limit
	return rpq
}

// Offset to start from.
func (rpq *RegisteredProviderQuery) Offset(offset int) *RegisteredProviderQuery {
	rpq.ctx.Offset = &offset
	return rpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rpq *RegisteredProviderQuery) Unique(unique bool) *RegisteredProviderQuery {
	rpq.ctx.Unique = &unique
	return rpq
}

// Order specifies how the records should be ordered.
func (rpq *RegisteredProviderQuery) Order(o ...registeredprovider.OrderOption) *RegisteredProviderQuery {
	rpq.order = append(rpq.order, o...)
	return rpq
}

// QueryContract chains the current query on the "contract" edge.
func (rpq *RegisteredProviderQuery) QueryContract() *RegisteredContractQuery {
	query := (&RegisteredContractClient{config: rpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(registeredprovider.Table, registeredprovider.FieldID, selector),
			sqlgraph.To(registeredcontract.Table, registeredcontract.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, registeredprovider.ContractTable, registeredprovider.ContractColumn),
		)
		fromU = sqlgraph.SetNeighbors(rpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RegisteredProvider entity from the query.
// Returns a *NotFoundError when no RegisteredProvider was found.
func (rpq *RegisteredProviderQuery) First(ctx context.Context) (*RegisteredProvider, error) {
	nodes, err := rpq.Limit(1).All(setContextOp(ctx, rpq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{registeredprovider.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rpq *RegisteredProviderQuery) FirstX(ctx context.Context) *RegisteredProvider {
	node, err := rpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RegisteredProvider ID from the query.
// Returns a *NotFoundError when no RegisteredProvider ID was found.
func (rpq *RegisteredProviderQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rpq.Limit(1).IDs(setContextOp(ctx, rpq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{registeredprovider.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rpq *RegisteredProviderQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := rpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RegisteredProvider entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RegisteredProvider entity is found.
// Returns a *NotFoundError when no RegisteredProvider entities are found.
func (rpq *RegisteredProviderQuery) Only(ctx context.Context) (*RegisteredProvider, error) {
	nodes, err := rpq.Limit(2).All(setContextOp(ctx, rpq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{registeredprovider.Label}
	default:
		return nil, &NotSingularError{registeredprovider.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rpq *RegisteredProviderQuery) OnlyX(ctx context.Context) *RegisteredProvider {
	node, err := rpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RegisteredProvider ID in the query.
// Returns a *NotSingularError when more than one RegisteredProvider ID is found.
// Returns a *NotFoundError when no entities are found.
func (rpq *RegisteredProviderQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rpq.Limit(2).IDs(setContextOp(ctx, rpq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{registeredprovider.Label}
	default:
		err = &NotSingularError{registeredprovider.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rpq *RegisteredProviderQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := rpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RegisteredProviders.
func (rpq *RegisteredProviderQuery) All(ctx context.Context) ([]*RegisteredProvider, error) {
	ctx = setContextOp(ctx, rpq.ctx, "All")
	if err := rpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RegisteredProvider, *RegisteredProviderQuery]()
	return withInterceptors[[]*RegisteredProvider](ctx, rpq, qr, rpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rpq *RegisteredProviderQuery) AllX(ctx context.Context) []*RegisteredProvider {
	nodes, err := rpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RegisteredProvider IDs.
func (rpq *RegisteredProviderQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if rpq.ctx.Unique == nil && rpq.path != nil {
		rpq.Unique(true)
	}
	ctx = setContextOp(ctx, rpq.ctx, "IDs")
	if err = rpq.Select(registeredprovider.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rpq *RegisteredProviderQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := rpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rpq *RegisteredProviderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rpq.ctx, "Count")
	if err := rpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rpq, querierCount[*RegisteredProviderQuery](), rpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rpq *RegisteredProviderQuery) CountX(ctx context.Context) int {
	count, err := rpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rpq *RegisteredProviderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rpq.ctx, "Exist")
	switch _, err := rpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rpq *RegisteredProviderQuery) ExistX(ctx context.Context) bool {
	exist, err := rpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RegisteredProviderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rpq *RegisteredProviderQuery) Clone() *RegisteredProviderQuery {
	if rpq == nil {
		return nil
	}
	return &RegisteredProviderQuery{
		config:       rpq.config,
		ctx:          rpq.ctx.Clone(),
		order:        append([]registeredprovider.OrderOption{}, rpq.order...),
		inters:       append([]Interceptor{}, rpq.inters...),
		predicates:   append([]predicate.RegisteredProvider{}, rpq.predicates...),
		withContract: rpq.withContract.Clone(),
		// clone intermediate query.
		sql:  rpq.sql.Clone(),
		path: rpq.path,
	}
}

// WithContract tells the query-builder to eager-load the nodes that are connected to
// the "contract" edge. The optional arguments are used to configure the query builder of the edge.
func (rpq *RegisteredProviderQuery) WithContract(opts ...func(*RegisteredContractQuery)) *RegisteredProviderQuery {
	query := (&RegisteredContractClient{config: rpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rpq.withContract = query
	return rpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		URL *url.URL `json:"url,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RegisteredProvider.Query().
//		GroupBy(registeredprovider.FieldURL).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rpq *RegisteredProviderQuery) GroupBy(field string, fields ...string) *RegisteredProviderGroupBy {
	rpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RegisteredProviderGroupBy{build: rpq}
	grbuild.flds = &rpq.ctx.Fields
	grbuild.label = registeredprovider.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		URL *url.URL `json:"url,omitempty"`
//	}
//
//	client.RegisteredProvider.Query().
//		Select(registeredprovider.FieldURL).
//		Scan(ctx, &v)
func (rpq *RegisteredProviderQuery) Select(fields ...string) *RegisteredProviderSelect {
	rpq.ctx.Fields = append(rpq.ctx.Fields, fields...)
	sbuild := &RegisteredProviderSelect{RegisteredProviderQuery: rpq}
	sbuild.label = registeredprovider.Label
	sbuild.flds, sbuild.scan = &rpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RegisteredProviderSelect configured with the given aggregations.
func (rpq *RegisteredProviderQuery) Aggregate(fns ...AggregateFunc) *RegisteredProviderSelect {
	return rpq.Select().Aggregate(fns...)
}

func (rpq *RegisteredProviderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rpq); err != nil {
				return err
			}
		}
	}
	for _, f := range rpq.ctx.Fields {
		if !registeredprovider.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rpq.path != nil {
		prev, err := rpq.path(ctx)
		if err != nil {
			return err
		}
		rpq.sql = prev
	}
	return nil
}

func (rpq *RegisteredProviderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RegisteredProvider, error) {
	var (
		nodes       = []*RegisteredProvider{}
		_spec       = rpq.querySpec()
		loadedTypes = [1]bool{
			rpq.withContract != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RegisteredProvider).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RegisteredProvider{config: rpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rpq.withContract; query != nil {
		if err := rpq.loadContract(ctx, query, nodes, nil,
			func(n *RegisteredProvider, e *RegisteredContract) { n.Edges.Contract = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rpq *RegisteredProviderQuery) loadContract(ctx context.Context, query *RegisteredContractQuery, nodes []*RegisteredProvider, init func(*RegisteredProvider), assign func(*RegisteredProvider, *RegisteredContract)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*RegisteredProvider)
	for i := range nodes {
		fk := nodes[i].ContractID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(registeredcontract.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "contract_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rpq *RegisteredProviderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rpq.querySpec()
	_spec.Node.Columns = rpq.ctx.Fields
	if len(rpq.ctx.Fields) > 0 {
		_spec.Unique = rpq.ctx.Unique != nil && *rpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rpq.driver, _spec)
}

func (rpq *RegisteredProviderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(registeredprovider.Table, registeredprovider.Columns, sqlgraph.NewFieldSpec(registeredprovider.FieldID, field.TypeUUID))
	_spec.From = rpq.sql
	if unique := rpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rpq.path != nil {
		_spec.Unique = true
	}
	if fields := rpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, registeredprovider.FieldID)
		for i := range fields {
			if fields[i] != registeredprovider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if rpq.withContract != nil {
			_spec.Node.AddColumnOnce(registeredprovider.FieldContractID)
		}
	}
	if ps := rpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rpq *RegisteredProviderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rpq.driver.Dialect())
	t1 := builder.Table(registeredprovider.Table)
	columns := rpq.ctx.Fields
	if len(columns) == 0 {
		columns = registeredprovider.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rpq.sql != nil {
		selector = rpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rpq.ctx.Unique != nil && *rpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rpq.predicates {
		p(selector)
	}
	for _, p := range rpq.order {
		p(selector)
	}
	if offset := rpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RegisteredProviderGroupBy is the group-by builder for RegisteredProvider entities.
type RegisteredProviderGroupBy struct {
	selector
	build *RegisteredProviderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rpgb *RegisteredProviderGroupBy) Aggregate(fns ...AggregateFunc) *RegisteredProviderGroupBy {
	rpgb.fns = append(rpgb.fns, fns...)
	return rpgb
}

// Scan applies the selector query and scans the result into the given value.
func (rpgb *RegisteredProviderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rpgb.build.ctx, "GroupBy")
	if err := rpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RegisteredProviderQuery, *RegisteredProviderGroupBy](ctx, rpgb.build, rpgb, rpgb.build.inters, v)
}

func (rpgb *RegisteredProviderGroupBy) sqlScan(ctx context.Context, root *RegisteredProviderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rpgb.fns))
	for _, fn := range rpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rpgb.flds)+len(rpgb.fns))
		for _, f := range *rpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RegisteredProviderSelect is the builder for selecting fields of RegisteredProvider entities.
type RegisteredProviderSelect struct {
	*RegisteredProviderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rps *RegisteredProviderSelect) Aggregate(fns ...AggregateFunc) *RegisteredProviderSelect {
	rps.fns = append(rps.fns, fns...)
	return rps
}

// Scan applies the selector query and scans the result into the given value.
func (rps *RegisteredProviderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rps.ctx, "Select")
	if err := rps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RegisteredProviderQuery, *RegisteredProviderSelect](ctx, rps.RegisteredProviderQuery, rps, rps.inters, v)
}

func (rps *RegisteredProviderSelect) sqlScan(ctx context.Context, root *RegisteredProviderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rps.fns))
	for _, fn := range rps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
