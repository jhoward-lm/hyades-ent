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

	"dependencytrack.io/hyades/ent/mappedldapgroup"
	"dependencytrack.io/hyades/ent/predicate"
	"dependencytrack.io/hyades/ent/team"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MappedLDAPGroupQuery is the builder for querying MappedLDAPGroup entities.
type MappedLDAPGroupQuery struct {
	config
	ctx        *QueryContext
	order      []mappedldapgroup.OrderOption
	inters     []Interceptor
	predicates []predicate.MappedLDAPGroup
	withTeam   *TeamQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MappedLDAPGroupQuery builder.
func (mlgq *MappedLDAPGroupQuery) Where(ps ...predicate.MappedLDAPGroup) *MappedLDAPGroupQuery {
	mlgq.predicates = append(mlgq.predicates, ps...)
	return mlgq
}

// Limit the number of records to be returned by this query.
func (mlgq *MappedLDAPGroupQuery) Limit(limit int) *MappedLDAPGroupQuery {
	mlgq.ctx.Limit = &limit
	return mlgq
}

// Offset to start from.
func (mlgq *MappedLDAPGroupQuery) Offset(offset int) *MappedLDAPGroupQuery {
	mlgq.ctx.Offset = &offset
	return mlgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mlgq *MappedLDAPGroupQuery) Unique(unique bool) *MappedLDAPGroupQuery {
	mlgq.ctx.Unique = &unique
	return mlgq
}

// Order specifies how the records should be ordered.
func (mlgq *MappedLDAPGroupQuery) Order(o ...mappedldapgroup.OrderOption) *MappedLDAPGroupQuery {
	mlgq.order = append(mlgq.order, o...)
	return mlgq
}

// QueryTeam chains the current query on the "team" edge.
func (mlgq *MappedLDAPGroupQuery) QueryTeam() *TeamQuery {
	query := (&TeamClient{config: mlgq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mlgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mlgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(mappedldapgroup.Table, mappedldapgroup.FieldID, selector),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, mappedldapgroup.TeamTable, mappedldapgroup.TeamColumn),
		)
		fromU = sqlgraph.SetNeighbors(mlgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MappedLDAPGroup entity from the query.
// Returns a *NotFoundError when no MappedLDAPGroup was found.
func (mlgq *MappedLDAPGroupQuery) First(ctx context.Context) (*MappedLDAPGroup, error) {
	nodes, err := mlgq.Limit(1).All(setContextOp(ctx, mlgq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{mappedldapgroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mlgq *MappedLDAPGroupQuery) FirstX(ctx context.Context) *MappedLDAPGroup {
	node, err := mlgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MappedLDAPGroup ID from the query.
// Returns a *NotFoundError when no MappedLDAPGroup ID was found.
func (mlgq *MappedLDAPGroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mlgq.Limit(1).IDs(setContextOp(ctx, mlgq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{mappedldapgroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mlgq *MappedLDAPGroupQuery) FirstIDX(ctx context.Context) int {
	id, err := mlgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MappedLDAPGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MappedLDAPGroup entity is found.
// Returns a *NotFoundError when no MappedLDAPGroup entities are found.
func (mlgq *MappedLDAPGroupQuery) Only(ctx context.Context) (*MappedLDAPGroup, error) {
	nodes, err := mlgq.Limit(2).All(setContextOp(ctx, mlgq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{mappedldapgroup.Label}
	default:
		return nil, &NotSingularError{mappedldapgroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mlgq *MappedLDAPGroupQuery) OnlyX(ctx context.Context) *MappedLDAPGroup {
	node, err := mlgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MappedLDAPGroup ID in the query.
// Returns a *NotSingularError when more than one MappedLDAPGroup ID is found.
// Returns a *NotFoundError when no entities are found.
func (mlgq *MappedLDAPGroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mlgq.Limit(2).IDs(setContextOp(ctx, mlgq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{mappedldapgroup.Label}
	default:
		err = &NotSingularError{mappedldapgroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mlgq *MappedLDAPGroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := mlgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MappedLDAPGroups.
func (mlgq *MappedLDAPGroupQuery) All(ctx context.Context) ([]*MappedLDAPGroup, error) {
	ctx = setContextOp(ctx, mlgq.ctx, ent.OpQueryAll)
	if err := mlgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MappedLDAPGroup, *MappedLDAPGroupQuery]()
	return withInterceptors[[]*MappedLDAPGroup](ctx, mlgq, qr, mlgq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mlgq *MappedLDAPGroupQuery) AllX(ctx context.Context) []*MappedLDAPGroup {
	nodes, err := mlgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MappedLDAPGroup IDs.
func (mlgq *MappedLDAPGroupQuery) IDs(ctx context.Context) (ids []int, err error) {
	if mlgq.ctx.Unique == nil && mlgq.path != nil {
		mlgq.Unique(true)
	}
	ctx = setContextOp(ctx, mlgq.ctx, ent.OpQueryIDs)
	if err = mlgq.Select(mappedldapgroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mlgq *MappedLDAPGroupQuery) IDsX(ctx context.Context) []int {
	ids, err := mlgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mlgq *MappedLDAPGroupQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mlgq.ctx, ent.OpQueryCount)
	if err := mlgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mlgq, querierCount[*MappedLDAPGroupQuery](), mlgq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mlgq *MappedLDAPGroupQuery) CountX(ctx context.Context) int {
	count, err := mlgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mlgq *MappedLDAPGroupQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mlgq.ctx, ent.OpQueryExist)
	switch _, err := mlgq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mlgq *MappedLDAPGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := mlgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MappedLDAPGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mlgq *MappedLDAPGroupQuery) Clone() *MappedLDAPGroupQuery {
	if mlgq == nil {
		return nil
	}
	return &MappedLDAPGroupQuery{
		config:     mlgq.config,
		ctx:        mlgq.ctx.Clone(),
		order:      append([]mappedldapgroup.OrderOption{}, mlgq.order...),
		inters:     append([]Interceptor{}, mlgq.inters...),
		predicates: append([]predicate.MappedLDAPGroup{}, mlgq.predicates...),
		withTeam:   mlgq.withTeam.Clone(),
		// clone intermediate query.
		sql:  mlgq.sql.Clone(),
		path: mlgq.path,
	}
}

// WithTeam tells the query-builder to eager-load the nodes that are connected to
// the "team" edge. The optional arguments are used to configure the query builder of the edge.
func (mlgq *MappedLDAPGroupQuery) WithTeam(opts ...func(*TeamQuery)) *MappedLDAPGroupQuery {
	query := (&TeamClient{config: mlgq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mlgq.withTeam = query
	return mlgq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UUID uuid.UUID `json:"uuid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MappedLDAPGroup.Query().
//		GroupBy(mappedldapgroup.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mlgq *MappedLDAPGroupQuery) GroupBy(field string, fields ...string) *MappedLDAPGroupGroupBy {
	mlgq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MappedLDAPGroupGroupBy{build: mlgq}
	grbuild.flds = &mlgq.ctx.Fields
	grbuild.label = mappedldapgroup.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UUID uuid.UUID `json:"uuid,omitempty"`
//	}
//
//	client.MappedLDAPGroup.Query().
//		Select(mappedldapgroup.FieldUUID).
//		Scan(ctx, &v)
func (mlgq *MappedLDAPGroupQuery) Select(fields ...string) *MappedLDAPGroupSelect {
	mlgq.ctx.Fields = append(mlgq.ctx.Fields, fields...)
	sbuild := &MappedLDAPGroupSelect{MappedLDAPGroupQuery: mlgq}
	sbuild.label = mappedldapgroup.Label
	sbuild.flds, sbuild.scan = &mlgq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MappedLDAPGroupSelect configured with the given aggregations.
func (mlgq *MappedLDAPGroupQuery) Aggregate(fns ...AggregateFunc) *MappedLDAPGroupSelect {
	return mlgq.Select().Aggregate(fns...)
}

func (mlgq *MappedLDAPGroupQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mlgq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mlgq); err != nil {
				return err
			}
		}
	}
	for _, f := range mlgq.ctx.Fields {
		if !mappedldapgroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mlgq.path != nil {
		prev, err := mlgq.path(ctx)
		if err != nil {
			return err
		}
		mlgq.sql = prev
	}
	return nil
}

func (mlgq *MappedLDAPGroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MappedLDAPGroup, error) {
	var (
		nodes       = []*MappedLDAPGroup{}
		_spec       = mlgq.querySpec()
		loadedTypes = [1]bool{
			mlgq.withTeam != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MappedLDAPGroup).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MappedLDAPGroup{config: mlgq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mlgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mlgq.withTeam; query != nil {
		if err := mlgq.loadTeam(ctx, query, nodes, nil,
			func(n *MappedLDAPGroup, e *Team) { n.Edges.Team = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mlgq *MappedLDAPGroupQuery) loadTeam(ctx context.Context, query *TeamQuery, nodes []*MappedLDAPGroup, init func(*MappedLDAPGroup), assign func(*MappedLDAPGroup, *Team)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*MappedLDAPGroup)
	for i := range nodes {
		fk := nodes[i].TeamID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(team.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "team_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (mlgq *MappedLDAPGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mlgq.querySpec()
	_spec.Node.Columns = mlgq.ctx.Fields
	if len(mlgq.ctx.Fields) > 0 {
		_spec.Unique = mlgq.ctx.Unique != nil && *mlgq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mlgq.driver, _spec)
}

func (mlgq *MappedLDAPGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(mappedldapgroup.Table, mappedldapgroup.Columns, sqlgraph.NewFieldSpec(mappedldapgroup.FieldID, field.TypeInt))
	_spec.From = mlgq.sql
	if unique := mlgq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mlgq.path != nil {
		_spec.Unique = true
	}
	if fields := mlgq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mappedldapgroup.FieldID)
		for i := range fields {
			if fields[i] != mappedldapgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if mlgq.withTeam != nil {
			_spec.Node.AddColumnOnce(mappedldapgroup.FieldTeamID)
		}
	}
	if ps := mlgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mlgq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mlgq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mlgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mlgq *MappedLDAPGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mlgq.driver.Dialect())
	t1 := builder.Table(mappedldapgroup.Table)
	columns := mlgq.ctx.Fields
	if len(columns) == 0 {
		columns = mappedldapgroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mlgq.sql != nil {
		selector = mlgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mlgq.ctx.Unique != nil && *mlgq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mlgq.predicates {
		p(selector)
	}
	for _, p := range mlgq.order {
		p(selector)
	}
	if offset := mlgq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mlgq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MappedLDAPGroupGroupBy is the group-by builder for MappedLDAPGroup entities.
type MappedLDAPGroupGroupBy struct {
	selector
	build *MappedLDAPGroupQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mlggb *MappedLDAPGroupGroupBy) Aggregate(fns ...AggregateFunc) *MappedLDAPGroupGroupBy {
	mlggb.fns = append(mlggb.fns, fns...)
	return mlggb
}

// Scan applies the selector query and scans the result into the given value.
func (mlggb *MappedLDAPGroupGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mlggb.build.ctx, ent.OpQueryGroupBy)
	if err := mlggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MappedLDAPGroupQuery, *MappedLDAPGroupGroupBy](ctx, mlggb.build, mlggb, mlggb.build.inters, v)
}

func (mlggb *MappedLDAPGroupGroupBy) sqlScan(ctx context.Context, root *MappedLDAPGroupQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mlggb.fns))
	for _, fn := range mlggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mlggb.flds)+len(mlggb.fns))
		for _, f := range *mlggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mlggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mlggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MappedLDAPGroupSelect is the builder for selecting fields of MappedLDAPGroup entities.
type MappedLDAPGroupSelect struct {
	*MappedLDAPGroupQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mlgs *MappedLDAPGroupSelect) Aggregate(fns ...AggregateFunc) *MappedLDAPGroupSelect {
	mlgs.fns = append(mlgs.fns, fns...)
	return mlgs
}

// Scan applies the selector query and scans the result into the given value.
func (mlgs *MappedLDAPGroupSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mlgs.ctx, ent.OpQuerySelect)
	if err := mlgs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MappedLDAPGroupQuery, *MappedLDAPGroupSelect](ctx, mlgs.MappedLDAPGroupQuery, mlgs, mlgs.inters, v)
}

func (mlgs *MappedLDAPGroupSelect) sqlScan(ctx context.Context, root *MappedLDAPGroupQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mlgs.fns))
	for _, fn := range mlgs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mlgs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mlgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
