// --------------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2025 The DependencyTrack Authors
// SPDX-FileName: ent/schema/api_key.go
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
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type APIKey struct {
	ent.Schema
}

func (APIKey) Fields() []ent.Field {
	return []ent.Field{
		field.String("apikey").Unique(),
		field.String("comment").Optional(),
		field.Time("created").Optional(),
		field.Time("last_used").Optional(),
		field.Bool("is_legacy"),
		field.String("public_id").Optional().Unique(),
	}
}

func (APIKey) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teams", Team.Type).
			Ref("api_keys"),
	}
}

func (APIKey) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("APIKEY"),
	}
}
