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

package oidcgroup

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the oidcgroup type in the database.
	Label = "oidc_group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeMappedOidcGroups holds the string denoting the mapped_oidc_groups edge name in mutations.
	EdgeMappedOidcGroups = "mapped_oidc_groups"
	// EdgeTeam holds the string denoting the team edge name in mutations.
	EdgeTeam = "team"
	// Table holds the table name of the oidcgroup in the database.
	Table = "OIDCGROUP"
	// MappedOidcGroupsTable is the table that holds the mapped_oidc_groups relation/edge.
	MappedOidcGroupsTable = "MAPPEDOIDCGROUP"
	// MappedOidcGroupsInverseTable is the table name for the MappedOIDCGroup entity.
	// It exists in this package in order to avoid circular dependency with the "mappedoidcgroup" package.
	MappedOidcGroupsInverseTable = "MAPPEDOIDCGROUP"
	// MappedOidcGroupsColumn is the table column denoting the mapped_oidc_groups relation/edge.
	MappedOidcGroupsColumn = "group_id"
	// TeamTable is the table that holds the team relation/edge. The primary key declared below.
	TeamTable = "team_oidc_groups"
	// TeamInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	TeamInverseTable = "TEAM"
)

// Columns holds all SQL columns for oidcgroup fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldName,
}

var (
	// TeamPrimaryKey and TeamColumn2 are the table columns denoting the
	// primary key for the team relation (M2M).
	TeamPrimaryKey = []string{"team_id", "oidc_group_id"}
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

var (
	// DefaultUUID holds the default value on creation for the "uuid" field.
	DefaultUUID func() uuid.UUID
)

// OrderOption defines the ordering options for the OIDCGroup queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUUID orders the results by the uuid field.
func ByUUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUUID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByMappedOidcGroupsCount orders the results by mapped_oidc_groups count.
func ByMappedOidcGroupsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMappedOidcGroupsStep(), opts...)
	}
}

// ByMappedOidcGroups orders the results by mapped_oidc_groups terms.
func ByMappedOidcGroups(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMappedOidcGroupsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTeamCount orders the results by team count.
func ByTeamCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTeamStep(), opts...)
	}
}

// ByTeam orders the results by team terms.
func ByTeam(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTeamStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMappedOidcGroupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MappedOidcGroupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MappedOidcGroupsTable, MappedOidcGroupsColumn),
	)
}
func newTeamStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TeamInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, TeamTable, TeamPrimaryKey...),
	)
}
