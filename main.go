// --------------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2025 The DependencyTrack Authors
// SPDX-FileName: main.go
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

package main

// #include <stdlib.h>
import "C"

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"slices"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	pgx "github.com/jackc/pgx/v5/stdlib"

	"dependencytrack.io/hyades/ent"
	"dependencytrack.io/hyades/ent/migrate"
)

const dsn = "postgresql://dtrack:dtrack@localhost:5432/dtrack"

type (
	Backend struct {
		client *ent.Client
		ctx    context.Context
	}

	TxFunc func(*ent.Tx) error
)

var errUninitializedClient = errors.New("backend client must be initialized")

func (backend *Backend) initClient() error {
	// Register the pgx driver as "postgres".
	if !slices.Contains(sql.Drivers(), dialect.Postgres) {
		sql.Register(dialect.Postgres, &pgx.Driver{})
	}

	client, err := ent.Open(dialect.Postgres, dsn)
	if err != nil {
		return fmt.Errorf("failed opening connection to sqlite: %w", err)
	}

	backend.client = client
	backend.ctx = ent.NewContext(context.Background(), client)

	// Run the auto migration tool.
	migrateOpts := []schema.MigrateOption{migrate.WithGlobalUniqueID(true), migrate.WithDropIndex(true)}
	if err := backend.client.Schema.Create(backend.ctx, migrateOpts...); err != nil {
		return fmt.Errorf("failed creating schema resources: %w", err)
	}

	return nil
}

func (backend *Backend) withTx(fns ...TxFunc) error {
	if backend.client == nil {
		return fmt.Errorf("%w", errUninitializedClient)
	}

	tx, err := backend.client.Tx(backend.ctx)
	if err != nil {
		return fmt.Errorf("creating transactional client: %w", err)
	}

	backend.ctx = ent.NewTxContext(backend.ctx, tx)

	for _, fn := range fns {
		if err := fn(tx); err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				err = fmt.Errorf("%w: rolling back transaction: %w", err, rollbackErr)
			}

			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}

//export CreateTeam
func CreateTeam(name *C.char) error {
	backend := &Backend{}
	if err := backend.initClient(); err != nil {
		return fmt.Errorf("initializing ent client: %w", err)
	}

	_, err := backend.client.Team.Create().
		SetName(C.GoString(name)).
		Save(backend.ctx)
	if err != nil {
		return fmt.Errorf("creating team: %w", err)
	}

	return nil
}

func main() {}
