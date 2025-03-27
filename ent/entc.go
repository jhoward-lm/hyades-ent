//go:build ignore
// +build ignore

// --------------------------------------------------------------------
// This file is part of Dependency-Track.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0
// Copyright (c) OWASP Foundation. All Rights Reserved.
// --------------------------------------------------------------------

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/edge"
)

func main() {
	config := &gen.Config{
		Features: []gen.Feature{
			gen.FeatureExecQuery,
			gen.FeatureIntercept,
			gen.FeatureUpsert,
			gen.FeatureVersionedMigration,
		},
		Hooks:     []gen.Hook{marshalEdgesHook},
		Templates: []*gen.Template{gen.MustParse(gen.NewTemplate("").ParseDir("template"))},
	}

	// Ensure the generated directory (environment) avoids cyclic imports.
	if undo, err := gen.PrepareEnv(config); err != nil {
		defer undo()
		log.Fatalf("preparing ent environment: %v", err)
	}

	if err := entc.Generate("./schema", config); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

// marshalEdgesHook is a generator hook to set the struct tag
// for the generated types' Edges field to `json:"-"`.
func marshalEdgesHook(next gen.Generator) gen.Generator {
	return gen.GenerateFunc(func(graph *gen.Graph) error {
		tag := edge.Annotation{StructTag: `json:"-"`}

		for _, node := range graph.Nodes {
			setFieldStructTag(node.Fields...)
			setInverseEdgeStructTag(node.Edges...)
			node.Annotations.Set(tag.Name(), tag)
		}

		return next.Generate(graph)
	})
}

// setFieldStructTag sets the struct tag for generated types' edge-fields to `json:"-"`.
func setFieldStructTag(fields ...*gen.Field) {
	for idx := range fields {
		if fields[idx].IsEdgeField() {
			fields[idx].StructTag = `json:"-"`
		}
	}
}

// setInverseEdgeStructTag sets the struct tag for generated types' inverse edges to `json:"-"`.
func setInverseEdgeStructTag(edges ...*gen.Edge) {
	for idx := range edges {
		if edges[idx].IsInverse() {
			edges[idx].StructTag = `json:"-"`
		}
	}
}
