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

package schemaversion

import (
	"dependencytrack.io/hyades/ent/predicate"
	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldLTE(FieldID, id))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v string) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldEQ(FieldVersion, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v string) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v string) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...string) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...string) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v string) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v string) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v string) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v string) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldLTE(FieldVersion, v))
}

// VersionContains applies the Contains predicate on the "version" field.
func VersionContains(v string) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldContains(FieldVersion, v))
}

// VersionHasPrefix applies the HasPrefix predicate on the "version" field.
func VersionHasPrefix(v string) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldHasPrefix(FieldVersion, v))
}

// VersionHasSuffix applies the HasSuffix predicate on the "version" field.
func VersionHasSuffix(v string) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldHasSuffix(FieldVersion, v))
}

// VersionIsNil applies the IsNil predicate on the "version" field.
func VersionIsNil() predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldIsNull(FieldVersion))
}

// VersionNotNil applies the NotNil predicate on the "version" field.
func VersionNotNil() predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldNotNull(FieldVersion))
}

// VersionEqualFold applies the EqualFold predicate on the "version" field.
func VersionEqualFold(v string) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldEqualFold(FieldVersion, v))
}

// VersionContainsFold applies the ContainsFold predicate on the "version" field.
func VersionContainsFold(v string) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.FieldContainsFold(FieldVersion, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SchemaVersion) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SchemaVersion) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SchemaVersion) predicate.SchemaVersion {
	return predicate.SchemaVersion(sql.NotPredicates(p))
}
