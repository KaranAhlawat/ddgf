package db

import (
	"database/sql"
	"fmt"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func (s *Suite) TestDeleteAdvice() {
	query := `-- name: DeleteAdvice :exec
	DELETE FROM "advices"
	WHERE "id" = $1
	`

	res := sqlmock.NewResult(0, 1)

	s.mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(s.advice.ID).
		WillReturnResult(res).
		WillReturnError(nil)

	err := s.querier.DeleteAdvice(s.ctx, s.advice.ID)

	require.NoError(s.T(), err)
}

func (s *Suite) TestInsertAdvice() {
	query := `-- name: InsertAdvice :one
	INSERT INTO "advices" ("id", "content")
	VALUES ($1, $2)
	RETURNING id, content
	`

	rows := sqlmock.NewRows([]string{"id", "content"}).
		AddRow(s.advice.ID, s.advice.Content)

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(s.advice.ID, s.advice.Content).
		WillReturnRows(rows)

	res, err := s.querier.InsertAdvice(s.ctx, InsertAdviceParams{
		ID:      s.advice.ID,
		Content: s.advice.Content,
	})

	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
	require.EqualValues(s.T(), *s.advice, res)
}

func (s *Suite) TestSelectAdvice() {
	query := `-- name: SelectAdvice :one
	SELECT id, content
	FROM "advices"
	WHERE "id" = $1
	LIMIT 1
	`

	rows := sqlmock.NewRows([]string{"id", "content"}).
		AddRow(s.advice.ID, s.advice.Content)

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(s.advice.ID).
		WillReturnRows(rows)

	res, err := s.querier.SelectAdvice(s.ctx, s.advice.ID)

	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
	require.EqualValues(s.T(), *s.advice, res)
}

func (s *Suite) TestSelectAdvices() {
	query := `-- name: SelectAdvices :many
	SELECT id, content
	FROM "advices"
	ORDER BY "id"
	`

	rows := sqlmock.NewRows([]string{"id", "content"}).
		AddRow(s.advice.ID, s.advice.Content).
		AddRow(s.advice.ID, s.advice.Content)

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows)

	res, err := s.querier.SelectAdvices(s.ctx)

	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)

	for _, advice := range res {
		require.EqualValues(s.T(), *s.advice, advice)
	}
}

func (s *Suite) TestAdviceNoRows() {
	query := `-- name: SelectAdvices :many
	SELECT id, content
	FROM "advices"
	ORDER BY "id"
	`

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnError(sql.ErrNoRows)

	_, err := s.querier.SelectAdvices(s.ctx)

	require.Error(s.T(), err)
}

func (s *Suite) TestAdviceRowError() {
	query := `-- name: SelectAdvices :many
	SELECT id, content
	FROM "advices"
	ORDER BY "id"
	`

	rows := sqlmock.NewRows([]string{"id", "content"}).
		AddRow(s.advice.ID, s.advice.Content).
		RowError(0, fmt.Errorf("row error"))

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows)

	res, err := s.querier.SelectAdvices(s.ctx)

	require.Error(s.T(), err)
	require.Nil(s.T(), res)
}

func (s *Suite) TestUpdateAdvice() {
	query := `-- name: UpdateAdvice :exec
	UPDATE "advices"
	SET "content" = $1
	WHERE "id" = $2
	`

	res := sqlmock.NewResult(0, 1)

	s.mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(s.advice.Content, s.advice.ID).
		WillReturnResult(res).
		WillReturnError(nil)

	err := s.querier.UpdateAdvice(s.ctx, UpdateAdviceParams{
		Content: s.advice.Content,
		ID:      s.advice.ID,
	})

	require.NoError(s.T(), err)
}
