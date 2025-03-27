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

package apikey

import (
	"time"

	"dependencytrack.io/hyades/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.APIKey {
	return predicate.APIKey(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.APIKey {
	return predicate.APIKey(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.APIKey {
	return predicate.APIKey(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.APIKey {
	return predicate.APIKey(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.APIKey {
	return predicate.APIKey(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.APIKey {
	return predicate.APIKey(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.APIKey {
	return predicate.APIKey(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.APIKey {
	return predicate.APIKey(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.APIKey {
	return predicate.APIKey(sql.FieldLTE(FieldID, id))
}

// Apikey applies equality check predicate on the "apikey" field. It's identical to ApikeyEQ.
func Apikey(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldEQ(FieldApikey, v))
}

// Comment applies equality check predicate on the "comment" field. It's identical to CommentEQ.
func Comment(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldEQ(FieldComment, v))
}

// Created applies equality check predicate on the "created" field. It's identical to CreatedEQ.
func Created(v time.Time) predicate.APIKey {
	return predicate.APIKey(sql.FieldEQ(FieldCreated, v))
}

// LastUsed applies equality check predicate on the "last_used" field. It's identical to LastUsedEQ.
func LastUsed(v time.Time) predicate.APIKey {
	return predicate.APIKey(sql.FieldEQ(FieldLastUsed, v))
}

// IsLegacy applies equality check predicate on the "is_legacy" field. It's identical to IsLegacyEQ.
func IsLegacy(v bool) predicate.APIKey {
	return predicate.APIKey(sql.FieldEQ(FieldIsLegacy, v))
}

// PublicID applies equality check predicate on the "public_id" field. It's identical to PublicIDEQ.
func PublicID(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldEQ(FieldPublicID, v))
}

// ApikeyEQ applies the EQ predicate on the "apikey" field.
func ApikeyEQ(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldEQ(FieldApikey, v))
}

// ApikeyNEQ applies the NEQ predicate on the "apikey" field.
func ApikeyNEQ(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldNEQ(FieldApikey, v))
}

// ApikeyIn applies the In predicate on the "apikey" field.
func ApikeyIn(vs ...string) predicate.APIKey {
	return predicate.APIKey(sql.FieldIn(FieldApikey, vs...))
}

// ApikeyNotIn applies the NotIn predicate on the "apikey" field.
func ApikeyNotIn(vs ...string) predicate.APIKey {
	return predicate.APIKey(sql.FieldNotIn(FieldApikey, vs...))
}

// ApikeyGT applies the GT predicate on the "apikey" field.
func ApikeyGT(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldGT(FieldApikey, v))
}

// ApikeyGTE applies the GTE predicate on the "apikey" field.
func ApikeyGTE(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldGTE(FieldApikey, v))
}

// ApikeyLT applies the LT predicate on the "apikey" field.
func ApikeyLT(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldLT(FieldApikey, v))
}

// ApikeyLTE applies the LTE predicate on the "apikey" field.
func ApikeyLTE(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldLTE(FieldApikey, v))
}

// ApikeyContains applies the Contains predicate on the "apikey" field.
func ApikeyContains(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldContains(FieldApikey, v))
}

// ApikeyHasPrefix applies the HasPrefix predicate on the "apikey" field.
func ApikeyHasPrefix(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldHasPrefix(FieldApikey, v))
}

// ApikeyHasSuffix applies the HasSuffix predicate on the "apikey" field.
func ApikeyHasSuffix(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldHasSuffix(FieldApikey, v))
}

// ApikeyEqualFold applies the EqualFold predicate on the "apikey" field.
func ApikeyEqualFold(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldEqualFold(FieldApikey, v))
}

// ApikeyContainsFold applies the ContainsFold predicate on the "apikey" field.
func ApikeyContainsFold(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldContainsFold(FieldApikey, v))
}

// CommentEQ applies the EQ predicate on the "comment" field.
func CommentEQ(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldEQ(FieldComment, v))
}

// CommentNEQ applies the NEQ predicate on the "comment" field.
func CommentNEQ(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldNEQ(FieldComment, v))
}

// CommentIn applies the In predicate on the "comment" field.
func CommentIn(vs ...string) predicate.APIKey {
	return predicate.APIKey(sql.FieldIn(FieldComment, vs...))
}

// CommentNotIn applies the NotIn predicate on the "comment" field.
func CommentNotIn(vs ...string) predicate.APIKey {
	return predicate.APIKey(sql.FieldNotIn(FieldComment, vs...))
}

// CommentGT applies the GT predicate on the "comment" field.
func CommentGT(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldGT(FieldComment, v))
}

// CommentGTE applies the GTE predicate on the "comment" field.
func CommentGTE(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldGTE(FieldComment, v))
}

// CommentLT applies the LT predicate on the "comment" field.
func CommentLT(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldLT(FieldComment, v))
}

// CommentLTE applies the LTE predicate on the "comment" field.
func CommentLTE(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldLTE(FieldComment, v))
}

// CommentContains applies the Contains predicate on the "comment" field.
func CommentContains(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldContains(FieldComment, v))
}

// CommentHasPrefix applies the HasPrefix predicate on the "comment" field.
func CommentHasPrefix(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldHasPrefix(FieldComment, v))
}

// CommentHasSuffix applies the HasSuffix predicate on the "comment" field.
func CommentHasSuffix(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldHasSuffix(FieldComment, v))
}

// CommentIsNil applies the IsNil predicate on the "comment" field.
func CommentIsNil() predicate.APIKey {
	return predicate.APIKey(sql.FieldIsNull(FieldComment))
}

// CommentNotNil applies the NotNil predicate on the "comment" field.
func CommentNotNil() predicate.APIKey {
	return predicate.APIKey(sql.FieldNotNull(FieldComment))
}

// CommentEqualFold applies the EqualFold predicate on the "comment" field.
func CommentEqualFold(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldEqualFold(FieldComment, v))
}

// CommentContainsFold applies the ContainsFold predicate on the "comment" field.
func CommentContainsFold(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldContainsFold(FieldComment, v))
}

// CreatedEQ applies the EQ predicate on the "created" field.
func CreatedEQ(v time.Time) predicate.APIKey {
	return predicate.APIKey(sql.FieldEQ(FieldCreated, v))
}

// CreatedNEQ applies the NEQ predicate on the "created" field.
func CreatedNEQ(v time.Time) predicate.APIKey {
	return predicate.APIKey(sql.FieldNEQ(FieldCreated, v))
}

// CreatedIn applies the In predicate on the "created" field.
func CreatedIn(vs ...time.Time) predicate.APIKey {
	return predicate.APIKey(sql.FieldIn(FieldCreated, vs...))
}

// CreatedNotIn applies the NotIn predicate on the "created" field.
func CreatedNotIn(vs ...time.Time) predicate.APIKey {
	return predicate.APIKey(sql.FieldNotIn(FieldCreated, vs...))
}

// CreatedGT applies the GT predicate on the "created" field.
func CreatedGT(v time.Time) predicate.APIKey {
	return predicate.APIKey(sql.FieldGT(FieldCreated, v))
}

// CreatedGTE applies the GTE predicate on the "created" field.
func CreatedGTE(v time.Time) predicate.APIKey {
	return predicate.APIKey(sql.FieldGTE(FieldCreated, v))
}

// CreatedLT applies the LT predicate on the "created" field.
func CreatedLT(v time.Time) predicate.APIKey {
	return predicate.APIKey(sql.FieldLT(FieldCreated, v))
}

// CreatedLTE applies the LTE predicate on the "created" field.
func CreatedLTE(v time.Time) predicate.APIKey {
	return predicate.APIKey(sql.FieldLTE(FieldCreated, v))
}

// CreatedIsNil applies the IsNil predicate on the "created" field.
func CreatedIsNil() predicate.APIKey {
	return predicate.APIKey(sql.FieldIsNull(FieldCreated))
}

// CreatedNotNil applies the NotNil predicate on the "created" field.
func CreatedNotNil() predicate.APIKey {
	return predicate.APIKey(sql.FieldNotNull(FieldCreated))
}

// LastUsedEQ applies the EQ predicate on the "last_used" field.
func LastUsedEQ(v time.Time) predicate.APIKey {
	return predicate.APIKey(sql.FieldEQ(FieldLastUsed, v))
}

// LastUsedNEQ applies the NEQ predicate on the "last_used" field.
func LastUsedNEQ(v time.Time) predicate.APIKey {
	return predicate.APIKey(sql.FieldNEQ(FieldLastUsed, v))
}

// LastUsedIn applies the In predicate on the "last_used" field.
func LastUsedIn(vs ...time.Time) predicate.APIKey {
	return predicate.APIKey(sql.FieldIn(FieldLastUsed, vs...))
}

// LastUsedNotIn applies the NotIn predicate on the "last_used" field.
func LastUsedNotIn(vs ...time.Time) predicate.APIKey {
	return predicate.APIKey(sql.FieldNotIn(FieldLastUsed, vs...))
}

// LastUsedGT applies the GT predicate on the "last_used" field.
func LastUsedGT(v time.Time) predicate.APIKey {
	return predicate.APIKey(sql.FieldGT(FieldLastUsed, v))
}

// LastUsedGTE applies the GTE predicate on the "last_used" field.
func LastUsedGTE(v time.Time) predicate.APIKey {
	return predicate.APIKey(sql.FieldGTE(FieldLastUsed, v))
}

// LastUsedLT applies the LT predicate on the "last_used" field.
func LastUsedLT(v time.Time) predicate.APIKey {
	return predicate.APIKey(sql.FieldLT(FieldLastUsed, v))
}

// LastUsedLTE applies the LTE predicate on the "last_used" field.
func LastUsedLTE(v time.Time) predicate.APIKey {
	return predicate.APIKey(sql.FieldLTE(FieldLastUsed, v))
}

// LastUsedIsNil applies the IsNil predicate on the "last_used" field.
func LastUsedIsNil() predicate.APIKey {
	return predicate.APIKey(sql.FieldIsNull(FieldLastUsed))
}

// LastUsedNotNil applies the NotNil predicate on the "last_used" field.
func LastUsedNotNil() predicate.APIKey {
	return predicate.APIKey(sql.FieldNotNull(FieldLastUsed))
}

// IsLegacyEQ applies the EQ predicate on the "is_legacy" field.
func IsLegacyEQ(v bool) predicate.APIKey {
	return predicate.APIKey(sql.FieldEQ(FieldIsLegacy, v))
}

// IsLegacyNEQ applies the NEQ predicate on the "is_legacy" field.
func IsLegacyNEQ(v bool) predicate.APIKey {
	return predicate.APIKey(sql.FieldNEQ(FieldIsLegacy, v))
}

// PublicIDEQ applies the EQ predicate on the "public_id" field.
func PublicIDEQ(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldEQ(FieldPublicID, v))
}

// PublicIDNEQ applies the NEQ predicate on the "public_id" field.
func PublicIDNEQ(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldNEQ(FieldPublicID, v))
}

// PublicIDIn applies the In predicate on the "public_id" field.
func PublicIDIn(vs ...string) predicate.APIKey {
	return predicate.APIKey(sql.FieldIn(FieldPublicID, vs...))
}

// PublicIDNotIn applies the NotIn predicate on the "public_id" field.
func PublicIDNotIn(vs ...string) predicate.APIKey {
	return predicate.APIKey(sql.FieldNotIn(FieldPublicID, vs...))
}

// PublicIDGT applies the GT predicate on the "public_id" field.
func PublicIDGT(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldGT(FieldPublicID, v))
}

// PublicIDGTE applies the GTE predicate on the "public_id" field.
func PublicIDGTE(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldGTE(FieldPublicID, v))
}

// PublicIDLT applies the LT predicate on the "public_id" field.
func PublicIDLT(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldLT(FieldPublicID, v))
}

// PublicIDLTE applies the LTE predicate on the "public_id" field.
func PublicIDLTE(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldLTE(FieldPublicID, v))
}

// PublicIDContains applies the Contains predicate on the "public_id" field.
func PublicIDContains(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldContains(FieldPublicID, v))
}

// PublicIDHasPrefix applies the HasPrefix predicate on the "public_id" field.
func PublicIDHasPrefix(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldHasPrefix(FieldPublicID, v))
}

// PublicIDHasSuffix applies the HasSuffix predicate on the "public_id" field.
func PublicIDHasSuffix(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldHasSuffix(FieldPublicID, v))
}

// PublicIDIsNil applies the IsNil predicate on the "public_id" field.
func PublicIDIsNil() predicate.APIKey {
	return predicate.APIKey(sql.FieldIsNull(FieldPublicID))
}

// PublicIDNotNil applies the NotNil predicate on the "public_id" field.
func PublicIDNotNil() predicate.APIKey {
	return predicate.APIKey(sql.FieldNotNull(FieldPublicID))
}

// PublicIDEqualFold applies the EqualFold predicate on the "public_id" field.
func PublicIDEqualFold(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldEqualFold(FieldPublicID, v))
}

// PublicIDContainsFold applies the ContainsFold predicate on the "public_id" field.
func PublicIDContainsFold(v string) predicate.APIKey {
	return predicate.APIKey(sql.FieldContainsFold(FieldPublicID, v))
}

// HasTeams applies the HasEdge predicate on the "teams" edge.
func HasTeams() predicate.APIKey {
	return predicate.APIKey(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, TeamsTable, TeamsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamsWith applies the HasEdge predicate on the "teams" edge with a given conditions (other predicates).
func HasTeamsWith(preds ...predicate.Team) predicate.APIKey {
	return predicate.APIKey(func(s *sql.Selector) {
		step := newTeamsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.APIKey) predicate.APIKey {
	return predicate.APIKey(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.APIKey) predicate.APIKey {
	return predicate.APIKey(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.APIKey) predicate.APIKey {
	return predicate.APIKey(sql.NotPredicates(p))
}
