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
	"errors"
	"fmt"

	"dependencytrack.io/hyades/ent/mappedoidcgroup"
	"dependencytrack.io/hyades/ent/oidcgroup"
	"dependencytrack.io/hyades/ent/predicate"
	"dependencytrack.io/hyades/ent/team"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OIDCGroupUpdate is the builder for updating OIDCGroup entities.
type OIDCGroupUpdate struct {
	config
	hooks    []Hook
	mutation *OIDCGroupMutation
}

// Where appends a list predicates to the OIDCGroupUpdate builder.
func (ogu *OIDCGroupUpdate) Where(ps ...predicate.OIDCGroup) *OIDCGroupUpdate {
	ogu.mutation.Where(ps...)
	return ogu
}

// SetName sets the "name" field.
func (ogu *OIDCGroupUpdate) SetName(s string) *OIDCGroupUpdate {
	ogu.mutation.SetName(s)
	return ogu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ogu *OIDCGroupUpdate) SetNillableName(s *string) *OIDCGroupUpdate {
	if s != nil {
		ogu.SetName(*s)
	}
	return ogu
}

// AddMappedOidcGroupIDs adds the "mapped_oidc_groups" edge to the MappedOIDCGroup entity by IDs.
func (ogu *OIDCGroupUpdate) AddMappedOidcGroupIDs(ids ...int) *OIDCGroupUpdate {
	ogu.mutation.AddMappedOidcGroupIDs(ids...)
	return ogu
}

// AddMappedOidcGroups adds the "mapped_oidc_groups" edges to the MappedOIDCGroup entity.
func (ogu *OIDCGroupUpdate) AddMappedOidcGroups(m ...*MappedOIDCGroup) *OIDCGroupUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ogu.AddMappedOidcGroupIDs(ids...)
}

// AddTeamIDs adds the "team" edge to the Team entity by IDs.
func (ogu *OIDCGroupUpdate) AddTeamIDs(ids ...int) *OIDCGroupUpdate {
	ogu.mutation.AddTeamIDs(ids...)
	return ogu
}

// AddTeam adds the "team" edges to the Team entity.
func (ogu *OIDCGroupUpdate) AddTeam(t ...*Team) *OIDCGroupUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ogu.AddTeamIDs(ids...)
}

// Mutation returns the OIDCGroupMutation object of the builder.
func (ogu *OIDCGroupUpdate) Mutation() *OIDCGroupMutation {
	return ogu.mutation
}

// ClearMappedOidcGroups clears all "mapped_oidc_groups" edges to the MappedOIDCGroup entity.
func (ogu *OIDCGroupUpdate) ClearMappedOidcGroups() *OIDCGroupUpdate {
	ogu.mutation.ClearMappedOidcGroups()
	return ogu
}

// RemoveMappedOidcGroupIDs removes the "mapped_oidc_groups" edge to MappedOIDCGroup entities by IDs.
func (ogu *OIDCGroupUpdate) RemoveMappedOidcGroupIDs(ids ...int) *OIDCGroupUpdate {
	ogu.mutation.RemoveMappedOidcGroupIDs(ids...)
	return ogu
}

// RemoveMappedOidcGroups removes "mapped_oidc_groups" edges to MappedOIDCGroup entities.
func (ogu *OIDCGroupUpdate) RemoveMappedOidcGroups(m ...*MappedOIDCGroup) *OIDCGroupUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ogu.RemoveMappedOidcGroupIDs(ids...)
}

// ClearTeam clears all "team" edges to the Team entity.
func (ogu *OIDCGroupUpdate) ClearTeam() *OIDCGroupUpdate {
	ogu.mutation.ClearTeam()
	return ogu
}

// RemoveTeamIDs removes the "team" edge to Team entities by IDs.
func (ogu *OIDCGroupUpdate) RemoveTeamIDs(ids ...int) *OIDCGroupUpdate {
	ogu.mutation.RemoveTeamIDs(ids...)
	return ogu
}

// RemoveTeam removes "team" edges to Team entities.
func (ogu *OIDCGroupUpdate) RemoveTeam(t ...*Team) *OIDCGroupUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ogu.RemoveTeamIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ogu *OIDCGroupUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ogu.sqlSave, ogu.mutation, ogu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ogu *OIDCGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := ogu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ogu *OIDCGroupUpdate) Exec(ctx context.Context) error {
	_, err := ogu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ogu *OIDCGroupUpdate) ExecX(ctx context.Context) {
	if err := ogu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ogu *OIDCGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(oidcgroup.Table, oidcgroup.Columns, sqlgraph.NewFieldSpec(oidcgroup.FieldID, field.TypeInt))
	if ps := ogu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ogu.mutation.Name(); ok {
		_spec.SetField(oidcgroup.FieldName, field.TypeString, value)
	}
	if ogu.mutation.MappedOidcGroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   oidcgroup.MappedOidcGroupsTable,
			Columns: []string{oidcgroup.MappedOidcGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mappedoidcgroup.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ogu.mutation.RemovedMappedOidcGroupsIDs(); len(nodes) > 0 && !ogu.mutation.MappedOidcGroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   oidcgroup.MappedOidcGroupsTable,
			Columns: []string{oidcgroup.MappedOidcGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mappedoidcgroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ogu.mutation.MappedOidcGroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   oidcgroup.MappedOidcGroupsTable,
			Columns: []string{oidcgroup.MappedOidcGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mappedoidcgroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ogu.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oidcgroup.TeamTable,
			Columns: oidcgroup.TeamPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ogu.mutation.RemovedTeamIDs(); len(nodes) > 0 && !ogu.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oidcgroup.TeamTable,
			Columns: oidcgroup.TeamPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ogu.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oidcgroup.TeamTable,
			Columns: oidcgroup.TeamPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ogu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oidcgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ogu.mutation.done = true
	return n, nil
}

// OIDCGroupUpdateOne is the builder for updating a single OIDCGroup entity.
type OIDCGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OIDCGroupMutation
}

// SetName sets the "name" field.
func (oguo *OIDCGroupUpdateOne) SetName(s string) *OIDCGroupUpdateOne {
	oguo.mutation.SetName(s)
	return oguo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (oguo *OIDCGroupUpdateOne) SetNillableName(s *string) *OIDCGroupUpdateOne {
	if s != nil {
		oguo.SetName(*s)
	}
	return oguo
}

// AddMappedOidcGroupIDs adds the "mapped_oidc_groups" edge to the MappedOIDCGroup entity by IDs.
func (oguo *OIDCGroupUpdateOne) AddMappedOidcGroupIDs(ids ...int) *OIDCGroupUpdateOne {
	oguo.mutation.AddMappedOidcGroupIDs(ids...)
	return oguo
}

// AddMappedOidcGroups adds the "mapped_oidc_groups" edges to the MappedOIDCGroup entity.
func (oguo *OIDCGroupUpdateOne) AddMappedOidcGroups(m ...*MappedOIDCGroup) *OIDCGroupUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return oguo.AddMappedOidcGroupIDs(ids...)
}

// AddTeamIDs adds the "team" edge to the Team entity by IDs.
func (oguo *OIDCGroupUpdateOne) AddTeamIDs(ids ...int) *OIDCGroupUpdateOne {
	oguo.mutation.AddTeamIDs(ids...)
	return oguo
}

// AddTeam adds the "team" edges to the Team entity.
func (oguo *OIDCGroupUpdateOne) AddTeam(t ...*Team) *OIDCGroupUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return oguo.AddTeamIDs(ids...)
}

// Mutation returns the OIDCGroupMutation object of the builder.
func (oguo *OIDCGroupUpdateOne) Mutation() *OIDCGroupMutation {
	return oguo.mutation
}

// ClearMappedOidcGroups clears all "mapped_oidc_groups" edges to the MappedOIDCGroup entity.
func (oguo *OIDCGroupUpdateOne) ClearMappedOidcGroups() *OIDCGroupUpdateOne {
	oguo.mutation.ClearMappedOidcGroups()
	return oguo
}

// RemoveMappedOidcGroupIDs removes the "mapped_oidc_groups" edge to MappedOIDCGroup entities by IDs.
func (oguo *OIDCGroupUpdateOne) RemoveMappedOidcGroupIDs(ids ...int) *OIDCGroupUpdateOne {
	oguo.mutation.RemoveMappedOidcGroupIDs(ids...)
	return oguo
}

// RemoveMappedOidcGroups removes "mapped_oidc_groups" edges to MappedOIDCGroup entities.
func (oguo *OIDCGroupUpdateOne) RemoveMappedOidcGroups(m ...*MappedOIDCGroup) *OIDCGroupUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return oguo.RemoveMappedOidcGroupIDs(ids...)
}

// ClearTeam clears all "team" edges to the Team entity.
func (oguo *OIDCGroupUpdateOne) ClearTeam() *OIDCGroupUpdateOne {
	oguo.mutation.ClearTeam()
	return oguo
}

// RemoveTeamIDs removes the "team" edge to Team entities by IDs.
func (oguo *OIDCGroupUpdateOne) RemoveTeamIDs(ids ...int) *OIDCGroupUpdateOne {
	oguo.mutation.RemoveTeamIDs(ids...)
	return oguo
}

// RemoveTeam removes "team" edges to Team entities.
func (oguo *OIDCGroupUpdateOne) RemoveTeam(t ...*Team) *OIDCGroupUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return oguo.RemoveTeamIDs(ids...)
}

// Where appends a list predicates to the OIDCGroupUpdate builder.
func (oguo *OIDCGroupUpdateOne) Where(ps ...predicate.OIDCGroup) *OIDCGroupUpdateOne {
	oguo.mutation.Where(ps...)
	return oguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oguo *OIDCGroupUpdateOne) Select(field string, fields ...string) *OIDCGroupUpdateOne {
	oguo.fields = append([]string{field}, fields...)
	return oguo
}

// Save executes the query and returns the updated OIDCGroup entity.
func (oguo *OIDCGroupUpdateOne) Save(ctx context.Context) (*OIDCGroup, error) {
	return withHooks(ctx, oguo.sqlSave, oguo.mutation, oguo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oguo *OIDCGroupUpdateOne) SaveX(ctx context.Context) *OIDCGroup {
	node, err := oguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oguo *OIDCGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := oguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oguo *OIDCGroupUpdateOne) ExecX(ctx context.Context) {
	if err := oguo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (oguo *OIDCGroupUpdateOne) sqlSave(ctx context.Context) (_node *OIDCGroup, err error) {
	_spec := sqlgraph.NewUpdateSpec(oidcgroup.Table, oidcgroup.Columns, sqlgraph.NewFieldSpec(oidcgroup.FieldID, field.TypeInt))
	id, ok := oguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OIDCGroup.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oidcgroup.FieldID)
		for _, f := range fields {
			if !oidcgroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != oidcgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oguo.mutation.Name(); ok {
		_spec.SetField(oidcgroup.FieldName, field.TypeString, value)
	}
	if oguo.mutation.MappedOidcGroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   oidcgroup.MappedOidcGroupsTable,
			Columns: []string{oidcgroup.MappedOidcGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mappedoidcgroup.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oguo.mutation.RemovedMappedOidcGroupsIDs(); len(nodes) > 0 && !oguo.mutation.MappedOidcGroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   oidcgroup.MappedOidcGroupsTable,
			Columns: []string{oidcgroup.MappedOidcGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mappedoidcgroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oguo.mutation.MappedOidcGroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   oidcgroup.MappedOidcGroupsTable,
			Columns: []string{oidcgroup.MappedOidcGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mappedoidcgroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if oguo.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oidcgroup.TeamTable,
			Columns: oidcgroup.TeamPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oguo.mutation.RemovedTeamIDs(); len(nodes) > 0 && !oguo.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oidcgroup.TeamTable,
			Columns: oidcgroup.TeamPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oguo.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   oidcgroup.TeamTable,
			Columns: oidcgroup.TeamPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OIDCGroup{config: oguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oidcgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oguo.mutation.done = true
	return _node, nil
}
