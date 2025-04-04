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
	"time"

	"dependencytrack.io/hyades/ent/installedupgrades"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// INSTALLEDUPGRADESCreate is the builder for creating a INSTALLEDUPGRADES entity.
type INSTALLEDUPGRADESCreate struct {
	config
	mutation *INSTALLEDUPGRADESMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetEndtime sets the "endtime" field.
func (ic *INSTALLEDUPGRADESCreate) SetEndtime(t time.Time) *INSTALLEDUPGRADESCreate {
	ic.mutation.SetEndtime(t)
	return ic
}

// SetNillableEndtime sets the "endtime" field if the given value is not nil.
func (ic *INSTALLEDUPGRADESCreate) SetNillableEndtime(t *time.Time) *INSTALLEDUPGRADESCreate {
	if t != nil {
		ic.SetEndtime(*t)
	}
	return ic
}

// SetStarttime sets the "starttime" field.
func (ic *INSTALLEDUPGRADESCreate) SetStarttime(t time.Time) *INSTALLEDUPGRADESCreate {
	ic.mutation.SetStarttime(t)
	return ic
}

// SetNillableStarttime sets the "starttime" field if the given value is not nil.
func (ic *INSTALLEDUPGRADESCreate) SetNillableStarttime(t *time.Time) *INSTALLEDUPGRADESCreate {
	if t != nil {
		ic.SetStarttime(*t)
	}
	return ic
}

// SetUpgradeclass sets the "upgradeclass" field.
func (ic *INSTALLEDUPGRADESCreate) SetUpgradeclass(s string) *INSTALLEDUPGRADESCreate {
	ic.mutation.SetUpgradeclass(s)
	return ic
}

// SetNillableUpgradeclass sets the "upgradeclass" field if the given value is not nil.
func (ic *INSTALLEDUPGRADESCreate) SetNillableUpgradeclass(s *string) *INSTALLEDUPGRADESCreate {
	if s != nil {
		ic.SetUpgradeclass(*s)
	}
	return ic
}

// Mutation returns the INSTALLEDUPGRADESMutation object of the builder.
func (ic *INSTALLEDUPGRADESCreate) Mutation() *INSTALLEDUPGRADESMutation {
	return ic.mutation
}

// Save creates the INSTALLEDUPGRADES in the database.
func (ic *INSTALLEDUPGRADESCreate) Save(ctx context.Context) (*INSTALLEDUPGRADES, error) {
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *INSTALLEDUPGRADESCreate) SaveX(ctx context.Context) *INSTALLEDUPGRADES {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *INSTALLEDUPGRADESCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *INSTALLEDUPGRADESCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *INSTALLEDUPGRADESCreate) check() error {
	return nil
}

func (ic *INSTALLEDUPGRADESCreate) sqlSave(ctx context.Context) (*INSTALLEDUPGRADES, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *INSTALLEDUPGRADESCreate) createSpec() (*INSTALLEDUPGRADES, *sqlgraph.CreateSpec) {
	var (
		_node = &INSTALLEDUPGRADES{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(installedupgrades.Table, sqlgraph.NewFieldSpec(installedupgrades.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ic.conflict
	if value, ok := ic.mutation.Endtime(); ok {
		_spec.SetField(installedupgrades.FieldEndtime, field.TypeTime, value)
		_node.Endtime = value
	}
	if value, ok := ic.mutation.Starttime(); ok {
		_spec.SetField(installedupgrades.FieldStarttime, field.TypeTime, value)
		_node.Starttime = value
	}
	if value, ok := ic.mutation.Upgradeclass(); ok {
		_spec.SetField(installedupgrades.FieldUpgradeclass, field.TypeString, value)
		_node.Upgradeclass = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.INSTALLEDUPGRADES.Create().
//		SetEndtime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.INSTALLEDUPGRADESUpsert) {
//			SetEndtime(v+v).
//		}).
//		Exec(ctx)
func (ic *INSTALLEDUPGRADESCreate) OnConflict(opts ...sql.ConflictOption) *INSTALLEDUPGRADESUpsertOne {
	ic.conflict = opts
	return &INSTALLEDUPGRADESUpsertOne{
		create: ic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.INSTALLEDUPGRADES.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ic *INSTALLEDUPGRADESCreate) OnConflictColumns(columns ...string) *INSTALLEDUPGRADESUpsertOne {
	ic.conflict = append(ic.conflict, sql.ConflictColumns(columns...))
	return &INSTALLEDUPGRADESUpsertOne{
		create: ic,
	}
}

type (
	// INSTALLEDUPGRADESUpsertOne is the builder for "upsert"-ing
	//  one INSTALLEDUPGRADES node.
	INSTALLEDUPGRADESUpsertOne struct {
		create *INSTALLEDUPGRADESCreate
	}

	// INSTALLEDUPGRADESUpsert is the "OnConflict" setter.
	INSTALLEDUPGRADESUpsert struct {
		*sql.UpdateSet
	}
)

// SetEndtime sets the "endtime" field.
func (u *INSTALLEDUPGRADESUpsert) SetEndtime(v time.Time) *INSTALLEDUPGRADESUpsert {
	u.Set(installedupgrades.FieldEndtime, v)
	return u
}

// UpdateEndtime sets the "endtime" field to the value that was provided on create.
func (u *INSTALLEDUPGRADESUpsert) UpdateEndtime() *INSTALLEDUPGRADESUpsert {
	u.SetExcluded(installedupgrades.FieldEndtime)
	return u
}

// ClearEndtime clears the value of the "endtime" field.
func (u *INSTALLEDUPGRADESUpsert) ClearEndtime() *INSTALLEDUPGRADESUpsert {
	u.SetNull(installedupgrades.FieldEndtime)
	return u
}

// SetStarttime sets the "starttime" field.
func (u *INSTALLEDUPGRADESUpsert) SetStarttime(v time.Time) *INSTALLEDUPGRADESUpsert {
	u.Set(installedupgrades.FieldStarttime, v)
	return u
}

// UpdateStarttime sets the "starttime" field to the value that was provided on create.
func (u *INSTALLEDUPGRADESUpsert) UpdateStarttime() *INSTALLEDUPGRADESUpsert {
	u.SetExcluded(installedupgrades.FieldStarttime)
	return u
}

// ClearStarttime clears the value of the "starttime" field.
func (u *INSTALLEDUPGRADESUpsert) ClearStarttime() *INSTALLEDUPGRADESUpsert {
	u.SetNull(installedupgrades.FieldStarttime)
	return u
}

// SetUpgradeclass sets the "upgradeclass" field.
func (u *INSTALLEDUPGRADESUpsert) SetUpgradeclass(v string) *INSTALLEDUPGRADESUpsert {
	u.Set(installedupgrades.FieldUpgradeclass, v)
	return u
}

// UpdateUpgradeclass sets the "upgradeclass" field to the value that was provided on create.
func (u *INSTALLEDUPGRADESUpsert) UpdateUpgradeclass() *INSTALLEDUPGRADESUpsert {
	u.SetExcluded(installedupgrades.FieldUpgradeclass)
	return u
}

// ClearUpgradeclass clears the value of the "upgradeclass" field.
func (u *INSTALLEDUPGRADESUpsert) ClearUpgradeclass() *INSTALLEDUPGRADESUpsert {
	u.SetNull(installedupgrades.FieldUpgradeclass)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.INSTALLEDUPGRADES.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *INSTALLEDUPGRADESUpsertOne) UpdateNewValues() *INSTALLEDUPGRADESUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.INSTALLEDUPGRADES.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *INSTALLEDUPGRADESUpsertOne) Ignore() *INSTALLEDUPGRADESUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *INSTALLEDUPGRADESUpsertOne) DoNothing() *INSTALLEDUPGRADESUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the INSTALLEDUPGRADESCreate.OnConflict
// documentation for more info.
func (u *INSTALLEDUPGRADESUpsertOne) Update(set func(*INSTALLEDUPGRADESUpsert)) *INSTALLEDUPGRADESUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&INSTALLEDUPGRADESUpsert{UpdateSet: update})
	}))
	return u
}

// SetEndtime sets the "endtime" field.
func (u *INSTALLEDUPGRADESUpsertOne) SetEndtime(v time.Time) *INSTALLEDUPGRADESUpsertOne {
	return u.Update(func(s *INSTALLEDUPGRADESUpsert) {
		s.SetEndtime(v)
	})
}

// UpdateEndtime sets the "endtime" field to the value that was provided on create.
func (u *INSTALLEDUPGRADESUpsertOne) UpdateEndtime() *INSTALLEDUPGRADESUpsertOne {
	return u.Update(func(s *INSTALLEDUPGRADESUpsert) {
		s.UpdateEndtime()
	})
}

// ClearEndtime clears the value of the "endtime" field.
func (u *INSTALLEDUPGRADESUpsertOne) ClearEndtime() *INSTALLEDUPGRADESUpsertOne {
	return u.Update(func(s *INSTALLEDUPGRADESUpsert) {
		s.ClearEndtime()
	})
}

// SetStarttime sets the "starttime" field.
func (u *INSTALLEDUPGRADESUpsertOne) SetStarttime(v time.Time) *INSTALLEDUPGRADESUpsertOne {
	return u.Update(func(s *INSTALLEDUPGRADESUpsert) {
		s.SetStarttime(v)
	})
}

// UpdateStarttime sets the "starttime" field to the value that was provided on create.
func (u *INSTALLEDUPGRADESUpsertOne) UpdateStarttime() *INSTALLEDUPGRADESUpsertOne {
	return u.Update(func(s *INSTALLEDUPGRADESUpsert) {
		s.UpdateStarttime()
	})
}

// ClearStarttime clears the value of the "starttime" field.
func (u *INSTALLEDUPGRADESUpsertOne) ClearStarttime() *INSTALLEDUPGRADESUpsertOne {
	return u.Update(func(s *INSTALLEDUPGRADESUpsert) {
		s.ClearStarttime()
	})
}

// SetUpgradeclass sets the "upgradeclass" field.
func (u *INSTALLEDUPGRADESUpsertOne) SetUpgradeclass(v string) *INSTALLEDUPGRADESUpsertOne {
	return u.Update(func(s *INSTALLEDUPGRADESUpsert) {
		s.SetUpgradeclass(v)
	})
}

// UpdateUpgradeclass sets the "upgradeclass" field to the value that was provided on create.
func (u *INSTALLEDUPGRADESUpsertOne) UpdateUpgradeclass() *INSTALLEDUPGRADESUpsertOne {
	return u.Update(func(s *INSTALLEDUPGRADESUpsert) {
		s.UpdateUpgradeclass()
	})
}

// ClearUpgradeclass clears the value of the "upgradeclass" field.
func (u *INSTALLEDUPGRADESUpsertOne) ClearUpgradeclass() *INSTALLEDUPGRADESUpsertOne {
	return u.Update(func(s *INSTALLEDUPGRADESUpsert) {
		s.ClearUpgradeclass()
	})
}

// Exec executes the query.
func (u *INSTALLEDUPGRADESUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for INSTALLEDUPGRADESCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *INSTALLEDUPGRADESUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *INSTALLEDUPGRADESUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *INSTALLEDUPGRADESUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// INSTALLEDUPGRADESCreateBulk is the builder for creating many INSTALLEDUPGRADES entities in bulk.
type INSTALLEDUPGRADESCreateBulk struct {
	config
	err      error
	builders []*INSTALLEDUPGRADESCreate
	conflict []sql.ConflictOption
}

// Save creates the INSTALLEDUPGRADES entities in the database.
func (icb *INSTALLEDUPGRADESCreateBulk) Save(ctx context.Context) ([]*INSTALLEDUPGRADES, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*INSTALLEDUPGRADES, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*INSTALLEDUPGRADESMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = icb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *INSTALLEDUPGRADESCreateBulk) SaveX(ctx context.Context) []*INSTALLEDUPGRADES {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *INSTALLEDUPGRADESCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *INSTALLEDUPGRADESCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.INSTALLEDUPGRADES.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.INSTALLEDUPGRADESUpsert) {
//			SetEndtime(v+v).
//		}).
//		Exec(ctx)
func (icb *INSTALLEDUPGRADESCreateBulk) OnConflict(opts ...sql.ConflictOption) *INSTALLEDUPGRADESUpsertBulk {
	icb.conflict = opts
	return &INSTALLEDUPGRADESUpsertBulk{
		create: icb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.INSTALLEDUPGRADES.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (icb *INSTALLEDUPGRADESCreateBulk) OnConflictColumns(columns ...string) *INSTALLEDUPGRADESUpsertBulk {
	icb.conflict = append(icb.conflict, sql.ConflictColumns(columns...))
	return &INSTALLEDUPGRADESUpsertBulk{
		create: icb,
	}
}

// INSTALLEDUPGRADESUpsertBulk is the builder for "upsert"-ing
// a bulk of INSTALLEDUPGRADES nodes.
type INSTALLEDUPGRADESUpsertBulk struct {
	create *INSTALLEDUPGRADESCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.INSTALLEDUPGRADES.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *INSTALLEDUPGRADESUpsertBulk) UpdateNewValues() *INSTALLEDUPGRADESUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.INSTALLEDUPGRADES.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *INSTALLEDUPGRADESUpsertBulk) Ignore() *INSTALLEDUPGRADESUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *INSTALLEDUPGRADESUpsertBulk) DoNothing() *INSTALLEDUPGRADESUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the INSTALLEDUPGRADESCreateBulk.OnConflict
// documentation for more info.
func (u *INSTALLEDUPGRADESUpsertBulk) Update(set func(*INSTALLEDUPGRADESUpsert)) *INSTALLEDUPGRADESUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&INSTALLEDUPGRADESUpsert{UpdateSet: update})
	}))
	return u
}

// SetEndtime sets the "endtime" field.
func (u *INSTALLEDUPGRADESUpsertBulk) SetEndtime(v time.Time) *INSTALLEDUPGRADESUpsertBulk {
	return u.Update(func(s *INSTALLEDUPGRADESUpsert) {
		s.SetEndtime(v)
	})
}

// UpdateEndtime sets the "endtime" field to the value that was provided on create.
func (u *INSTALLEDUPGRADESUpsertBulk) UpdateEndtime() *INSTALLEDUPGRADESUpsertBulk {
	return u.Update(func(s *INSTALLEDUPGRADESUpsert) {
		s.UpdateEndtime()
	})
}

// ClearEndtime clears the value of the "endtime" field.
func (u *INSTALLEDUPGRADESUpsertBulk) ClearEndtime() *INSTALLEDUPGRADESUpsertBulk {
	return u.Update(func(s *INSTALLEDUPGRADESUpsert) {
		s.ClearEndtime()
	})
}

// SetStarttime sets the "starttime" field.
func (u *INSTALLEDUPGRADESUpsertBulk) SetStarttime(v time.Time) *INSTALLEDUPGRADESUpsertBulk {
	return u.Update(func(s *INSTALLEDUPGRADESUpsert) {
		s.SetStarttime(v)
	})
}

// UpdateStarttime sets the "starttime" field to the value that was provided on create.
func (u *INSTALLEDUPGRADESUpsertBulk) UpdateStarttime() *INSTALLEDUPGRADESUpsertBulk {
	return u.Update(func(s *INSTALLEDUPGRADESUpsert) {
		s.UpdateStarttime()
	})
}

// ClearStarttime clears the value of the "starttime" field.
func (u *INSTALLEDUPGRADESUpsertBulk) ClearStarttime() *INSTALLEDUPGRADESUpsertBulk {
	return u.Update(func(s *INSTALLEDUPGRADESUpsert) {
		s.ClearStarttime()
	})
}

// SetUpgradeclass sets the "upgradeclass" field.
func (u *INSTALLEDUPGRADESUpsertBulk) SetUpgradeclass(v string) *INSTALLEDUPGRADESUpsertBulk {
	return u.Update(func(s *INSTALLEDUPGRADESUpsert) {
		s.SetUpgradeclass(v)
	})
}

// UpdateUpgradeclass sets the "upgradeclass" field to the value that was provided on create.
func (u *INSTALLEDUPGRADESUpsertBulk) UpdateUpgradeclass() *INSTALLEDUPGRADESUpsertBulk {
	return u.Update(func(s *INSTALLEDUPGRADESUpsert) {
		s.UpdateUpgradeclass()
	})
}

// ClearUpgradeclass clears the value of the "upgradeclass" field.
func (u *INSTALLEDUPGRADESUpsertBulk) ClearUpgradeclass() *INSTALLEDUPGRADESUpsertBulk {
	return u.Update(func(s *INSTALLEDUPGRADESUpsert) {
		s.ClearUpgradeclass()
	})
}

// Exec executes the query.
func (u *INSTALLEDUPGRADESUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the INSTALLEDUPGRADESCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for INSTALLEDUPGRADESCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *INSTALLEDUPGRADESUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
