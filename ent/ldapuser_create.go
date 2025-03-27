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

	"dependencytrack.io/hyades/ent/ldapuser"
	"dependencytrack.io/hyades/ent/permission"
	"dependencytrack.io/hyades/ent/team"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LDAPUserCreate is the builder for creating a LDAPUser entity.
type LDAPUserCreate struct {
	config
	mutation *LDAPUserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDn sets the "dn" field.
func (luc *LDAPUserCreate) SetDn(s string) *LDAPUserCreate {
	luc.mutation.SetDn(s)
	return luc
}

// SetEmail sets the "email" field.
func (luc *LDAPUserCreate) SetEmail(s string) *LDAPUserCreate {
	luc.mutation.SetEmail(s)
	return luc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (luc *LDAPUserCreate) SetNillableEmail(s *string) *LDAPUserCreate {
	if s != nil {
		luc.SetEmail(*s)
	}
	return luc
}

// SetUsername sets the "username" field.
func (luc *LDAPUserCreate) SetUsername(s string) *LDAPUserCreate {
	luc.mutation.SetUsername(s)
	return luc
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (luc *LDAPUserCreate) SetNillableUsername(s *string) *LDAPUserCreate {
	if s != nil {
		luc.SetUsername(*s)
	}
	return luc
}

// AddPermissionIDs adds the "permissions" edge to the Permission entity by IDs.
func (luc *LDAPUserCreate) AddPermissionIDs(ids ...int) *LDAPUserCreate {
	luc.mutation.AddPermissionIDs(ids...)
	return luc
}

// AddPermissions adds the "permissions" edges to the Permission entity.
func (luc *LDAPUserCreate) AddPermissions(p ...*Permission) *LDAPUserCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return luc.AddPermissionIDs(ids...)
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (luc *LDAPUserCreate) AddTeamIDs(ids ...int) *LDAPUserCreate {
	luc.mutation.AddTeamIDs(ids...)
	return luc
}

// AddTeams adds the "teams" edges to the Team entity.
func (luc *LDAPUserCreate) AddTeams(t ...*Team) *LDAPUserCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return luc.AddTeamIDs(ids...)
}

// Mutation returns the LDAPUserMutation object of the builder.
func (luc *LDAPUserCreate) Mutation() *LDAPUserMutation {
	return luc.mutation
}

// Save creates the LDAPUser in the database.
func (luc *LDAPUserCreate) Save(ctx context.Context) (*LDAPUser, error) {
	return withHooks(ctx, luc.sqlSave, luc.mutation, luc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (luc *LDAPUserCreate) SaveX(ctx context.Context) *LDAPUser {
	v, err := luc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (luc *LDAPUserCreate) Exec(ctx context.Context) error {
	_, err := luc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luc *LDAPUserCreate) ExecX(ctx context.Context) {
	if err := luc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luc *LDAPUserCreate) check() error {
	if _, ok := luc.mutation.Dn(); !ok {
		return &ValidationError{Name: "dn", err: errors.New(`ent: missing required field "LDAPUser.dn"`)}
	}
	return nil
}

func (luc *LDAPUserCreate) sqlSave(ctx context.Context) (*LDAPUser, error) {
	if err := luc.check(); err != nil {
		return nil, err
	}
	_node, _spec := luc.createSpec()
	if err := sqlgraph.CreateNode(ctx, luc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	luc.mutation.id = &_node.ID
	luc.mutation.done = true
	return _node, nil
}

func (luc *LDAPUserCreate) createSpec() (*LDAPUser, *sqlgraph.CreateSpec) {
	var (
		_node = &LDAPUser{config: luc.config}
		_spec = sqlgraph.NewCreateSpec(ldapuser.Table, sqlgraph.NewFieldSpec(ldapuser.FieldID, field.TypeInt))
	)
	_spec.OnConflict = luc.conflict
	if value, ok := luc.mutation.Dn(); ok {
		_spec.SetField(ldapuser.FieldDn, field.TypeString, value)
		_node.Dn = value
	}
	if value, ok := luc.mutation.Email(); ok {
		_spec.SetField(ldapuser.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := luc.mutation.Username(); ok {
		_spec.SetField(ldapuser.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if nodes := luc.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   ldapuser.PermissionsTable,
			Columns: ldapuser.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := luc.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   ldapuser.TeamsTable,
			Columns: ldapuser.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.LDAPUser.Create().
//		SetDn(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LDAPUserUpsert) {
//			SetDn(v+v).
//		}).
//		Exec(ctx)
func (luc *LDAPUserCreate) OnConflict(opts ...sql.ConflictOption) *LDAPUserUpsertOne {
	luc.conflict = opts
	return &LDAPUserUpsertOne{
		create: luc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.LDAPUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (luc *LDAPUserCreate) OnConflictColumns(columns ...string) *LDAPUserUpsertOne {
	luc.conflict = append(luc.conflict, sql.ConflictColumns(columns...))
	return &LDAPUserUpsertOne{
		create: luc,
	}
}

type (
	// LDAPUserUpsertOne is the builder for "upsert"-ing
	//  one LDAPUser node.
	LDAPUserUpsertOne struct {
		create *LDAPUserCreate
	}

	// LDAPUserUpsert is the "OnConflict" setter.
	LDAPUserUpsert struct {
		*sql.UpdateSet
	}
)

// SetDn sets the "dn" field.
func (u *LDAPUserUpsert) SetDn(v string) *LDAPUserUpsert {
	u.Set(ldapuser.FieldDn, v)
	return u
}

// UpdateDn sets the "dn" field to the value that was provided on create.
func (u *LDAPUserUpsert) UpdateDn() *LDAPUserUpsert {
	u.SetExcluded(ldapuser.FieldDn)
	return u
}

// SetEmail sets the "email" field.
func (u *LDAPUserUpsert) SetEmail(v string) *LDAPUserUpsert {
	u.Set(ldapuser.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *LDAPUserUpsert) UpdateEmail() *LDAPUserUpsert {
	u.SetExcluded(ldapuser.FieldEmail)
	return u
}

// ClearEmail clears the value of the "email" field.
func (u *LDAPUserUpsert) ClearEmail() *LDAPUserUpsert {
	u.SetNull(ldapuser.FieldEmail)
	return u
}

// SetUsername sets the "username" field.
func (u *LDAPUserUpsert) SetUsername(v string) *LDAPUserUpsert {
	u.Set(ldapuser.FieldUsername, v)
	return u
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *LDAPUserUpsert) UpdateUsername() *LDAPUserUpsert {
	u.SetExcluded(ldapuser.FieldUsername)
	return u
}

// ClearUsername clears the value of the "username" field.
func (u *LDAPUserUpsert) ClearUsername() *LDAPUserUpsert {
	u.SetNull(ldapuser.FieldUsername)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.LDAPUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *LDAPUserUpsertOne) UpdateNewValues() *LDAPUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.LDAPUser.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *LDAPUserUpsertOne) Ignore() *LDAPUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LDAPUserUpsertOne) DoNothing() *LDAPUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LDAPUserCreate.OnConflict
// documentation for more info.
func (u *LDAPUserUpsertOne) Update(set func(*LDAPUserUpsert)) *LDAPUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LDAPUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetDn sets the "dn" field.
func (u *LDAPUserUpsertOne) SetDn(v string) *LDAPUserUpsertOne {
	return u.Update(func(s *LDAPUserUpsert) {
		s.SetDn(v)
	})
}

// UpdateDn sets the "dn" field to the value that was provided on create.
func (u *LDAPUserUpsertOne) UpdateDn() *LDAPUserUpsertOne {
	return u.Update(func(s *LDAPUserUpsert) {
		s.UpdateDn()
	})
}

// SetEmail sets the "email" field.
func (u *LDAPUserUpsertOne) SetEmail(v string) *LDAPUserUpsertOne {
	return u.Update(func(s *LDAPUserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *LDAPUserUpsertOne) UpdateEmail() *LDAPUserUpsertOne {
	return u.Update(func(s *LDAPUserUpsert) {
		s.UpdateEmail()
	})
}

// ClearEmail clears the value of the "email" field.
func (u *LDAPUserUpsertOne) ClearEmail() *LDAPUserUpsertOne {
	return u.Update(func(s *LDAPUserUpsert) {
		s.ClearEmail()
	})
}

// SetUsername sets the "username" field.
func (u *LDAPUserUpsertOne) SetUsername(v string) *LDAPUserUpsertOne {
	return u.Update(func(s *LDAPUserUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *LDAPUserUpsertOne) UpdateUsername() *LDAPUserUpsertOne {
	return u.Update(func(s *LDAPUserUpsert) {
		s.UpdateUsername()
	})
}

// ClearUsername clears the value of the "username" field.
func (u *LDAPUserUpsertOne) ClearUsername() *LDAPUserUpsertOne {
	return u.Update(func(s *LDAPUserUpsert) {
		s.ClearUsername()
	})
}

// Exec executes the query.
func (u *LDAPUserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LDAPUserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LDAPUserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *LDAPUserUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *LDAPUserUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// LDAPUserCreateBulk is the builder for creating many LDAPUser entities in bulk.
type LDAPUserCreateBulk struct {
	config
	err      error
	builders []*LDAPUserCreate
	conflict []sql.ConflictOption
}

// Save creates the LDAPUser entities in the database.
func (lucb *LDAPUserCreateBulk) Save(ctx context.Context) ([]*LDAPUser, error) {
	if lucb.err != nil {
		return nil, lucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(lucb.builders))
	nodes := make([]*LDAPUser, len(lucb.builders))
	mutators := make([]Mutator, len(lucb.builders))
	for i := range lucb.builders {
		func(i int, root context.Context) {
			builder := lucb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LDAPUserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = lucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lucb *LDAPUserCreateBulk) SaveX(ctx context.Context) []*LDAPUser {
	v, err := lucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lucb *LDAPUserCreateBulk) Exec(ctx context.Context) error {
	_, err := lucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lucb *LDAPUserCreateBulk) ExecX(ctx context.Context) {
	if err := lucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.LDAPUser.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LDAPUserUpsert) {
//			SetDn(v+v).
//		}).
//		Exec(ctx)
func (lucb *LDAPUserCreateBulk) OnConflict(opts ...sql.ConflictOption) *LDAPUserUpsertBulk {
	lucb.conflict = opts
	return &LDAPUserUpsertBulk{
		create: lucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.LDAPUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lucb *LDAPUserCreateBulk) OnConflictColumns(columns ...string) *LDAPUserUpsertBulk {
	lucb.conflict = append(lucb.conflict, sql.ConflictColumns(columns...))
	return &LDAPUserUpsertBulk{
		create: lucb,
	}
}

// LDAPUserUpsertBulk is the builder for "upsert"-ing
// a bulk of LDAPUser nodes.
type LDAPUserUpsertBulk struct {
	create *LDAPUserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.LDAPUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *LDAPUserUpsertBulk) UpdateNewValues() *LDAPUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.LDAPUser.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *LDAPUserUpsertBulk) Ignore() *LDAPUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LDAPUserUpsertBulk) DoNothing() *LDAPUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LDAPUserCreateBulk.OnConflict
// documentation for more info.
func (u *LDAPUserUpsertBulk) Update(set func(*LDAPUserUpsert)) *LDAPUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LDAPUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetDn sets the "dn" field.
func (u *LDAPUserUpsertBulk) SetDn(v string) *LDAPUserUpsertBulk {
	return u.Update(func(s *LDAPUserUpsert) {
		s.SetDn(v)
	})
}

// UpdateDn sets the "dn" field to the value that was provided on create.
func (u *LDAPUserUpsertBulk) UpdateDn() *LDAPUserUpsertBulk {
	return u.Update(func(s *LDAPUserUpsert) {
		s.UpdateDn()
	})
}

// SetEmail sets the "email" field.
func (u *LDAPUserUpsertBulk) SetEmail(v string) *LDAPUserUpsertBulk {
	return u.Update(func(s *LDAPUserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *LDAPUserUpsertBulk) UpdateEmail() *LDAPUserUpsertBulk {
	return u.Update(func(s *LDAPUserUpsert) {
		s.UpdateEmail()
	})
}

// ClearEmail clears the value of the "email" field.
func (u *LDAPUserUpsertBulk) ClearEmail() *LDAPUserUpsertBulk {
	return u.Update(func(s *LDAPUserUpsert) {
		s.ClearEmail()
	})
}

// SetUsername sets the "username" field.
func (u *LDAPUserUpsertBulk) SetUsername(v string) *LDAPUserUpsertBulk {
	return u.Update(func(s *LDAPUserUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *LDAPUserUpsertBulk) UpdateUsername() *LDAPUserUpsertBulk {
	return u.Update(func(s *LDAPUserUpsert) {
		s.UpdateUsername()
	})
}

// ClearUsername clears the value of the "username" field.
func (u *LDAPUserUpsertBulk) ClearUsername() *LDAPUserUpsertBulk {
	return u.Update(func(s *LDAPUserUpsert) {
		s.ClearUsername()
	})
}

// Exec executes the query.
func (u *LDAPUserUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the LDAPUserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LDAPUserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LDAPUserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
