// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: writers.sql

package writers

import (
	"context"
)

const delete = `-- name: Delete :exec
DELETE FROM strings WHERE name = ?1
`

func (q *Queries) Delete(ctx context.Context, name string) error {
	_, err := q.exec(ctx, q.deleteStmt, delete, name)
	return err
}

const flushAll = `-- name: FlushAll :exec
DELETE FROM strings
`

func (q *Queries) FlushAll(ctx context.Context) error {
	_, err := q.exec(ctx, q.flushAllStmt, flushAll)
	return err
}

const set = `-- name: Set :exec
INSERT INTO strings (name, value)
VALUES (?1, ?2) ON CONFLICT(name) DO
UPDATE
SET value = excluded.value
`

type SetParams struct {
	Name  string
	Value string
}

func (q *Queries) Set(ctx context.Context, arg SetParams) error {
	_, err := q.exec(ctx, q.setStmt, set, arg.Name, arg.Value)
	return err
}
