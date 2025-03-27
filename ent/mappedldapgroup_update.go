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

	"dependencytrack.io/hyades/ent/mappedldapgroup"
	"dependencytrack.io/hyades/ent/predicate"
	"dependencytrack.io/hyades/ent/team"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MappedLDAPGroupUpdate is the builder for updating MappedLDAPGroup entities.
type MappedLDAPGroupUpdate struct {
	config
	hooks    []Hook
	mutation *MappedLDAPGroupMutation
}

// Where appends a list predicates to the MappedLDAPGroupUpdate builder.
func (mlgu *MappedLDAPGroupUpdate) Where(ps ...predicate.MappedLDAPGroup) *MappedLDAPGroupUpdate {
	mlgu.mutation.Where(ps...)
	return mlgu
}

// SetDn sets the "dn" field.
func (mlgu *MappedLDAPGroupUpdate) SetDn(s string) *MappedLDAPGroupUpdate {
	mlgu.mutation.SetDn(s)
	return mlgu
}

// SetNillableDn sets the "dn" field if the given value is not nil.
func (mlgu *MappedLDAPGroupUpdate) SetNillableDn(s *string) *MappedLDAPGroupUpdate {
	if s != nil {
		mlgu.SetDn(*s)
	}
	return mlgu
}

// SetTeamID sets the "team_id" field.
func (mlgu *MappedLDAPGroupUpdate) SetTeamID(i int) *MappedLDAPGroupUpdate {
	mlgu.mutation.SetTeamID(i)
	return mlgu
}

// SetNillableTeamID sets the "team_id" field if the given value is not nil.
func (mlgu *MappedLDAPGroupUpdate) SetNillableTeamID(i *int) *MappedLDAPGroupUpdate {
	if i != nil {
		mlgu.SetTeamID(*i)
	}
	return mlgu
}

// ClearTeamID clears the value of the "team_id" field.
func (mlgu *MappedLDAPGroupUpdate) ClearTeamID() *MappedLDAPGroupUpdate {
	mlgu.mutation.ClearTeamID()
	return mlgu
}

// SetTeam sets the "team" edge to the Team entity.
func (mlgu *MappedLDAPGroupUpdate) SetTeam(t *Team) *MappedLDAPGroupUpdate {
	return mlgu.SetTeamID(t.ID)
}

// Mutation returns the MappedLDAPGroupMutation object of the builder.
func (mlgu *MappedLDAPGroupUpdate) Mutation() *MappedLDAPGroupMutation {
	return mlgu.mutation
}

// ClearTeam clears the "team" edge to the Team entity.
func (mlgu *MappedLDAPGroupUpdate) ClearTeam() *MappedLDAPGroupUpdate {
	mlgu.mutation.ClearTeam()
	return mlgu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mlgu *MappedLDAPGroupUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mlgu.sqlSave, mlgu.mutation, mlgu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mlgu *MappedLDAPGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := mlgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mlgu *MappedLDAPGroupUpdate) Exec(ctx context.Context) error {
	_, err := mlgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mlgu *MappedLDAPGroupUpdate) ExecX(ctx context.Context) {
	if err := mlgu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mlgu *MappedLDAPGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(mappedldapgroup.Table, mappedldapgroup.Columns, sqlgraph.NewFieldSpec(mappedldapgroup.FieldID, field.TypeInt))
	if ps := mlgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mlgu.mutation.Dn(); ok {
		_spec.SetField(mappedldapgroup.FieldDn, field.TypeString, value)
	}
	if mlgu.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mappedldapgroup.TeamTable,
			Columns: []string{mappedldapgroup.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mlgu.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mappedldapgroup.TeamTable,
			Columns: []string{mappedldapgroup.TeamColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, mlgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mappedldapgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mlgu.mutation.done = true
	return n, nil
}

// MappedLDAPGroupUpdateOne is the builder for updating a single MappedLDAPGroup entity.
type MappedLDAPGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MappedLDAPGroupMutation
}

// SetDn sets the "dn" field.
func (mlguo *MappedLDAPGroupUpdateOne) SetDn(s string) *MappedLDAPGroupUpdateOne {
	mlguo.mutation.SetDn(s)
	return mlguo
}

// SetNillableDn sets the "dn" field if the given value is not nil.
func (mlguo *MappedLDAPGroupUpdateOne) SetNillableDn(s *string) *MappedLDAPGroupUpdateOne {
	if s != nil {
		mlguo.SetDn(*s)
	}
	return mlguo
}

// SetTeamID sets the "team_id" field.
func (mlguo *MappedLDAPGroupUpdateOne) SetTeamID(i int) *MappedLDAPGroupUpdateOne {
	mlguo.mutation.SetTeamID(i)
	return mlguo
}

// SetNillableTeamID sets the "team_id" field if the given value is not nil.
func (mlguo *MappedLDAPGroupUpdateOne) SetNillableTeamID(i *int) *MappedLDAPGroupUpdateOne {
	if i != nil {
		mlguo.SetTeamID(*i)
	}
	return mlguo
}

// ClearTeamID clears the value of the "team_id" field.
func (mlguo *MappedLDAPGroupUpdateOne) ClearTeamID() *MappedLDAPGroupUpdateOne {
	mlguo.mutation.ClearTeamID()
	return mlguo
}

// SetTeam sets the "team" edge to the Team entity.
func (mlguo *MappedLDAPGroupUpdateOne) SetTeam(t *Team) *MappedLDAPGroupUpdateOne {
	return mlguo.SetTeamID(t.ID)
}

// Mutation returns the MappedLDAPGroupMutation object of the builder.
func (mlguo *MappedLDAPGroupUpdateOne) Mutation() *MappedLDAPGroupMutation {
	return mlguo.mutation
}

// ClearTeam clears the "team" edge to the Team entity.
func (mlguo *MappedLDAPGroupUpdateOne) ClearTeam() *MappedLDAPGroupUpdateOne {
	mlguo.mutation.ClearTeam()
	return mlguo
}

// Where appends a list predicates to the MappedLDAPGroupUpdate builder.
func (mlguo *MappedLDAPGroupUpdateOne) Where(ps ...predicate.MappedLDAPGroup) *MappedLDAPGroupUpdateOne {
	mlguo.mutation.Where(ps...)
	return mlguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mlguo *MappedLDAPGroupUpdateOne) Select(field string, fields ...string) *MappedLDAPGroupUpdateOne {
	mlguo.fields = append([]string{field}, fields...)
	return mlguo
}

// Save executes the query and returns the updated MappedLDAPGroup entity.
func (mlguo *MappedLDAPGroupUpdateOne) Save(ctx context.Context) (*MappedLDAPGroup, error) {
	return withHooks(ctx, mlguo.sqlSave, mlguo.mutation, mlguo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mlguo *MappedLDAPGroupUpdateOne) SaveX(ctx context.Context) *MappedLDAPGroup {
	node, err := mlguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mlguo *MappedLDAPGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := mlguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mlguo *MappedLDAPGroupUpdateOne) ExecX(ctx context.Context) {
	if err := mlguo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mlguo *MappedLDAPGroupUpdateOne) sqlSave(ctx context.Context) (_node *MappedLDAPGroup, err error) {
	_spec := sqlgraph.NewUpdateSpec(mappedldapgroup.Table, mappedldapgroup.Columns, sqlgraph.NewFieldSpec(mappedldapgroup.FieldID, field.TypeInt))
	id, ok := mlguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MappedLDAPGroup.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mlguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mappedldapgroup.FieldID)
		for _, f := range fields {
			if !mappedldapgroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mappedldapgroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mlguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mlguo.mutation.Dn(); ok {
		_spec.SetField(mappedldapgroup.FieldDn, field.TypeString, value)
	}
	if mlguo.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mappedldapgroup.TeamTable,
			Columns: []string{mappedldapgroup.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mlguo.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mappedldapgroup.TeamTable,
			Columns: []string{mappedldapgroup.TeamColumn},
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
	_node = &MappedLDAPGroup{config: mlguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mlguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mappedldapgroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mlguo.mutation.done = true
	return _node, nil
}
