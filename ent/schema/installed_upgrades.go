// --------------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2025 The DependencyTrack Authors
// SPDX-FileName: ent/schema/installed_upgrades.go
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

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type INSTALLEDUPGRADES struct {
	ent.Schema
}

func (INSTALLEDUPGRADES) Fields() []ent.Field {
	return []ent.Field{
		field.Time("endtime").Optional(),
		field.Time("starttime").Optional(),
		field.String("upgradeclass").Optional(),
	}
}

func (INSTALLEDUPGRADES) Edges() []ent.Edge {
	return nil
}

func (INSTALLEDUPGRADES) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("INSTALLEDUPGRADES"),
	}
}
