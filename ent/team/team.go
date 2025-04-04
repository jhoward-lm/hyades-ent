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

package team

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the team type in the database.
	Label = "team"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeAPIKeys holds the string denoting the api_keys edge name in mutations.
	EdgeAPIKeys = "api_keys"
	// EdgePermissions holds the string denoting the permissions edge name in mutations.
	EdgePermissions = "permissions"
	// EdgeOidcGroups holds the string denoting the oidc_groups edge name in mutations.
	EdgeOidcGroups = "oidc_groups"
	// EdgeMappedLdapGroups holds the string denoting the mapped_ldap_groups edge name in mutations.
	EdgeMappedLdapGroups = "mapped_ldap_groups"
	// EdgeMappedOidcGroups holds the string denoting the mapped_oidc_groups edge name in mutations.
	EdgeMappedOidcGroups = "mapped_oidc_groups"
	// EdgeLdapUsers holds the string denoting the ldap_users edge name in mutations.
	EdgeLdapUsers = "ldap_users"
	// EdgeManagedUsers holds the string denoting the managed_users edge name in mutations.
	EdgeManagedUsers = "managed_users"
	// EdgeOidcUsers holds the string denoting the oidc_users edge name in mutations.
	EdgeOidcUsers = "oidc_users"
	// Table holds the table name of the team in the database.
	Table = "TEAM"
	// APIKeysTable is the table that holds the api_keys relation/edge. The primary key declared below.
	APIKeysTable = "team_api_keys"
	// APIKeysInverseTable is the table name for the APIKey entity.
	// It exists in this package in order to avoid circular dependency with the "apikey" package.
	APIKeysInverseTable = "APIKEY"
	// PermissionsTable is the table that holds the permissions relation/edge. The primary key declared below.
	PermissionsTable = "team_permissions"
	// PermissionsInverseTable is the table name for the Permission entity.
	// It exists in this package in order to avoid circular dependency with the "permission" package.
	PermissionsInverseTable = "PERMISSION"
	// OidcGroupsTable is the table that holds the oidc_groups relation/edge. The primary key declared below.
	OidcGroupsTable = "team_oidc_groups"
	// OidcGroupsInverseTable is the table name for the OIDCGroup entity.
	// It exists in this package in order to avoid circular dependency with the "oidcgroup" package.
	OidcGroupsInverseTable = "OIDCGROUP"
	// MappedLdapGroupsTable is the table that holds the mapped_ldap_groups relation/edge.
	MappedLdapGroupsTable = "MAPPEDLDAPGROUP"
	// MappedLdapGroupsInverseTable is the table name for the MappedLDAPGroup entity.
	// It exists in this package in order to avoid circular dependency with the "mappedldapgroup" package.
	MappedLdapGroupsInverseTable = "MAPPEDLDAPGROUP"
	// MappedLdapGroupsColumn is the table column denoting the mapped_ldap_groups relation/edge.
	MappedLdapGroupsColumn = "team_id"
	// MappedOidcGroupsTable is the table that holds the mapped_oidc_groups relation/edge.
	MappedOidcGroupsTable = "MAPPEDOIDCGROUP"
	// MappedOidcGroupsInverseTable is the table name for the MappedOIDCGroup entity.
	// It exists in this package in order to avoid circular dependency with the "mappedoidcgroup" package.
	MappedOidcGroupsInverseTable = "MAPPEDOIDCGROUP"
	// MappedOidcGroupsColumn is the table column denoting the mapped_oidc_groups relation/edge.
	MappedOidcGroupsColumn = "team_id"
	// LdapUsersTable is the table that holds the ldap_users relation/edge. The primary key declared below.
	LdapUsersTable = "ldap_user_teams"
	// LdapUsersInverseTable is the table name for the LDAPUser entity.
	// It exists in this package in order to avoid circular dependency with the "ldapuser" package.
	LdapUsersInverseTable = "LDAPUSER"
	// ManagedUsersTable is the table that holds the managed_users relation/edge. The primary key declared below.
	ManagedUsersTable = "managed_user_teams"
	// ManagedUsersInverseTable is the table name for the ManagedUser entity.
	// It exists in this package in order to avoid circular dependency with the "manageduser" package.
	ManagedUsersInverseTable = "MANAGEDUSER"
	// OidcUsersTable is the table that holds the oidc_users relation/edge. The primary key declared below.
	OidcUsersTable = "oidc_user_teams"
	// OidcUsersInverseTable is the table name for the OIDCUser entity.
	// It exists in this package in order to avoid circular dependency with the "oidcuser" package.
	OidcUsersInverseTable = "OIDCUSER"
)

// Columns holds all SQL columns for team fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldName,
}

var (
	// APIKeysPrimaryKey and APIKeysColumn2 are the table columns denoting the
	// primary key for the api_keys relation (M2M).
	APIKeysPrimaryKey = []string{"team_id", "api_key_id"}
	// PermissionsPrimaryKey and PermissionsColumn2 are the table columns denoting the
	// primary key for the permissions relation (M2M).
	PermissionsPrimaryKey = []string{"team_id", "permission_id"}
	// OidcGroupsPrimaryKey and OidcGroupsColumn2 are the table columns denoting the
	// primary key for the oidc_groups relation (M2M).
	OidcGroupsPrimaryKey = []string{"team_id", "oidc_group_id"}
	// LdapUsersPrimaryKey and LdapUsersColumn2 are the table columns denoting the
	// primary key for the ldap_users relation (M2M).
	LdapUsersPrimaryKey = []string{"ldap_user_id", "team_id"}
	// ManagedUsersPrimaryKey and ManagedUsersColumn2 are the table columns denoting the
	// primary key for the managed_users relation (M2M).
	ManagedUsersPrimaryKey = []string{"managed_user_id", "team_id"}
	// OidcUsersPrimaryKey and OidcUsersColumn2 are the table columns denoting the
	// primary key for the oidc_users relation (M2M).
	OidcUsersPrimaryKey = []string{"oidc_user_id", "team_id"}
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

// OrderOption defines the ordering options for the Team queries.
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

// ByAPIKeysCount orders the results by api_keys count.
func ByAPIKeysCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAPIKeysStep(), opts...)
	}
}

// ByAPIKeys orders the results by api_keys terms.
func ByAPIKeys(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAPIKeysStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPermissionsCount orders the results by permissions count.
func ByPermissionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPermissionsStep(), opts...)
	}
}

// ByPermissions orders the results by permissions terms.
func ByPermissions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPermissionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOidcGroupsCount orders the results by oidc_groups count.
func ByOidcGroupsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOidcGroupsStep(), opts...)
	}
}

// ByOidcGroups orders the results by oidc_groups terms.
func ByOidcGroups(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOidcGroupsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMappedLdapGroupsCount orders the results by mapped_ldap_groups count.
func ByMappedLdapGroupsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMappedLdapGroupsStep(), opts...)
	}
}

// ByMappedLdapGroups orders the results by mapped_ldap_groups terms.
func ByMappedLdapGroups(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMappedLdapGroupsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
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

// ByLdapUsersCount orders the results by ldap_users count.
func ByLdapUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLdapUsersStep(), opts...)
	}
}

// ByLdapUsers orders the results by ldap_users terms.
func ByLdapUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLdapUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByManagedUsersCount orders the results by managed_users count.
func ByManagedUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newManagedUsersStep(), opts...)
	}
}

// ByManagedUsers orders the results by managed_users terms.
func ByManagedUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newManagedUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOidcUsersCount orders the results by oidc_users count.
func ByOidcUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOidcUsersStep(), opts...)
	}
}

// ByOidcUsers orders the results by oidc_users terms.
func ByOidcUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOidcUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAPIKeysStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(APIKeysInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, APIKeysTable, APIKeysPrimaryKey...),
	)
}
func newPermissionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PermissionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PermissionsTable, PermissionsPrimaryKey...),
	)
}
func newOidcGroupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OidcGroupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, OidcGroupsTable, OidcGroupsPrimaryKey...),
	)
}
func newMappedLdapGroupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MappedLdapGroupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MappedLdapGroupsTable, MappedLdapGroupsColumn),
	)
}
func newMappedOidcGroupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MappedOidcGroupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MappedOidcGroupsTable, MappedOidcGroupsColumn),
	)
}
func newLdapUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LdapUsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, LdapUsersTable, LdapUsersPrimaryKey...),
	)
}
func newManagedUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ManagedUsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ManagedUsersTable, ManagedUsersPrimaryKey...),
	)
}
func newOidcUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OidcUsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, OidcUsersTable, OidcUsersPrimaryKey...),
	)
}
