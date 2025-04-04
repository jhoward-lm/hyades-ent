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

package eventservicelog

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the eventservicelog type in the database.
	Label = "event_service_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSubscriberclass holds the string denoting the subscriberclass field in the database.
	FieldSubscriberclass = "subscriberclass"
	// FieldStarted holds the string denoting the started field in the database.
	FieldStarted = "started"
	// FieldCompleted holds the string denoting the completed field in the database.
	FieldCompleted = "completed"
	// Table holds the table name of the eventservicelog in the database.
	Table = "EVENTSERVICELOG"
)

// Columns holds all SQL columns for eventservicelog fields.
var Columns = []string{
	FieldID,
	FieldSubscriberclass,
	FieldStarted,
	FieldCompleted,
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

// OrderOption defines the ordering options for the EventServiceLog queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySubscriberclass orders the results by the subscriberclass field.
func BySubscriberclass(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubscriberclass, opts...).ToFunc()
}

// ByStarted orders the results by the started field.
func ByStarted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStarted, opts...).ToFunc()
}

// ByCompleted orders the results by the completed field.
func ByCompleted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCompleted, opts...).ToFunc()
}
