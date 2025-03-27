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

package manageduser

import (
	"time"

	"dependencytrack.io/hyades/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldLTE(FieldID, id))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEQ(FieldEmail, v))
}

// ForcePasswordChange applies equality check predicate on the "force_password_change" field. It's identical to ForcePasswordChangeEQ.
func ForcePasswordChange(v bool) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEQ(FieldForcePasswordChange, v))
}

// Fullname applies equality check predicate on the "fullname" field. It's identical to FullnameEQ.
func Fullname(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEQ(FieldFullname, v))
}

// LastPasswordChange applies equality check predicate on the "last_password_change" field. It's identical to LastPasswordChangeEQ.
func LastPasswordChange(v time.Time) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEQ(FieldLastPasswordChange, v))
}

// NonExpiryPassword applies equality check predicate on the "non_expiry_password" field. It's identical to NonExpiryPasswordEQ.
func NonExpiryPassword(v bool) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEQ(FieldNonExpiryPassword, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEQ(FieldPassword, v))
}

// Suspended applies equality check predicate on the "suspended" field. It's identical to SuspendedEQ.
func Suspended(v bool) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEQ(FieldSuspended, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEQ(FieldUsername, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldContainsFold(FieldEmail, v))
}

// ForcePasswordChangeEQ applies the EQ predicate on the "force_password_change" field.
func ForcePasswordChangeEQ(v bool) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEQ(FieldForcePasswordChange, v))
}

// ForcePasswordChangeNEQ applies the NEQ predicate on the "force_password_change" field.
func ForcePasswordChangeNEQ(v bool) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldNEQ(FieldForcePasswordChange, v))
}

// FullnameEQ applies the EQ predicate on the "fullname" field.
func FullnameEQ(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEQ(FieldFullname, v))
}

// FullnameNEQ applies the NEQ predicate on the "fullname" field.
func FullnameNEQ(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldNEQ(FieldFullname, v))
}

// FullnameIn applies the In predicate on the "fullname" field.
func FullnameIn(vs ...string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldIn(FieldFullname, vs...))
}

// FullnameNotIn applies the NotIn predicate on the "fullname" field.
func FullnameNotIn(vs ...string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldNotIn(FieldFullname, vs...))
}

// FullnameGT applies the GT predicate on the "fullname" field.
func FullnameGT(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldGT(FieldFullname, v))
}

// FullnameGTE applies the GTE predicate on the "fullname" field.
func FullnameGTE(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldGTE(FieldFullname, v))
}

// FullnameLT applies the LT predicate on the "fullname" field.
func FullnameLT(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldLT(FieldFullname, v))
}

// FullnameLTE applies the LTE predicate on the "fullname" field.
func FullnameLTE(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldLTE(FieldFullname, v))
}

// FullnameContains applies the Contains predicate on the "fullname" field.
func FullnameContains(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldContains(FieldFullname, v))
}

// FullnameHasPrefix applies the HasPrefix predicate on the "fullname" field.
func FullnameHasPrefix(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldHasPrefix(FieldFullname, v))
}

// FullnameHasSuffix applies the HasSuffix predicate on the "fullname" field.
func FullnameHasSuffix(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldHasSuffix(FieldFullname, v))
}

// FullnameIsNil applies the IsNil predicate on the "fullname" field.
func FullnameIsNil() predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldIsNull(FieldFullname))
}

// FullnameNotNil applies the NotNil predicate on the "fullname" field.
func FullnameNotNil() predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldNotNull(FieldFullname))
}

// FullnameEqualFold applies the EqualFold predicate on the "fullname" field.
func FullnameEqualFold(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEqualFold(FieldFullname, v))
}

// FullnameContainsFold applies the ContainsFold predicate on the "fullname" field.
func FullnameContainsFold(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldContainsFold(FieldFullname, v))
}

// LastPasswordChangeEQ applies the EQ predicate on the "last_password_change" field.
func LastPasswordChangeEQ(v time.Time) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEQ(FieldLastPasswordChange, v))
}

// LastPasswordChangeNEQ applies the NEQ predicate on the "last_password_change" field.
func LastPasswordChangeNEQ(v time.Time) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldNEQ(FieldLastPasswordChange, v))
}

// LastPasswordChangeIn applies the In predicate on the "last_password_change" field.
func LastPasswordChangeIn(vs ...time.Time) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldIn(FieldLastPasswordChange, vs...))
}

// LastPasswordChangeNotIn applies the NotIn predicate on the "last_password_change" field.
func LastPasswordChangeNotIn(vs ...time.Time) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldNotIn(FieldLastPasswordChange, vs...))
}

// LastPasswordChangeGT applies the GT predicate on the "last_password_change" field.
func LastPasswordChangeGT(v time.Time) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldGT(FieldLastPasswordChange, v))
}

// LastPasswordChangeGTE applies the GTE predicate on the "last_password_change" field.
func LastPasswordChangeGTE(v time.Time) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldGTE(FieldLastPasswordChange, v))
}

// LastPasswordChangeLT applies the LT predicate on the "last_password_change" field.
func LastPasswordChangeLT(v time.Time) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldLT(FieldLastPasswordChange, v))
}

// LastPasswordChangeLTE applies the LTE predicate on the "last_password_change" field.
func LastPasswordChangeLTE(v time.Time) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldLTE(FieldLastPasswordChange, v))
}

// NonExpiryPasswordEQ applies the EQ predicate on the "non_expiry_password" field.
func NonExpiryPasswordEQ(v bool) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEQ(FieldNonExpiryPassword, v))
}

// NonExpiryPasswordNEQ applies the NEQ predicate on the "non_expiry_password" field.
func NonExpiryPasswordNEQ(v bool) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldNEQ(FieldNonExpiryPassword, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldContainsFold(FieldPassword, v))
}

// SuspendedEQ applies the EQ predicate on the "suspended" field.
func SuspendedEQ(v bool) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEQ(FieldSuspended, v))
}

// SuspendedNEQ applies the NEQ predicate on the "suspended" field.
func SuspendedNEQ(v bool) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldNEQ(FieldSuspended, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameIsNil applies the IsNil predicate on the "username" field.
func UsernameIsNil() predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldIsNull(FieldUsername))
}

// UsernameNotNil applies the NotNil predicate on the "username" field.
func UsernameNotNil() predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldNotNull(FieldUsername))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.ManagedUser {
	return predicate.ManagedUser(sql.FieldContainsFold(FieldUsername, v))
}

// HasPermissions applies the HasEdge predicate on the "permissions" edge.
func HasPermissions() predicate.ManagedUser {
	return predicate.ManagedUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, PermissionsTable, PermissionsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPermissionsWith applies the HasEdge predicate on the "permissions" edge with a given conditions (other predicates).
func HasPermissionsWith(preds ...predicate.Permission) predicate.ManagedUser {
	return predicate.ManagedUser(func(s *sql.Selector) {
		step := newPermissionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTeams applies the HasEdge predicate on the "teams" edge.
func HasTeams() predicate.ManagedUser {
	return predicate.ManagedUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TeamsTable, TeamsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamsWith applies the HasEdge predicate on the "teams" edge with a given conditions (other predicates).
func HasTeamsWith(preds ...predicate.Team) predicate.ManagedUser {
	return predicate.ManagedUser(func(s *sql.Selector) {
		step := newTeamsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ManagedUser) predicate.ManagedUser {
	return predicate.ManagedUser(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ManagedUser) predicate.ManagedUser {
	return predicate.ManagedUser(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ManagedUser) predicate.ManagedUser {
	return predicate.ManagedUser(sql.NotPredicates(p))
}
