// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vicanso/forest/ent/predicate"
	"github.com/vicanso/forest/ent/userlogin"
)

// UserLoginQuery is the builder for querying UserLogin entities.
type UserLoginQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	predicates []predicate.UserLogin
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (ulq *UserLoginQuery) Where(ps ...predicate.UserLogin) *UserLoginQuery {
	ulq.predicates = append(ulq.predicates, ps...)
	return ulq
}

// Limit adds a limit step to the query.
func (ulq *UserLoginQuery) Limit(limit int) *UserLoginQuery {
	ulq.limit = &limit
	return ulq
}

// Offset adds an offset step to the query.
func (ulq *UserLoginQuery) Offset(offset int) *UserLoginQuery {
	ulq.offset = &offset
	return ulq
}

// Order adds an order step to the query.
func (ulq *UserLoginQuery) Order(o ...OrderFunc) *UserLoginQuery {
	ulq.order = append(ulq.order, o...)
	return ulq
}

// First returns the first UserLogin entity in the query. Returns *NotFoundError when no userlogin was found.
func (ulq *UserLoginQuery) First(ctx context.Context) (*UserLogin, error) {
	nodes, err := ulq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userlogin.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ulq *UserLoginQuery) FirstX(ctx context.Context) *UserLogin {
	node, err := ulq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserLogin id in the query. Returns *NotFoundError when no id was found.
func (ulq *UserLoginQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ulq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userlogin.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ulq *UserLoginQuery) FirstIDX(ctx context.Context) int {
	id, err := ulq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only UserLogin entity in the query, returns an error if not exactly one entity was returned.
func (ulq *UserLoginQuery) Only(ctx context.Context) (*UserLogin, error) {
	nodes, err := ulq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userlogin.Label}
	default:
		return nil, &NotSingularError{userlogin.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ulq *UserLoginQuery) OnlyX(ctx context.Context) *UserLogin {
	node, err := ulq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only UserLogin id in the query, returns an error if not exactly one id was returned.
func (ulq *UserLoginQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ulq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userlogin.Label}
	default:
		err = &NotSingularError{userlogin.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ulq *UserLoginQuery) OnlyIDX(ctx context.Context) int {
	id, err := ulq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserLogins.
func (ulq *UserLoginQuery) All(ctx context.Context) ([]*UserLogin, error) {
	if err := ulq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ulq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ulq *UserLoginQuery) AllX(ctx context.Context) []*UserLogin {
	nodes, err := ulq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserLogin ids.
func (ulq *UserLoginQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ulq.Select(userlogin.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ulq *UserLoginQuery) IDsX(ctx context.Context) []int {
	ids, err := ulq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ulq *UserLoginQuery) Count(ctx context.Context) (int, error) {
	if err := ulq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ulq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ulq *UserLoginQuery) CountX(ctx context.Context) int {
	count, err := ulq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ulq *UserLoginQuery) Exist(ctx context.Context) (bool, error) {
	if err := ulq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ulq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ulq *UserLoginQuery) ExistX(ctx context.Context) bool {
	exist, err := ulq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ulq *UserLoginQuery) Clone() *UserLoginQuery {
	if ulq == nil {
		return nil
	}
	return &UserLoginQuery{
		config:     ulq.config,
		limit:      ulq.limit,
		offset:     ulq.offset,
		order:      append([]OrderFunc{}, ulq.order...),
		predicates: append([]predicate.UserLogin{}, ulq.predicates...),
		// clone intermediate query.
		sql:  ulq.sql.Clone(),
		path: ulq.path,
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty" sql:"created_at"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserLogin.Query().
//		GroupBy(userlogin.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ulq *UserLoginQuery) GroupBy(field string, fields ...string) *UserLoginGroupBy {
	group := &UserLoginGroupBy{config: ulq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ulq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ulq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty" sql:"created_at"`
//	}
//
//	client.UserLogin.Query().
//		Select(userlogin.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (ulq *UserLoginQuery) Select(field string, fields ...string) *UserLoginSelect {
	selector := &UserLoginSelect{config: ulq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ulq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ulq.sqlQuery(), nil
	}
	return selector
}

func (ulq *UserLoginQuery) prepareQuery(ctx context.Context) error {
	if ulq.path != nil {
		prev, err := ulq.path(ctx)
		if err != nil {
			return err
		}
		ulq.sql = prev
	}
	return nil
}

func (ulq *UserLoginQuery) sqlAll(ctx context.Context) ([]*UserLogin, error) {
	var (
		nodes = []*UserLogin{}
		_spec = ulq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &UserLogin{config: ulq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, ulq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ulq *UserLoginQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ulq.querySpec()
	return sqlgraph.CountNodes(ctx, ulq.driver, _spec)
}

func (ulq *UserLoginQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ulq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (ulq *UserLoginQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userlogin.Table,
			Columns: userlogin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userlogin.FieldID,
			},
		},
		From:   ulq.sql,
		Unique: true,
	}
	if ps := ulq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ulq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ulq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ulq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, userlogin.ValidColumn)
			}
		}
	}
	return _spec
}

func (ulq *UserLoginQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(ulq.driver.Dialect())
	t1 := builder.Table(userlogin.Table)
	selector := builder.Select(t1.Columns(userlogin.Columns...)...).From(t1)
	if ulq.sql != nil {
		selector = ulq.sql
		selector.Select(selector.Columns(userlogin.Columns...)...)
	}
	for _, p := range ulq.predicates {
		p(selector)
	}
	for _, p := range ulq.order {
		p(selector, userlogin.ValidColumn)
	}
	if offset := ulq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ulq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserLoginGroupBy is the builder for group-by UserLogin entities.
type UserLoginGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ulgb *UserLoginGroupBy) Aggregate(fns ...AggregateFunc) *UserLoginGroupBy {
	ulgb.fns = append(ulgb.fns, fns...)
	return ulgb
}

// Scan applies the group-by query and scan the result into the given value.
func (ulgb *UserLoginGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ulgb.path(ctx)
	if err != nil {
		return err
	}
	ulgb.sql = query
	return ulgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ulgb *UserLoginGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ulgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ulgb *UserLoginGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ulgb.fields) > 1 {
		return nil, errors.New("ent: UserLoginGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ulgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ulgb *UserLoginGroupBy) StringsX(ctx context.Context) []string {
	v, err := ulgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (ulgb *UserLoginGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ulgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userlogin.Label}
	default:
		err = fmt.Errorf("ent: UserLoginGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ulgb *UserLoginGroupBy) StringX(ctx context.Context) string {
	v, err := ulgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ulgb *UserLoginGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ulgb.fields) > 1 {
		return nil, errors.New("ent: UserLoginGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ulgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ulgb *UserLoginGroupBy) IntsX(ctx context.Context) []int {
	v, err := ulgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (ulgb *UserLoginGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ulgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userlogin.Label}
	default:
		err = fmt.Errorf("ent: UserLoginGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ulgb *UserLoginGroupBy) IntX(ctx context.Context) int {
	v, err := ulgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ulgb *UserLoginGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ulgb.fields) > 1 {
		return nil, errors.New("ent: UserLoginGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ulgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ulgb *UserLoginGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ulgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (ulgb *UserLoginGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ulgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userlogin.Label}
	default:
		err = fmt.Errorf("ent: UserLoginGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ulgb *UserLoginGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ulgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ulgb *UserLoginGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ulgb.fields) > 1 {
		return nil, errors.New("ent: UserLoginGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ulgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ulgb *UserLoginGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ulgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (ulgb *UserLoginGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ulgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userlogin.Label}
	default:
		err = fmt.Errorf("ent: UserLoginGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ulgb *UserLoginGroupBy) BoolX(ctx context.Context) bool {
	v, err := ulgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ulgb *UserLoginGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ulgb.fields {
		if !userlogin.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ulgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ulgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ulgb *UserLoginGroupBy) sqlQuery() *sql.Selector {
	selector := ulgb.sql
	columns := make([]string, 0, len(ulgb.fields)+len(ulgb.fns))
	columns = append(columns, ulgb.fields...)
	for _, fn := range ulgb.fns {
		columns = append(columns, fn(selector, userlogin.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(ulgb.fields...)
}

// UserLoginSelect is the builder for select fields of UserLogin entities.
type UserLoginSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (uls *UserLoginSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := uls.path(ctx)
	if err != nil {
		return err
	}
	uls.sql = query
	return uls.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uls *UserLoginSelect) ScanX(ctx context.Context, v interface{}) {
	if err := uls.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (uls *UserLoginSelect) Strings(ctx context.Context) ([]string, error) {
	if len(uls.fields) > 1 {
		return nil, errors.New("ent: UserLoginSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := uls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uls *UserLoginSelect) StringsX(ctx context.Context) []string {
	v, err := uls.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (uls *UserLoginSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uls.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userlogin.Label}
	default:
		err = fmt.Errorf("ent: UserLoginSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uls *UserLoginSelect) StringX(ctx context.Context) string {
	v, err := uls.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (uls *UserLoginSelect) Ints(ctx context.Context) ([]int, error) {
	if len(uls.fields) > 1 {
		return nil, errors.New("ent: UserLoginSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := uls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uls *UserLoginSelect) IntsX(ctx context.Context) []int {
	v, err := uls.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (uls *UserLoginSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uls.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userlogin.Label}
	default:
		err = fmt.Errorf("ent: UserLoginSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uls *UserLoginSelect) IntX(ctx context.Context) int {
	v, err := uls.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (uls *UserLoginSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(uls.fields) > 1 {
		return nil, errors.New("ent: UserLoginSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := uls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uls *UserLoginSelect) Float64sX(ctx context.Context) []float64 {
	v, err := uls.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (uls *UserLoginSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uls.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userlogin.Label}
	default:
		err = fmt.Errorf("ent: UserLoginSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uls *UserLoginSelect) Float64X(ctx context.Context) float64 {
	v, err := uls.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (uls *UserLoginSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(uls.fields) > 1 {
		return nil, errors.New("ent: UserLoginSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := uls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uls *UserLoginSelect) BoolsX(ctx context.Context) []bool {
	v, err := uls.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (uls *UserLoginSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uls.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userlogin.Label}
	default:
		err = fmt.Errorf("ent: UserLoginSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uls *UserLoginSelect) BoolX(ctx context.Context) bool {
	v, err := uls.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uls *UserLoginSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range uls.fields {
		if !userlogin.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := uls.sqlQuery().Query()
	if err := uls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uls *UserLoginSelect) sqlQuery() sql.Querier {
	selector := uls.sql
	selector.Select(selector.Columns(uls.fields...)...)
	return selector
}
