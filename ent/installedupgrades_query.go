// Code generated by ent, DO NOT EDIT.
// --------------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2025 The DependencyTrack Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------------
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// --------------------------------------------------------------------

package ent

import (
	"context"
	"fmt"
	"math"

	"dependencytrack.io/hyades/ent/installedupgrades"
	"dependencytrack.io/hyades/ent/predicate"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// INSTALLEDUPGRADESQuery is the builder for querying INSTALLEDUPGRADES entities.
type INSTALLEDUPGRADESQuery struct {
	config
	ctx        *QueryContext
	order      []installedupgrades.OrderOption
	inters     []Interceptor
	predicates []predicate.INSTALLEDUPGRADES
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the INSTALLEDUPGRADESQuery builder.
func (iq *INSTALLEDUPGRADESQuery) Where(ps ...predicate.INSTALLEDUPGRADES) *INSTALLEDUPGRADESQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *INSTALLEDUPGRADESQuery) Limit(limit int) *INSTALLEDUPGRADESQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *INSTALLEDUPGRADESQuery) Offset(offset int) *INSTALLEDUPGRADESQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *INSTALLEDUPGRADESQuery) Unique(unique bool) *INSTALLEDUPGRADESQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *INSTALLEDUPGRADESQuery) Order(o ...installedupgrades.OrderOption) *INSTALLEDUPGRADESQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// First returns the first INSTALLEDUPGRADES entity from the query.
// Returns a *NotFoundError when no INSTALLEDUPGRADES was found.
func (iq *INSTALLEDUPGRADESQuery) First(ctx context.Context) (*INSTALLEDUPGRADES, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{installedupgrades.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *INSTALLEDUPGRADESQuery) FirstX(ctx context.Context) *INSTALLEDUPGRADES {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first INSTALLEDUPGRADES ID from the query.
// Returns a *NotFoundError when no INSTALLEDUPGRADES ID was found.
func (iq *INSTALLEDUPGRADESQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{installedupgrades.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *INSTALLEDUPGRADESQuery) FirstIDX(ctx context.Context) int {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single INSTALLEDUPGRADES entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one INSTALLEDUPGRADES entity is found.
// Returns a *NotFoundError when no INSTALLEDUPGRADES entities are found.
func (iq *INSTALLEDUPGRADESQuery) Only(ctx context.Context) (*INSTALLEDUPGRADES, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{installedupgrades.Label}
	default:
		return nil, &NotSingularError{installedupgrades.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *INSTALLEDUPGRADESQuery) OnlyX(ctx context.Context) *INSTALLEDUPGRADES {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only INSTALLEDUPGRADES ID in the query.
// Returns a *NotSingularError when more than one INSTALLEDUPGRADES ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *INSTALLEDUPGRADESQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{installedupgrades.Label}
	default:
		err = &NotSingularError{installedupgrades.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *INSTALLEDUPGRADESQuery) OnlyIDX(ctx context.Context) int {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of INSTALLEDUPGRADESs.
func (iq *INSTALLEDUPGRADESQuery) All(ctx context.Context) ([]*INSTALLEDUPGRADES, error) {
	ctx = setContextOp(ctx, iq.ctx, ent.OpQueryAll)
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*INSTALLEDUPGRADES, *INSTALLEDUPGRADESQuery]()
	return withInterceptors[[]*INSTALLEDUPGRADES](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *INSTALLEDUPGRADESQuery) AllX(ctx context.Context) []*INSTALLEDUPGRADES {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of INSTALLEDUPGRADES IDs.
func (iq *INSTALLEDUPGRADESQuery) IDs(ctx context.Context) (ids []int, err error) {
	if iq.ctx.Unique == nil && iq.path != nil {
		iq.Unique(true)
	}
	ctx = setContextOp(ctx, iq.ctx, ent.OpQueryIDs)
	if err = iq.Select(installedupgrades.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *INSTALLEDUPGRADESQuery) IDsX(ctx context.Context) []int {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *INSTALLEDUPGRADESQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, ent.OpQueryCount)
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*INSTALLEDUPGRADESQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *INSTALLEDUPGRADESQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *INSTALLEDUPGRADESQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, ent.OpQueryExist)
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *INSTALLEDUPGRADESQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the INSTALLEDUPGRADESQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *INSTALLEDUPGRADESQuery) Clone() *INSTALLEDUPGRADESQuery {
	if iq == nil {
		return nil
	}
	return &INSTALLEDUPGRADESQuery{
		config:     iq.config,
		ctx:        iq.ctx.Clone(),
		order:      append([]installedupgrades.OrderOption{}, iq.order...),
		inters:     append([]Interceptor{}, iq.inters...),
		predicates: append([]predicate.INSTALLEDUPGRADES{}, iq.predicates...),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Endtime time.Time `json:"endtime,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.INSTALLEDUPGRADES.Query().
//		GroupBy(installedupgrades.FieldEndtime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iq *INSTALLEDUPGRADESQuery) GroupBy(field string, fields ...string) *INSTALLEDUPGRADESGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &INSTALLEDUPGRADESGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = installedupgrades.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Endtime time.Time `json:"endtime,omitempty"`
//	}
//
//	client.INSTALLEDUPGRADES.Query().
//		Select(installedupgrades.FieldEndtime).
//		Scan(ctx, &v)
func (iq *INSTALLEDUPGRADESQuery) Select(fields ...string) *INSTALLEDUPGRADESSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &INSTALLEDUPGRADESSelect{INSTALLEDUPGRADESQuery: iq}
	sbuild.label = installedupgrades.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a INSTALLEDUPGRADESSelect configured with the given aggregations.
func (iq *INSTALLEDUPGRADESQuery) Aggregate(fns ...AggregateFunc) *INSTALLEDUPGRADESSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *INSTALLEDUPGRADESQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !installedupgrades.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *INSTALLEDUPGRADESQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*INSTALLEDUPGRADES, error) {
	var (
		nodes = []*INSTALLEDUPGRADES{}
		_spec = iq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*INSTALLEDUPGRADES).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &INSTALLEDUPGRADES{config: iq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (iq *INSTALLEDUPGRADESQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *INSTALLEDUPGRADESQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(installedupgrades.Table, installedupgrades.Columns, sqlgraph.NewFieldSpec(installedupgrades.FieldID, field.TypeInt))
	_spec.From = iq.sql
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iq.path != nil {
		_spec.Unique = true
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, installedupgrades.FieldID)
		for i := range fields {
			if fields[i] != installedupgrades.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *INSTALLEDUPGRADESQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(installedupgrades.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = installedupgrades.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// INSTALLEDUPGRADESGroupBy is the group-by builder for INSTALLEDUPGRADES entities.
type INSTALLEDUPGRADESGroupBy struct {
	selector
	build *INSTALLEDUPGRADESQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *INSTALLEDUPGRADESGroupBy) Aggregate(fns ...AggregateFunc) *INSTALLEDUPGRADESGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *INSTALLEDUPGRADESGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, ent.OpQueryGroupBy)
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*INSTALLEDUPGRADESQuery, *INSTALLEDUPGRADESGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *INSTALLEDUPGRADESGroupBy) sqlScan(ctx context.Context, root *INSTALLEDUPGRADESQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// INSTALLEDUPGRADESSelect is the builder for selecting fields of INSTALLEDUPGRADES entities.
type INSTALLEDUPGRADESSelect struct {
	*INSTALLEDUPGRADESQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *INSTALLEDUPGRADESSelect) Aggregate(fns ...AggregateFunc) *INSTALLEDUPGRADESSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *INSTALLEDUPGRADESSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, ent.OpQuerySelect)
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*INSTALLEDUPGRADESQuery, *INSTALLEDUPGRADESSelect](ctx, is.INSTALLEDUPGRADESQuery, is, is.inters, v)
}

func (is *INSTALLEDUPGRADESSelect) sqlScan(ctx context.Context, root *INSTALLEDUPGRADESQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
