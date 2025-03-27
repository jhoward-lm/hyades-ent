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
	"database/sql/driver"
	"fmt"
	"math"

	"dependencytrack.io/hyades/ent/manageduser"
	"dependencytrack.io/hyades/ent/permission"
	"dependencytrack.io/hyades/ent/predicate"
	"dependencytrack.io/hyades/ent/team"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ManagedUserQuery is the builder for querying ManagedUser entities.
type ManagedUserQuery struct {
	config
	ctx             *QueryContext
	order           []manageduser.OrderOption
	inters          []Interceptor
	predicates      []predicate.ManagedUser
	withPermissions *PermissionQuery
	withTeams       *TeamQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ManagedUserQuery builder.
func (muq *ManagedUserQuery) Where(ps ...predicate.ManagedUser) *ManagedUserQuery {
	muq.predicates = append(muq.predicates, ps...)
	return muq
}

// Limit the number of records to be returned by this query.
func (muq *ManagedUserQuery) Limit(limit int) *ManagedUserQuery {
	muq.ctx.Limit = &limit
	return muq
}

// Offset to start from.
func (muq *ManagedUserQuery) Offset(offset int) *ManagedUserQuery {
	muq.ctx.Offset = &offset
	return muq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (muq *ManagedUserQuery) Unique(unique bool) *ManagedUserQuery {
	muq.ctx.Unique = &unique
	return muq
}

// Order specifies how the records should be ordered.
func (muq *ManagedUserQuery) Order(o ...manageduser.OrderOption) *ManagedUserQuery {
	muq.order = append(muq.order, o...)
	return muq
}

// QueryPermissions chains the current query on the "permissions" edge.
func (muq *ManagedUserQuery) QueryPermissions() *PermissionQuery {
	query := (&PermissionClient{config: muq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := muq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := muq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(manageduser.Table, manageduser.FieldID, selector),
			sqlgraph.To(permission.Table, permission.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, manageduser.PermissionsTable, manageduser.PermissionsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(muq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTeams chains the current query on the "teams" edge.
func (muq *ManagedUserQuery) QueryTeams() *TeamQuery {
	query := (&TeamClient{config: muq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := muq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := muq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(manageduser.Table, manageduser.FieldID, selector),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, manageduser.TeamsTable, manageduser.TeamsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(muq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ManagedUser entity from the query.
// Returns a *NotFoundError when no ManagedUser was found.
func (muq *ManagedUserQuery) First(ctx context.Context) (*ManagedUser, error) {
	nodes, err := muq.Limit(1).All(setContextOp(ctx, muq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{manageduser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (muq *ManagedUserQuery) FirstX(ctx context.Context) *ManagedUser {
	node, err := muq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ManagedUser ID from the query.
// Returns a *NotFoundError when no ManagedUser ID was found.
func (muq *ManagedUserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = muq.Limit(1).IDs(setContextOp(ctx, muq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{manageduser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (muq *ManagedUserQuery) FirstIDX(ctx context.Context) int {
	id, err := muq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ManagedUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ManagedUser entity is found.
// Returns a *NotFoundError when no ManagedUser entities are found.
func (muq *ManagedUserQuery) Only(ctx context.Context) (*ManagedUser, error) {
	nodes, err := muq.Limit(2).All(setContextOp(ctx, muq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{manageduser.Label}
	default:
		return nil, &NotSingularError{manageduser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (muq *ManagedUserQuery) OnlyX(ctx context.Context) *ManagedUser {
	node, err := muq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ManagedUser ID in the query.
// Returns a *NotSingularError when more than one ManagedUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (muq *ManagedUserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = muq.Limit(2).IDs(setContextOp(ctx, muq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{manageduser.Label}
	default:
		err = &NotSingularError{manageduser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (muq *ManagedUserQuery) OnlyIDX(ctx context.Context) int {
	id, err := muq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ManagedUsers.
func (muq *ManagedUserQuery) All(ctx context.Context) ([]*ManagedUser, error) {
	ctx = setContextOp(ctx, muq.ctx, ent.OpQueryAll)
	if err := muq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ManagedUser, *ManagedUserQuery]()
	return withInterceptors[[]*ManagedUser](ctx, muq, qr, muq.inters)
}

// AllX is like All, but panics if an error occurs.
func (muq *ManagedUserQuery) AllX(ctx context.Context) []*ManagedUser {
	nodes, err := muq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ManagedUser IDs.
func (muq *ManagedUserQuery) IDs(ctx context.Context) (ids []int, err error) {
	if muq.ctx.Unique == nil && muq.path != nil {
		muq.Unique(true)
	}
	ctx = setContextOp(ctx, muq.ctx, ent.OpQueryIDs)
	if err = muq.Select(manageduser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (muq *ManagedUserQuery) IDsX(ctx context.Context) []int {
	ids, err := muq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (muq *ManagedUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, muq.ctx, ent.OpQueryCount)
	if err := muq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, muq, querierCount[*ManagedUserQuery](), muq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (muq *ManagedUserQuery) CountX(ctx context.Context) int {
	count, err := muq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (muq *ManagedUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, muq.ctx, ent.OpQueryExist)
	switch _, err := muq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (muq *ManagedUserQuery) ExistX(ctx context.Context) bool {
	exist, err := muq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ManagedUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (muq *ManagedUserQuery) Clone() *ManagedUserQuery {
	if muq == nil {
		return nil
	}
	return &ManagedUserQuery{
		config:          muq.config,
		ctx:             muq.ctx.Clone(),
		order:           append([]manageduser.OrderOption{}, muq.order...),
		inters:          append([]Interceptor{}, muq.inters...),
		predicates:      append([]predicate.ManagedUser{}, muq.predicates...),
		withPermissions: muq.withPermissions.Clone(),
		withTeams:       muq.withTeams.Clone(),
		// clone intermediate query.
		sql:  muq.sql.Clone(),
		path: muq.path,
	}
}

// WithPermissions tells the query-builder to eager-load the nodes that are connected to
// the "permissions" edge. The optional arguments are used to configure the query builder of the edge.
func (muq *ManagedUserQuery) WithPermissions(opts ...func(*PermissionQuery)) *ManagedUserQuery {
	query := (&PermissionClient{config: muq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	muq.withPermissions = query
	return muq
}

// WithTeams tells the query-builder to eager-load the nodes that are connected to
// the "teams" edge. The optional arguments are used to configure the query builder of the edge.
func (muq *ManagedUserQuery) WithTeams(opts ...func(*TeamQuery)) *ManagedUserQuery {
	query := (&TeamClient{config: muq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	muq.withTeams = query
	return muq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Email string `json:"email,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ManagedUser.Query().
//		GroupBy(manageduser.FieldEmail).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (muq *ManagedUserQuery) GroupBy(field string, fields ...string) *ManagedUserGroupBy {
	muq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ManagedUserGroupBy{build: muq}
	grbuild.flds = &muq.ctx.Fields
	grbuild.label = manageduser.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Email string `json:"email,omitempty"`
//	}
//
//	client.ManagedUser.Query().
//		Select(manageduser.FieldEmail).
//		Scan(ctx, &v)
func (muq *ManagedUserQuery) Select(fields ...string) *ManagedUserSelect {
	muq.ctx.Fields = append(muq.ctx.Fields, fields...)
	sbuild := &ManagedUserSelect{ManagedUserQuery: muq}
	sbuild.label = manageduser.Label
	sbuild.flds, sbuild.scan = &muq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ManagedUserSelect configured with the given aggregations.
func (muq *ManagedUserQuery) Aggregate(fns ...AggregateFunc) *ManagedUserSelect {
	return muq.Select().Aggregate(fns...)
}

func (muq *ManagedUserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range muq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, muq); err != nil {
				return err
			}
		}
	}
	for _, f := range muq.ctx.Fields {
		if !manageduser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if muq.path != nil {
		prev, err := muq.path(ctx)
		if err != nil {
			return err
		}
		muq.sql = prev
	}
	return nil
}

func (muq *ManagedUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ManagedUser, error) {
	var (
		nodes       = []*ManagedUser{}
		_spec       = muq.querySpec()
		loadedTypes = [2]bool{
			muq.withPermissions != nil,
			muq.withTeams != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ManagedUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ManagedUser{config: muq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, muq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := muq.withPermissions; query != nil {
		if err := muq.loadPermissions(ctx, query, nodes,
			func(n *ManagedUser) { n.Edges.Permissions = []*Permission{} },
			func(n *ManagedUser, e *Permission) { n.Edges.Permissions = append(n.Edges.Permissions, e) }); err != nil {
			return nil, err
		}
	}
	if query := muq.withTeams; query != nil {
		if err := muq.loadTeams(ctx, query, nodes,
			func(n *ManagedUser) { n.Edges.Teams = []*Team{} },
			func(n *ManagedUser, e *Team) { n.Edges.Teams = append(n.Edges.Teams, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (muq *ManagedUserQuery) loadPermissions(ctx context.Context, query *PermissionQuery, nodes []*ManagedUser, init func(*ManagedUser), assign func(*ManagedUser, *Permission)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*ManagedUser)
	nids := make(map[int]map[*ManagedUser]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(manageduser.PermissionsTable)
		s.Join(joinT).On(s.C(permission.FieldID), joinT.C(manageduser.PermissionsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(manageduser.PermissionsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(manageduser.PermissionsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*ManagedUser]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Permission](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "permissions" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (muq *ManagedUserQuery) loadTeams(ctx context.Context, query *TeamQuery, nodes []*ManagedUser, init func(*ManagedUser), assign func(*ManagedUser, *Team)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*ManagedUser)
	nids := make(map[int]map[*ManagedUser]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(manageduser.TeamsTable)
		s.Join(joinT).On(s.C(team.FieldID), joinT.C(manageduser.TeamsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(manageduser.TeamsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(manageduser.TeamsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*ManagedUser]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Team](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "teams" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (muq *ManagedUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := muq.querySpec()
	_spec.Node.Columns = muq.ctx.Fields
	if len(muq.ctx.Fields) > 0 {
		_spec.Unique = muq.ctx.Unique != nil && *muq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, muq.driver, _spec)
}

func (muq *ManagedUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(manageduser.Table, manageduser.Columns, sqlgraph.NewFieldSpec(manageduser.FieldID, field.TypeInt))
	_spec.From = muq.sql
	if unique := muq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if muq.path != nil {
		_spec.Unique = true
	}
	if fields := muq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, manageduser.FieldID)
		for i := range fields {
			if fields[i] != manageduser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := muq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := muq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := muq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := muq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (muq *ManagedUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(muq.driver.Dialect())
	t1 := builder.Table(manageduser.Table)
	columns := muq.ctx.Fields
	if len(columns) == 0 {
		columns = manageduser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if muq.sql != nil {
		selector = muq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if muq.ctx.Unique != nil && *muq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range muq.predicates {
		p(selector)
	}
	for _, p := range muq.order {
		p(selector)
	}
	if offset := muq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := muq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ManagedUserGroupBy is the group-by builder for ManagedUser entities.
type ManagedUserGroupBy struct {
	selector
	build *ManagedUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mugb *ManagedUserGroupBy) Aggregate(fns ...AggregateFunc) *ManagedUserGroupBy {
	mugb.fns = append(mugb.fns, fns...)
	return mugb
}

// Scan applies the selector query and scans the result into the given value.
func (mugb *ManagedUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mugb.build.ctx, ent.OpQueryGroupBy)
	if err := mugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ManagedUserQuery, *ManagedUserGroupBy](ctx, mugb.build, mugb, mugb.build.inters, v)
}

func (mugb *ManagedUserGroupBy) sqlScan(ctx context.Context, root *ManagedUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mugb.fns))
	for _, fn := range mugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mugb.flds)+len(mugb.fns))
		for _, f := range *mugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ManagedUserSelect is the builder for selecting fields of ManagedUser entities.
type ManagedUserSelect struct {
	*ManagedUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mus *ManagedUserSelect) Aggregate(fns ...AggregateFunc) *ManagedUserSelect {
	mus.fns = append(mus.fns, fns...)
	return mus
}

// Scan applies the selector query and scans the result into the given value.
func (mus *ManagedUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mus.ctx, ent.OpQuerySelect)
	if err := mus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ManagedUserQuery, *ManagedUserSelect](ctx, mus.ManagedUserQuery, mus, mus.inters, v)
}

func (mus *ManagedUserSelect) sqlScan(ctx context.Context, root *ManagedUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mus.fns))
	for _, fn := range mus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
