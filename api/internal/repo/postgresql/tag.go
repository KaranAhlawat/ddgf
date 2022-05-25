package repo

import (
	"context"
	"database/sql"

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
	return err
}

func (t *Tag) Insert(ctx context.Context, id uuid.UUID, tag string) (model.Tag, error) {
	res, err := t.q.InsertTag(ctx, db.InsertTagParams{
		ID:  id,
		Tag: tag,
	})
	if err != nil {
		return model.Tag{}, err
	}

	return model.Tag{
		ID:  res.ID,
		Tag: res.Tag,
	}, nil
}

func (t *Tag) Select(ctx context.Context, id uuid.UUID) (model.Tag, error) {
	res, err := t.q.SelectTag(ctx, id)
	if err != nil {
		return model.Tag{}, err
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
		return modelTags, err
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
