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

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"dependencytrack.io/hyades/ent/mappedoidcgroup"
	"dependencytrack.io/hyades/ent/oidcgroup"
	"dependencytrack.io/hyades/ent/team"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// MappedOIDCGroup is the model entity for the MappedOIDCGroup schema.
type MappedOIDCGroup struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID uuid.UUID `json:"uuid,omitempty"`
	// GroupID holds the value of the "group_id" field.
	GroupID int `json:"-"`
	// TeamID holds the value of the "team_id" field.
	TeamID int `json:"-"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MappedOIDCGroupQuery when eager-loading is set.
	Edges        MappedOIDCGroupEdges `json:"-"`
	selectValues sql.SelectValues
}

// MappedOIDCGroupEdges holds the relations/edges for other nodes in the graph.
type MappedOIDCGroupEdges struct {
	// OidcGroup holds the value of the oidc_group edge.
	OidcGroup *OIDCGroup `json:"-"`
	// Team holds the value of the team edge.
	Team *Team `json:"-"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OidcGroupOrErr returns the OidcGroup value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MappedOIDCGroupEdges) OidcGroupOrErr() (*OIDCGroup, error) {
	if e.OidcGroup != nil {
		return e.OidcGroup, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: oidcgroup.Label}
	}
	return nil, &NotLoadedError{edge: "oidc_group"}
}

// TeamOrErr returns the Team value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MappedOIDCGroupEdges) TeamOrErr() (*Team, error) {
	if e.Team != nil {
		return e.Team, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: team.Label}
	}
	return nil, &NotLoadedError{edge: "team"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MappedOIDCGroup) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case mappedoidcgroup.FieldID, mappedoidcgroup.FieldGroupID, mappedoidcgroup.FieldTeamID:
			values[i] = new(sql.NullInt64)
		case mappedoidcgroup.FieldUUID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MappedOIDCGroup fields.
func (mog *MappedOIDCGroup) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mappedoidcgroup.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mog.ID = int(value.Int64)
		case mappedoidcgroup.FieldUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value != nil {
				mog.UUID = *value
			}
		case mappedoidcgroup.FieldGroupID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				mog.GroupID = int(value.Int64)
			}
		case mappedoidcgroup.FieldTeamID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field team_id", values[i])
			} else if value.Valid {
				mog.TeamID = int(value.Int64)
			}
		default:
			mog.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MappedOIDCGroup.
// This includes values selected through modifiers, order, etc.
func (mog *MappedOIDCGroup) Value(name string) (ent.Value, error) {
	return mog.selectValues.Get(name)
}

// QueryOidcGroup queries the "oidc_group" edge of the MappedOIDCGroup entity.
func (mog *MappedOIDCGroup) QueryOidcGroup() *OIDCGroupQuery {
	return NewMappedOIDCGroupClient(mog.config).QueryOidcGroup(mog)
}

// QueryTeam queries the "team" edge of the MappedOIDCGroup entity.
func (mog *MappedOIDCGroup) QueryTeam() *TeamQuery {
	return NewMappedOIDCGroupClient(mog.config).QueryTeam(mog)
}

// Update returns a builder for updating this MappedOIDCGroup.
// Note that you need to call MappedOIDCGroup.Unwrap() before calling this method if this MappedOIDCGroup
// was returned from a transaction, and the transaction was committed or rolled back.
func (mog *MappedOIDCGroup) Update() *MappedOIDCGroupUpdateOne {
	return NewMappedOIDCGroupClient(mog.config).UpdateOne(mog)
}

// Unwrap unwraps the MappedOIDCGroup entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mog *MappedOIDCGroup) Unwrap() *MappedOIDCGroup {
	_tx, ok := mog.config.driver.(*txDriver)
	if !ok {
		panic("ent: MappedOIDCGroup is not a transactional entity")
	}
	mog.config.driver = _tx.drv
	return mog
}

// String implements the fmt.Stringer.
func (mog *MappedOIDCGroup) String() string {
	var builder strings.Builder
	builder.WriteString("MappedOIDCGroup(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mog.ID))
	builder.WriteString("uuid=")
	builder.WriteString(fmt.Sprintf("%v", mog.UUID))
	builder.WriteString(", ")
	builder.WriteString("group_id=")
	builder.WriteString(fmt.Sprintf("%v", mog.GroupID))
	builder.WriteString(", ")
	builder.WriteString("team_id=")
	builder.WriteString(fmt.Sprintf("%v", mog.TeamID))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (mog *MappedOIDCGroup) MarshalJSON() ([]byte, error) {
	type Alias MappedOIDCGroup
	return json.Marshal(&struct {
		*Alias
		MappedOIDCGroupEdges
	}{
		Alias:                (*Alias)(mog),
		MappedOIDCGroupEdges: mog.Edges,
	})
}

// MappedOIDCGroups is a parsable slice of MappedOIDCGroup.
type MappedOIDCGroups []*MappedOIDCGroup
