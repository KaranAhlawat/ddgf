package db

import (
	"context"
	"database/sql"
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
	tag     *Tag
	ctx     context.Context
}

func (s *Suite) SetupSuite() {
	var err error

	s.DB, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.querier = New(s.DB)

	id := uuid.New()

	s.page = &Page{
		ID:       id,
		Datetime: time.Now(),
		Content:  "Testing page with random id",
	}

	s.tag = &Tag{
		ID:  id,
		Tag: "Testing",
	}

	s.ctx = context.Background()
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestMain(t *testing.T) {
	suite.Run(t, new(Suite))
}
