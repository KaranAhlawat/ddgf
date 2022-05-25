package db

import (
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func (s *Suite) TestDeleteTagFromAdvice() {
	query := `-- name: DeleteTagFromAdvice :exec
    DELETE FROM "advices_tags"
    WHERE "advice_id" = $1
        AND "tag_id" = $2
    `

	res := sqlmock.NewResult(0, 1)

	s.mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(s.advice.ID, s.tag.ID).
		WillReturnResult(res).
		WillReturnError(nil)

	err := s.querier.DeleteTagFromAdvice(s.ctx, DeleteTagFromAdviceParams{
		AdviceID: s.advice.ID,
		TagID:    s.tag.ID,
	})

	require.NoError(s.T(), err)
}

func (s *Suite) TestInsertAdviceTagEntry() {
	query := `-- name: InsertAdviceTagEntry :one
    INSERT INTO "advices_tags" ("advice_id", "tag_id")
        VALUES ($1, $2)
    RETURNING
        advice_id, tag_id
    `

	rows := sqlmock.NewRows([]string{"advice_id", "tag_id"}).
		AddRow(s.advice.ID, s.tag.ID)

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(s.advice.ID, s.tag.ID).
		WillReturnRows(rows)

	res, err := s.querier.InsertAdviceTagEntry(s.ctx, InsertAdviceTagEntryParams{
		AdviceID: s.advice.ID,
		TagID:    s.tag.ID,
	})

	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
	require.Equal(s.T(), s.advice.ID, res.AdviceID)
	require.Equal(s.T(), s.tag.ID, res.TagID)
}

func (s *Suite) TestSelectAdvicesForTag() {
	query := `-- name: SelectAdvicesForTag :many
    SELECT
        "at"."advice_id",
        "a"."content"
    FROM
        "advices_tags" "at"
        JOIN "advices" "a" ON "at"."advice_id" = "a"."id"
    WHERE
        "at"."tag_id" = $1
    `

	rows := sqlmock.NewRows([]string{"advice_id", "content"}).
		AddRow(s.advice.ID, s.advice.Content).
		AddRow(s.advice.ID, s.advice.Content)

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(s.tag.ID).
		WillReturnRows(rows)

	res, err := s.querier.SelectAdvicesForTag(s.ctx, s.tag.ID)

	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)

	for _, advice := range res {
		require.Equal(s.T(), s.advice.ID, advice.AdviceID)
		require.Equal(s.T(), s.advice.Content, advice.Content)
	}
}

func (s *Suite) TestSelectAllEntries() {
	query := `-- name: SelectAllEntries :many
    SELECT
        advice_id, tag_id
    FROM
        "advices_tags"
    ORDER BY
        "advice_id"
    `

	rows := sqlmock.NewRows([]string{"advice_id", "tag_id"}).
		AddRow(s.advice.ID, s.tag.ID).
		AddRow(s.advice.ID, s.tag.ID)

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows)

	res, err := s.querier.SelectAllEntries(s.ctx)

	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)

	for _, adviceTag := range res {
		require.Equal(s.T(), adviceTag.AdviceID, s.advice.ID)
		require.Equal(s.T(), adviceTag.TagID, s.tag.ID)
	}
}

func (s *Suite) TestSelectTagsForAdvice() {
	query := `-- name: SelectTagsForAdvice :many
    SELECT
        "at"."tag_id",
        "t"."tag"
    FROM
        "advices_tags" "at"
        JOIN "tags" "t" ON "at"."tag_id" = "t"."id"
    WHERE
        "at"."advice_id" = $1
    `

	rows := sqlmock.NewRows([]string{"tag_id", "tag"}).
		AddRow(s.tag.ID, s.tag.Tag).
		AddRow(s.tag.ID, s.tag.Tag)

	s.mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(s.advice.ID).
		WillReturnRows(rows)

	res, err := s.querier.SelectTagsForAdvice(s.ctx, s.advice.ID)

	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)

	for _, tag := range res {
		require.EqualValues(s.T(), s.tag.ID, tag.TagID)
		require.EqualValues(s.T(), s.tag.Tag, tag.Tag)
	}
}
