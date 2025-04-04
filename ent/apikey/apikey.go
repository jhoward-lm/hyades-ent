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
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the apikey type in the database.
	Label = "api_key"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldApikey holds the string denoting the apikey field in the database.
	FieldApikey = "apikey"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// FieldCreated holds the string denoting the created field in the database.
	FieldCreated = "created"
	// FieldLastUsed holds the string denoting the last_used field in the database.
	FieldLastUsed = "last_used"
	// FieldIsLegacy holds the string denoting the is_legacy field in the database.
	FieldIsLegacy = "is_legacy"
	// FieldPublicID holds the string denoting the public_id field in the database.
	FieldPublicID = "public_id"
	// EdgeTeams holds the string denoting the teams edge name in mutations.
	EdgeTeams = "teams"
	// Table holds the table name of the apikey in the database.
	Table = "APIKEY"
	// TeamsTable is the table that holds the teams relation/edge. The primary key declared below.
	TeamsTable = "team_api_keys"
	// TeamsInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	TeamsInverseTable = "TEAM"
)

// Columns holds all SQL columns for apikey fields.
var Columns = []string{
	FieldID,
	FieldApikey,
	FieldComment,
	FieldCreated,
	FieldLastUsed,
	FieldIsLegacy,
	FieldPublicID,
}

var (
	// TeamsPrimaryKey and TeamsColumn2 are the table columns denoting the
	// primary key for the teams relation (M2M).
	TeamsPrimaryKey = []string{"team_id", "api_key_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the APIKey queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByApikey orders the results by the apikey field.
func ByApikey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApikey, opts...).ToFunc()
}

// ByComment orders the results by the comment field.
func ByComment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComment, opts...).ToFunc()
}

// ByCreated orders the results by the created field.
func ByCreated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreated, opts...).ToFunc()
}

// ByLastUsed orders the results by the last_used field.
func ByLastUsed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastUsed, opts...).ToFunc()
}

// ByIsLegacy orders the results by the is_legacy field.
func ByIsLegacy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsLegacy, opts...).ToFunc()
}

// ByPublicID orders the results by the public_id field.
func ByPublicID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublicID, opts...).ToFunc()
}

// ByTeamsCount orders the results by teams count.
func ByTeamsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTeamsStep(), opts...)
	}
}

// ByTeams orders the results by teams terms.
func ByTeams(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTeamsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTeamsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TeamsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, TeamsTable, TeamsPrimaryKey...),
	)
}
