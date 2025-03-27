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

package oidcuser

import (
	"dependencytrack.io/hyades/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldLTE(FieldID, id))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldEQ(FieldEmail, v))
}

// SubjectIdentifier applies equality check predicate on the "subject_identifier" field. It's identical to SubjectIdentifierEQ.
func SubjectIdentifier(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldEQ(FieldSubjectIdentifier, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldEQ(FieldUsername, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldContainsFold(FieldEmail, v))
}

// SubjectIdentifierEQ applies the EQ predicate on the "subject_identifier" field.
func SubjectIdentifierEQ(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldEQ(FieldSubjectIdentifier, v))
}

// SubjectIdentifierNEQ applies the NEQ predicate on the "subject_identifier" field.
func SubjectIdentifierNEQ(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldNEQ(FieldSubjectIdentifier, v))
}

// SubjectIdentifierIn applies the In predicate on the "subject_identifier" field.
func SubjectIdentifierIn(vs ...string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldIn(FieldSubjectIdentifier, vs...))
}

// SubjectIdentifierNotIn applies the NotIn predicate on the "subject_identifier" field.
func SubjectIdentifierNotIn(vs ...string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldNotIn(FieldSubjectIdentifier, vs...))
}

// SubjectIdentifierGT applies the GT predicate on the "subject_identifier" field.
func SubjectIdentifierGT(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldGT(FieldSubjectIdentifier, v))
}

// SubjectIdentifierGTE applies the GTE predicate on the "subject_identifier" field.
func SubjectIdentifierGTE(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldGTE(FieldSubjectIdentifier, v))
}

// SubjectIdentifierLT applies the LT predicate on the "subject_identifier" field.
func SubjectIdentifierLT(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldLT(FieldSubjectIdentifier, v))
}

// SubjectIdentifierLTE applies the LTE predicate on the "subject_identifier" field.
func SubjectIdentifierLTE(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldLTE(FieldSubjectIdentifier, v))
}

// SubjectIdentifierContains applies the Contains predicate on the "subject_identifier" field.
func SubjectIdentifierContains(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldContains(FieldSubjectIdentifier, v))
}

// SubjectIdentifierHasPrefix applies the HasPrefix predicate on the "subject_identifier" field.
func SubjectIdentifierHasPrefix(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldHasPrefix(FieldSubjectIdentifier, v))
}

// SubjectIdentifierHasSuffix applies the HasSuffix predicate on the "subject_identifier" field.
func SubjectIdentifierHasSuffix(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldHasSuffix(FieldSubjectIdentifier, v))
}

// SubjectIdentifierIsNil applies the IsNil predicate on the "subject_identifier" field.
func SubjectIdentifierIsNil() predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldIsNull(FieldSubjectIdentifier))
}

// SubjectIdentifierNotNil applies the NotNil predicate on the "subject_identifier" field.
func SubjectIdentifierNotNil() predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldNotNull(FieldSubjectIdentifier))
}

// SubjectIdentifierEqualFold applies the EqualFold predicate on the "subject_identifier" field.
func SubjectIdentifierEqualFold(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldEqualFold(FieldSubjectIdentifier, v))
}

// SubjectIdentifierContainsFold applies the ContainsFold predicate on the "subject_identifier" field.
func SubjectIdentifierContainsFold(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldContainsFold(FieldSubjectIdentifier, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.OIDCUser {
	return predicate.OIDCUser(sql.FieldContainsFold(FieldUsername, v))
}

// HasPermissions applies the HasEdge predicate on the "permissions" edge.
func HasPermissions() predicate.OIDCUser {
	return predicate.OIDCUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, PermissionsTable, PermissionsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPermissionsWith applies the HasEdge predicate on the "permissions" edge with a given conditions (other predicates).
func HasPermissionsWith(preds ...predicate.Permission) predicate.OIDCUser {
	return predicate.OIDCUser(func(s *sql.Selector) {
		step := newPermissionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTeams applies the HasEdge predicate on the "teams" edge.
func HasTeams() predicate.OIDCUser {
	return predicate.OIDCUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TeamsTable, TeamsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamsWith applies the HasEdge predicate on the "teams" edge with a given conditions (other predicates).
func HasTeamsWith(preds ...predicate.Team) predicate.OIDCUser {
	return predicate.OIDCUser(func(s *sql.Selector) {
		step := newTeamsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OIDCUser) predicate.OIDCUser {
	return predicate.OIDCUser(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OIDCUser) predicate.OIDCUser {
	return predicate.OIDCUser(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OIDCUser) predicate.OIDCUser {
	return predicate.OIDCUser(sql.NotPredicates(p))
}
