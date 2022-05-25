package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/KaranAhlawat/ddgf/internal/app/model"
	db "github.com/KaranAhlawat/ddgf/internal/repo/gen"
	"github.com/google/uuid"
)

type Advice struct {
	q *db.Queries
}

func NewAdvice(database *sql.DB) *Advice {
	return &Advice{
		q: db.New(database),
	}
}

func (a *Advice) Remove(ctx context.Context, id uuid.UUID) error {
	err := a.q.DeleteAdvice(ctx, id)
	if err != nil {
		return fmt.Errorf("db delete: %w", err)
	} else {
		return nil
	}
}

func (a *Advice) Insert(ctx context.Context, id uuid.UUID, content string) (model.Advice, error) {
	_, err := a.q.InsertAdvice(ctx, db.InsertAdviceParams{
		ID:      id,
		Content: content,
	})
	if err != nil {
		return model.Advice{}, fmt.Errorf("db insert: %w", err)
	} else {
		return model.Advice{
			ID:      id,
			Content: content,
			Tags:    []model.Tag{},
		}, nil
	}
}

func (a *Advice) Select(ctx context.Context, id uuid.UUID) (model.Advice, error) {
	modelTags := []model.Tag{}
	res, err := a.q.SelectAdvice(ctx, id)

	if err != nil {
		return model.Advice{}, fmt.Errorf("db select: %w", err)
	}

	return model.Advice{
		ID:      res.ID,
		Content: res.Content,
		Tags:    modelTags,
	}, nil
}

func (a *Advice) SelectAll(ctx context.Context) ([]model.Advice, error) {
	modelAdvices := []model.Advice{}
	res, err := a.q.SelectAdvices(ctx)
	if err != nil {
		return modelAdvices, fmt.Errorf("db select all: %w", err)
	}

	for _, dbAdvice := range res {
		temp := model.Advice{
			ID:      dbAdvice.ID,
			Content: dbAdvice.Content,
			Tags:    []model.Tag{},
		}

		modelAdvices = append(modelAdvices, temp)
	}

	return modelAdvices, nil
}

func (a *Advice) Update(ctx context.Context, content string, id uuid.UUID) error {
	err := a.q.UpdateAdvice(ctx, db.UpdateAdviceParams{
		Content: content,
		ID:      id,
	})
	if err != nil {
		return fmt.Errorf("db update: %w", err)
	} else {
		return nil
	}
}

func (a *Advice) SelectTags(ctx context.Context, id uuid.UUID) ([]model.Tag, error) {
	modelTags := []model.Tag{}
	tags, err := a.q.SelectTagsForAdvice(ctx, id)
	if err != nil {
		return modelTags, fmt.Errorf("db select ad-tag: %w", err)
	}

	for _, tag := range tags {
		temp := model.Tag{
			ID:  tag.TagID,
			Tag: tag.Tag,
		}
		modelTags = append(modelTags, temp)
	}

	return modelTags, nil
}

func (a *Advice) SelectAllTags(ctx context.Context) (map[uuid.UUID][]model.Tag, error) {
	gtmap := map[uuid.UUID][]model.Tag{}
	res, err := a.q.SelectAllEntries(ctx)
	if err != nil {
		return gtmap, fmt.Errorf("db all ad-tags: %w", err)
	}

	for _, row := range res {
		temp := model.Tag{
			ID:  row.TagID,
			Tag: row.Tag,
		}
		gtmap[row.AdviceID] = append(gtmap[row.AdviceID], temp)
	}

	return gtmap, nil
}

func (a *Advice) InsertTag(ctx context.Context, a_id uuid.UUID, t_id uuid.UUID) error {
	_, err := a.q.InsertAdviceTagEntry(ctx, db.InsertAdviceTagEntryParams{
		AdviceID: a_id,
		TagID:    t_id,
	})

	if err != nil {
		return fmt.Errorf("db insert ad-tag: %w", err)
	}

	return nil
}

func (a *Advice) DeleteTag(ctx context.Context, a_id uuid.UUID, t_id uuid.UUID) error {
	err := a.q.DeleteTagFromAdvice(ctx, db.DeleteTagFromAdviceParams{
		AdviceID: a_id,
		TagID:    t_id,
	})

	if err != nil {
		return fmt.Errorf("db delete ad-tag: %w", err)
	}

	return nil
}
