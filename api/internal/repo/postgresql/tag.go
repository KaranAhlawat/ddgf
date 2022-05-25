package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/KaranAhlawat/ddgf/internal/app/model"
	db "github.com/KaranAhlawat/ddgf/internal/repo/gen"
	"github.com/google/uuid"
)

type Tag struct {
	q *db.Queries
}

func NewTag(database *sql.DB) *Tag {
	return &Tag{
		q: db.New(database),
	}
}

func (t *Tag) Remove(ctx context.Context, id uuid.UUID) error {
	err := t.q.DeleteTag(ctx, id)
	return fmt.Errorf("db delete: %w", err)
}

func (t *Tag) Insert(ctx context.Context, id uuid.UUID, tag string) (model.Tag, error) {
	res, err := t.q.InsertTag(ctx, db.InsertTagParams{
		ID:  id,
		Tag: tag,
	})
	if err != nil {
		return model.Tag{}, fmt.Errorf("db insert: %w", err)
	}

	return model.Tag{
		ID:  res.ID,
		Tag: res.Tag,
	}, nil
}

func (t *Tag) Select(ctx context.Context, id uuid.UUID) (model.Tag, error) {
	res, err := t.q.SelectTag(ctx, id)
	if err != nil {
		return model.Tag{}, fmt.Errorf("db select: %w", err)
	}

	return model.Tag{
		ID:  res.ID,
		Tag: res.Tag,
	}, nil
}

func (t *Tag) SelectAll(ctx context.Context) ([]model.Tag, error) {
	modelTags := []model.Tag{}
	res, err := t.q.SelectTags(ctx)
	if err != nil {
		return modelTags, fmt.Errorf("db all: %w", err)
	}

	for _, dbTag := range res {
		temp := model.Tag{
			ID:  dbTag.ID,
			Tag: dbTag.Tag,
		}
		modelTags = append(modelTags, temp)
	}

	return modelTags, nil
}

func (t *Tag) SelectAdvices(ctx context.Context, id uuid.UUID) ([]model.Advice, error) {
	modelAdvices := []model.Advice{}
	advices, err := t.q.SelectAdvicesForTag(ctx, id)
	if err != nil {
		return modelAdvices, fmt.Errorf("db select tag-ad: %w", err)
	}

	for _, advice := range advices {
		temp := model.Advice{
			Content: advice.Content,
			ID:      advice.AdviceID,
			Tags:    []model.Tag{},
		}
		modelAdvices = append(modelAdvices, temp)
	}

	return modelAdvices, nil
}

func (a *Tag) SelectTagsForList(ctx context.Context, ids []uuid.UUID) (map[uuid.UUID][]model.Tag, error) {
	tags, err := a.q.SelectTagsForList(ctx, ids)
	m := map[uuid.UUID][]model.Tag{}
	if err != nil {
		return m, fmt.Errorf("db select ad-tag: %w", err)
	}

	for _, tag := range tags {
		temp := model.Tag{
			ID:  tag.TagID,
			Tag: tag.Tag,
		}
		m[tag.AdviceID] = append(m[tag.AdviceID], temp)
	}

	return m, nil
}
