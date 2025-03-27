// --------------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2025 The DependencyTrack Authors
// SPDX-FileName: ent/schema/team.go
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

	"dependencytrack.io/hyades/ent/schema/mixin"
)

type Team struct {
	ent.Schema
}

func (Team) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.UUID{},
	}
}

func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("api_keys", APIKey.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("permissions", Permission.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("oidc_groups", OIDCGroup.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("mapped_ldap_groups", MappedLDAPGroup.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("mapped_oidc_groups", MappedOIDCGroup.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("ldap_users", LDAPUser.Type).
			Ref("teams"),
		edge.From("managed_users", ManagedUser.Type).
			Ref("teams"),
		edge.From("oidc_users", OIDCUser.Type).
			Ref("teams"),
	}
}

func (Team) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("TEAM"),
	}
}
