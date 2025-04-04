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

package configproperty

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the configproperty type in the database.
	Label = "config_property"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldGroupname holds the string denoting the groupname field in the database.
	FieldGroupname = "groupname"
	// FieldPropertyname holds the string denoting the propertyname field in the database.
	FieldPropertyname = "propertyname"
	// FieldPropertytype holds the string denoting the propertytype field in the database.
	FieldPropertytype = "propertytype"
	// FieldPropertyvalue holds the string denoting the propertyvalue field in the database.
	FieldPropertyvalue = "propertyvalue"
	// Table holds the table name of the configproperty in the database.
	Table = "CONFIGPROPERTY"
)

// Columns holds all SQL columns for configproperty fields.
var Columns = []string{
	FieldID,
	FieldDescription,
	FieldGroupname,
	FieldPropertyname,
	FieldPropertytype,
	FieldPropertyvalue,
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

// OrderOption defines the ordering options for the ConfigProperty queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByGroupname orders the results by the groupname field.
func ByGroupname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGroupname, opts...).ToFunc()
}

// ByPropertyname orders the results by the propertyname field.
func ByPropertyname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPropertyname, opts...).ToFunc()
}

// ByPropertytype orders the results by the propertytype field.
func ByPropertytype(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPropertytype, opts...).ToFunc()
}

// ByPropertyvalue orders the results by the propertyvalue field.
func ByPropertyvalue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPropertyvalue, opts...).ToFunc()
}
