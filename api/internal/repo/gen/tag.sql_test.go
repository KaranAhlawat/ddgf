package db

import (
	"database/sql"
	"fmt"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func (s *Suite) TestDeleteTag() {
	query := `-- name: DeleteTag :exec
	DELETE FROM "tags"
	WHERE "id" = $1
	`

	res := sqlmock.NewResult(0, 1)

	s.mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(s.tag.ID).
		WillReturnResult(res).
		WillReturnError(nil)

	err := s.querier.DeleteTag(s.ctx, s.tag.ID)

	require.NoError(s.T(), err)
}

func (s *Suite) TestInsertTag() {
	query := `-- name: InsertTag :one
	INSERT INTO "tags" ("id", "tag")
	VALUES ($1, $2)
	RETURNING id, tag
	`

	rows := s.mock.NewRows([]string{"id", "tag"}).
		AddRow(s.tag.ID, s.tag.Tag)

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(s.tag.ID, s.tag.Tag).
		WillReturnRows(rows)

	res, err := s.querier.InsertTag(s.ctx, InsertTagParams{
		ID:  s.tag.ID,
		Tag: s.tag.Tag,
	})

	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
	require.EqualValues(s.T(), *s.tag, res)
}

func (s *Suite) TestSelectTag() {
	query := `-- name: SelectTag :one
	SELECT id, tag
	FROM "tags"
	WHERE "id" = $1
	LIMIT 1
	`

	rows := s.mock.NewRows([]string{"id", "tag"}).
		AddRow(s.tag.ID, s.tag.Tag)

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(s.tag.ID).
		WillReturnRows(rows)

	res, err := s.querier.SelectTag(s.ctx, s.tag.ID)

	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
	require.EqualValues(s.T(), *s.tag, res)
}

func (s *Suite) TestSelectTags() {
	query := `-- name: SelectTags :many
	SELECT id, tag
	FROM "tags"
	ORDER BY "tag"
	`

	rows := s.mock.NewRows([]string{"id", "tag"}).
		AddRow(s.tag.ID, s.tag.Tag).
		AddRow(s.tag.ID, s.tag.Tag)

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows)

	res, err := s.querier.SelectTags(s.ctx)

	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)

	for _, dbTag := range res {
		require.EqualValues(s.T(), *s.tag, dbTag)
	}
}

func (s *Suite) TestTagNoRows() {
	query := `-- name: SelectTags :many
	SELECT id, tag
	FROM "tags"
	ORDER BY "tag"
	`

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnError(sql.ErrNoRows)

	_, err := s.querier.SelectTags(s.ctx)

	require.Error(s.T(), err)
}

func (s *Suite) TestTagRowError() {
	query := `-- name: SelectTags :many
	SELECT id, tag
	FROM "tags"
	ORDER BY "tag"
	`

	rows := sqlmock.NewRows([]string{"id", "tag"}).
		AddRow(s.tag.ID, s.tag.Tag).
		RowError(0, fmt.Errorf("row error"))

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows)

	res, err := s.querier.SelectTags(s.ctx)

	require.Error(s.T(), err)
	require.Nil(s.T(), res)
}
