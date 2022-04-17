package repo

import (
	"context"
	"database/sql"
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
}
