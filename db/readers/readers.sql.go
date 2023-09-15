// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: readers.sql

package readers

import (
	"context"
)

const get = `-- name: Get :one
SELECT value
FROM keys
WHERE name = ?1
`

func (q *Queries) Get(ctx context.Context, name string) (string, error) {
	row := q.queryRow(ctx, q.getStmt, get, name)
	var value string
	err := row.Scan(&value)
	return value, err
}

const substr = `-- name: Substr :one
SELECT SUBSTR(
  value,
  IIF(?1 < 0,
    ?1,
    ?1 + 1
  ),
  IIF(?2 < 0,
    LENGTH(value) - ?2,
    ?1 + ?2 + 1
  )
) FROM keys WHERE name = ?3
`

type SubstrParams struct {
	Start interface{}
	End   interface{}
	Name  string
}

func (q *Queries) Substr(ctx context.Context, arg SubstrParams) (string, error) {
	row := q.queryRow(ctx, q.substrStmt, substr, arg.Start, arg.End, arg.Name)
	var substr string
	err := row.Scan(&substr)
	return substr, err
}
