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

func (a *Advice) Delete(ctx context.Context, id uuid.UUID) error {
	err := a.q.DeleteAdvice(ctx, id)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	} else {
		return nil
	}
}

func (a *Advice) Create(ctx context.Context, id uuid.UUID, content string) (model.Advice, error) {
	_, err := a.q.InsertAdvice(ctx, db.InsertAdviceParams{
		ID:      id,
		Content: content,
	})
	if err != nil {
		return model.Advice{}, fmt.Errorf("insert: %w", err)
	} else {
		return model.Advice{
			ID:      id,
			Content: content,
		}, nil
	}
}

func (a *Advice) Find(ctx context.Context, id uuid.UUID) (model.Advice, error) {
	res, err := a.q.SelectAdvice(ctx, id)
	if err != nil {
		return model.Advice{}, fmt.Errorf("select: %w", err)
	} else {
		return model.Advice{
			ID:      res.ID,
			Content: res.Content,
		}, nil
	}
}

func (a *Advice) All(ctx context.Context, id uuid.UUID) ([]model.Advice, error) {
	modelAdvices := []model.Advice{}
	res, err := a.q.SelectAdvices(ctx)
	if err != nil {
		return modelAdvices, fmt.Errorf("select all: %w", err)
	} else {
		for _, dbAdvice := range res {
			temp := model.Advice{
				ID:      dbAdvice.ID,
				Content: dbAdvice.Content,
			}
			modelAdvices = append(modelAdvices, temp)
		}

		return modelAdvices, nil
	}
}

func (a *Advice) Update(ctx context.Context, content string, id uuid.UUID) error {
	err := a.q.UpdateAdvice(ctx, db.UpdateAdviceParams{
		Content: content,
		ID:      id,
	})
	if err != nil {
		return fmt.Errorf("update: %w", err)
	} else {
		return nil
	}
}