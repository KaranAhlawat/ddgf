// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: tag.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const deleteTag = `-- name: DeleteTag :exec
DELETE FROM "tags"
WHERE "id" = $1
`

func (q *Queries) DeleteTag(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTag, id)
	return err
}

const insertTag = `-- name: InsertTag :one
INSERT INTO "tags" ("id", "tag")
VALUES ($1, $2)
RETURNING id, tag
`

type InsertTagParams struct {
	ID  uuid.UUID `json:"id"`
	Tag string    `json:"tag"`
}

func (q *Queries) InsertTag(ctx context.Context, arg InsertTagParams) (Tag, error) {
	row := q.db.QueryRowContext(ctx, insertTag, arg.ID, arg.Tag)
	var i Tag
	err := row.Scan(&i.ID, &i.Tag)
	return i, err
}

const selectTag = `-- name: SelectTag :one
SELECT id, tag
FROM "tags"
WHERE "id" = $1
LIMIT 1
`

func (q *Queries) SelectTag(ctx context.Context, id uuid.UUID) (Tag, error) {
	row := q.db.QueryRowContext(ctx, selectTag, id)
	var i Tag
	err := row.Scan(&i.ID, &i.Tag)
	return i, err
}

const selectTags = `-- name: SelectTags :many
SELECT id, tag
FROM "tags"
ORDER BY "tag"
`

func (q *Queries) SelectTags(ctx context.Context) ([]Tag, error) {
	rows, err := q.db.QueryContext(ctx, selectTags)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tag{}
	for rows.Next() {
		var i Tag
		if err := rows.Scan(&i.ID, &i.Tag); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
