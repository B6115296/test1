// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team05/ent/bank"
	"github.com/sut63/team05/ent/member"
	"github.com/sut63/team05/ent/officer"
	"github.com/sut63/team05/ent/payback"
	"github.com/sut63/team05/ent/predicate"
	"github.com/sut63/team05/ent/product"
)

// PaybackQuery is the builder for querying Payback entities.
type PaybackQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Payback
	// eager-loading edges.
	withOfficer *OfficerQuery
	withMember  *MemberQuery
	withProduct *ProductQuery
	withBank    *BankQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (pq *PaybackQuery) Where(ps ...predicate.Payback) *PaybackQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PaybackQuery) Limit(limit int) *PaybackQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PaybackQuery) Offset(offset int) *PaybackQuery {
	pq.offset = &offset
	return pq
}

// Order adds an order step to the query.
func (pq *PaybackQuery) Order(o ...OrderFunc) *PaybackQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryOfficer chains the current query on the Officer edge.
func (pq *PaybackQuery) QueryOfficer() *OfficerQuery {
	query := &OfficerQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(payback.Table, payback.FieldID, pq.sqlQuery()),
			sqlgraph.To(officer.Table, officer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, payback.OfficerTable, payback.OfficerColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMember chains the current query on the Member edge.
func (pq *PaybackQuery) QueryMember() *MemberQuery {
	query := &MemberQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(payback.Table, payback.FieldID, pq.sqlQuery()),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, payback.MemberTable, payback.MemberColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProduct chains the current query on the Product edge.
func (pq *PaybackQuery) QueryProduct() *ProductQuery {
	query := &ProductQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(payback.Table, payback.FieldID, pq.sqlQuery()),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, payback.ProductTable, payback.ProductColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBank chains the current query on the Bank edge.
func (pq *PaybackQuery) QueryBank() *BankQuery {
	query := &BankQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(payback.Table, payback.FieldID, pq.sqlQuery()),
			sqlgraph.To(bank.Table, bank.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, payback.BankTable, payback.BankColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Payback entity in the query. Returns *NotFoundError when no payback was found.
func (pq *PaybackQuery) First(ctx context.Context) (*Payback, error) {
	pas, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(pas) == 0 {
		return nil, &NotFoundError{payback.Label}
	}
	return pas[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PaybackQuery) FirstX(ctx context.Context) *Payback {
	pa, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return pa
}

// FirstID returns the first Payback id in the query. Returns *NotFoundError when no id was found.
func (pq *PaybackQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{payback.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (pq *PaybackQuery) FirstXID(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Payback entity in the query, returns an error if not exactly one entity was returned.
func (pq *PaybackQuery) Only(ctx context.Context) (*Payback, error) {
	pas, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(pas) {
	case 1:
		return pas[0], nil
	case 0:
		return nil, &NotFoundError{payback.Label}
	default:
		return nil, &NotSingularError{payback.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PaybackQuery) OnlyX(ctx context.Context) *Payback {
	pa, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return pa
}

// OnlyID returns the only Payback id in the query, returns an error if not exactly one id was returned.
func (pq *PaybackQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{payback.Label}
	default:
		err = &NotSingularError{payback.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PaybackQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Paybacks.
func (pq *PaybackQuery) All(ctx context.Context) ([]*Payback, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *PaybackQuery) AllX(ctx context.Context) []*Payback {
	pas, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return pas
}

// IDs executes the query and returns a list of Payback ids.
func (pq *PaybackQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pq.Select(payback.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PaybackQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PaybackQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PaybackQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PaybackQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PaybackQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PaybackQuery) Clone() *PaybackQuery {
	return &PaybackQuery{
		config:     pq.config,
		limit:      pq.limit,
		offset:     pq.offset,
		order:      append([]OrderFunc{}, pq.order...),
		unique:     append([]string{}, pq.unique...),
		predicates: append([]predicate.Payback{}, pq.predicates...),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

//  WithOfficer tells the query-builder to eager-loads the nodes that are connected to
// the "Officer" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PaybackQuery) WithOfficer(opts ...func(*OfficerQuery)) *PaybackQuery {
	query := &OfficerQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withOfficer = query
	return pq
}

//  WithMember tells the query-builder to eager-loads the nodes that are connected to
// the "Member" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PaybackQuery) WithMember(opts ...func(*MemberQuery)) *PaybackQuery {
	query := &MemberQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withMember = query
	return pq
}

//  WithProduct tells the query-builder to eager-loads the nodes that are connected to
// the "Product" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PaybackQuery) WithProduct(opts ...func(*ProductQuery)) *PaybackQuery {
	query := &ProductQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withProduct = query
	return pq
}

//  WithBank tells the query-builder to eager-loads the nodes that are connected to
// the "Bank" edge. The optional arguments used to configure the query builder of the edge.
func (pq *PaybackQuery) WithBank(opts ...func(*BankQuery)) *PaybackQuery {
	query := &BankQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withBank = query
	return pq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Accountnumber string `json:"Accountnumber,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Payback.Query().
//		GroupBy(payback.FieldAccountnumber).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pq *PaybackQuery) GroupBy(field string, fields ...string) *PaybackGroupBy {
	group := &PaybackGroupBy{config: pq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Accountnumber string `json:"Accountnumber,omitempty"`
//	}
//
//	client.Payback.Query().
//		Select(payback.FieldAccountnumber).
//		Scan(ctx, &v)
//
func (pq *PaybackQuery) Select(field string, fields ...string) *PaybackSelect {
	selector := &PaybackSelect{config: pq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(), nil
	}
	return selector
}

func (pq *PaybackQuery) prepareQuery(ctx context.Context) error {
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PaybackQuery) sqlAll(ctx context.Context) ([]*Payback, error) {
	var (
		nodes       = []*Payback{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [4]bool{
			pq.withOfficer != nil,
			pq.withMember != nil,
			pq.withProduct != nil,
			pq.withBank != nil,
		}
	)
	if pq.withOfficer != nil || pq.withMember != nil || pq.withProduct != nil || pq.withBank != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, payback.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Payback{config: pq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pq.withOfficer; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Payback)
		for i := range nodes {
			if fk := nodes[i].officer_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(officer.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "officer_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Officer = n
			}
		}
	}

	if query := pq.withMember; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Payback)
		for i := range nodes {
			if fk := nodes[i].member_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(member.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "member_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Member = n
			}
		}
	}

	if query := pq.withProduct; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Payback)
		for i := range nodes {
			if fk := nodes[i].product_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(product.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "product_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Product = n
			}
		}
	}

	if query := pq.withBank; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Payback)
		for i := range nodes {
			if fk := nodes[i].bank_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(bank.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "bank_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Bank = n
			}
		}
	}

	return nodes, nil
}

func (pq *PaybackQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PaybackQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (pq *PaybackQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   payback.Table,
			Columns: payback.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: payback.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PaybackQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(payback.Table)
	selector := builder.Select(t1.Columns(payback.Columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(payback.Columns...)...)
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PaybackGroupBy is the builder for group-by Payback entities.
type PaybackGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PaybackGroupBy) Aggregate(fns ...AggregateFunc) *PaybackGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scan the result into the given value.
func (pgb *PaybackGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.sql = query
	return pgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pgb *PaybackGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (pgb *PaybackGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PaybackGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pgb *PaybackGroupBy) StringsX(ctx context.Context) []string {
	v, err := pgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (pgb *PaybackGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{payback.Label}
	default:
		err = fmt.Errorf("ent: PaybackGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pgb *PaybackGroupBy) StringX(ctx context.Context) string {
	v, err := pgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (pgb *PaybackGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PaybackGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pgb *PaybackGroupBy) IntsX(ctx context.Context) []int {
	v, err := pgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (pgb *PaybackGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{payback.Label}
	default:
		err = fmt.Errorf("ent: PaybackGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pgb *PaybackGroupBy) IntX(ctx context.Context) int {
	v, err := pgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (pgb *PaybackGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PaybackGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pgb *PaybackGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (pgb *PaybackGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{payback.Label}
	default:
		err = fmt.Errorf("ent: PaybackGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pgb *PaybackGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (pgb *PaybackGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PaybackGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pgb *PaybackGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (pgb *PaybackGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{payback.Label}
	default:
		err = fmt.Errorf("ent: PaybackGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pgb *PaybackGroupBy) BoolX(ctx context.Context) bool {
	v, err := pgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pgb *PaybackGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pgb.sqlQuery().Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *PaybackGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql
	columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
	columns = append(columns, pgb.fields...)
	for _, fn := range pgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(pgb.fields...)
}

// PaybackSelect is the builder for select fields of Payback entities.
type PaybackSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ps *PaybackSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ps.path(ctx)
	if err != nil {
		return err
	}
	ps.sql = query
	return ps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ps *PaybackSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ps *PaybackSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PaybackSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ps *PaybackSelect) StringsX(ctx context.Context) []string {
	v, err := ps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ps *PaybackSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{payback.Label}
	default:
		err = fmt.Errorf("ent: PaybackSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ps *PaybackSelect) StringX(ctx context.Context) string {
	v, err := ps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ps *PaybackSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PaybackSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ps *PaybackSelect) IntsX(ctx context.Context) []int {
	v, err := ps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ps *PaybackSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{payback.Label}
	default:
		err = fmt.Errorf("ent: PaybackSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ps *PaybackSelect) IntX(ctx context.Context) int {
	v, err := ps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ps *PaybackSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PaybackSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ps *PaybackSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ps *PaybackSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{payback.Label}
	default:
		err = fmt.Errorf("ent: PaybackSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ps *PaybackSelect) Float64X(ctx context.Context) float64 {
	v, err := ps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ps *PaybackSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PaybackSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ps *PaybackSelect) BoolsX(ctx context.Context) []bool {
	v, err := ps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ps *PaybackSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{payback.Label}
	default:
		err = fmt.Errorf("ent: PaybackSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ps *PaybackSelect) BoolX(ctx context.Context) bool {
	v, err := ps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ps *PaybackSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ps.sqlQuery().Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ps *PaybackSelect) sqlQuery() sql.Querier {
	selector := ps.sql
	selector.Select(selector.Columns(ps.fields...)...)
	return selector
}