package repo

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	DB      *sql.DB
	mock    sqlmock.Sqlmock
	querier *Queries
	page    *Page
	ctx     context.Context
}

func (s *Suite) SetupSuite() {
	var err error

	s.DB, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.querier = New(s.DB)

	id, err := uuid.NewUUID()
	require.NoError(s.T(), err)

	s.page = &Page{
		ID:       id,
		Datetime: time.Now(),
		Content:  "Testing page with random id",
	}

	s.ctx = context.Background()
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestMain(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestDeletePage() {
	query := `-- name: DeletePage :exec
	DELETE FROM "pages"
	WHERE "id" = $1
	`
	res := sqlmock.NewResult(0, 1)

	s.mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(s.page.ID).
		WillReturnResult(res).
		WillReturnError(nil)

	err := s.querier.DeletePage(s.ctx, s.page.ID)
	require.NoError(s.T(), err)
}

func (s *Suite) TestInsertPage() {
	query := `-- name: InsertPage :one
	INSERT INTO "pages" ("id", "datetime", "content")
	VALUES ($1, $2, $3)
	RETURNING id, datetime, content
	`

	rows := s.mock.NewRows([]string{"id", "datetime", "content"}).
		AddRow(s.page.ID, s.page.Datetime, s.page.Content)

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(s.page.ID, s.page.Datetime, s.page.Content).
		WillReturnRows(rows)

	res, err := s.querier.InsertPage(s.ctx, InsertPageParams{
		ID:       s.page.ID,
		Datetime: s.page.Datetime,
		Content:  s.page.Content,
	})

	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
	require.EqualValues(s.T(), *s.page, res)
}

func (s *Suite) TestSelectPage() {
	query := `-- name: SelectPage :one
	SELECT id, datetime, content
	FROM "pages"
	WHERE "id" = $1
	LIMIT 1
	`

	rows := s.mock.NewRows([]string{"id", "datetime", "content"}).
		AddRow(s.page.ID, s.page.Datetime, s.page.Content)

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(s.page.ID).
		WillReturnRows(rows)

	res, err := s.querier.SelectPage(s.ctx, s.page.ID)

	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
	require.EqualValues(s.T(), *s.page, res)
}

func (s *Suite) TestSelectPages() {
	query := `-- name: SelectPages :many
	SELECT id, datetime, content
	FROM "pages"
	ORDER BY "datetime"
	`

	rows := s.mock.NewRows([]string{"id", "datetime", "content"}).
		AddRow(s.page.ID, s.page.Datetime, s.page.Content).
		AddRow(s.page.ID, s.page.Datetime, s.page.Content)

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows)

	res, err := s.querier.SelectPages(s.ctx)

	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)

	for _, page := range res {
		require.EqualValues(s.T(), *s.page, page)
	}
}

func (s *Suite) TestErrNoRows() {
	query := `-- name: SelectPages :many
	SELECT id, datetime, content
	FROM "pages"
	ORDER BY "datetime"
	`

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnError(sql.ErrNoRows)

	_, err := s.querier.SelectPages(s.ctx)

	require.Error(s.T(), err)
}

func (s *Suite) TestErrRowError() {
	query := `-- name: SelectPages :many
	SELECT id, datetime, content
	FROM "pages"
	ORDER BY "datetime"
	`

	rows := sqlmock.NewRows([]string{"id", "datetime", "content"}).
		AddRow(s.page.ID, s.page.Datetime, s.page.Content).
		RowError(0, fmt.Errorf("row error"))

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows)

	res, err := s.querier.SelectPages(s.ctx)
	require.Error(s.T(), err)
	require.Nil(s.T(), res)
}

func (s *Suite) TestUpdatePage() {
	query := `-- name: UpdatePage :exec
	UPDATE "pages"
	SET "content" = $1,
	"datetime" = $2
	WHERE "id" = $3
	`

	res := sqlmock.NewResult(0, 1)
	params := UpdatePageParams{
		Content:  "Another string for testing purposes",
		Datetime: time.Now(),
		ID:       s.page.ID,
	}

	s.mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(params.Content, params.Datetime, params.ID).
		WillReturnResult(res).
		WillReturnError(nil)

	err := s.querier.UpdatePage(s.ctx, params)

	require.NoError(s.T(), err)
}
