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

package installedupgrades

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the installedupgrades type in the database.
	Label = "installedupgrades"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEndtime holds the string denoting the endtime field in the database.
	FieldEndtime = "endtime"
	// FieldStarttime holds the string denoting the starttime field in the database.
	FieldStarttime = "starttime"
	// FieldUpgradeclass holds the string denoting the upgradeclass field in the database.
	FieldUpgradeclass = "upgradeclass"
	// Table holds the table name of the installedupgrades in the database.
	Table = "INSTALLEDUPGRADES"
)

// Columns holds all SQL columns for installedupgrades fields.
var Columns = []string{
	FieldID,
	FieldEndtime,
	FieldStarttime,
	FieldUpgradeclass,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the INSTALLEDUPGRADES queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEndtime orders the results by the endtime field.
func ByEndtime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndtime, opts...).ToFunc()
}

// ByStarttime orders the results by the starttime field.
func ByStarttime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStarttime, opts...).ToFunc()
}

// ByUpgradeclass orders the results by the upgradeclass field.
func ByUpgradeclass(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpgradeclass, opts...).ToFunc()
}
