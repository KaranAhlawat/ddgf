package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/KaranAhlawat/ddgf/internal/app/model"
	db "github.com/KaranAhlawat/ddgf/internal/repo/gen"
	"github.com/google/uuid"
)

type Page struct {
	q *db.Queries
}

func NewPage(database *sql.DB) *Page {
	return &Page{
		q: db.New(database),
	}
}

func (p *Page) Remove(ctx context.Context, id uuid.UUID) error {
	err := p.q.DeletePage(ctx, id)
	if err != nil {
		return fmt.Errorf("db delete: %w", err)
	} else {
		return nil
	}
}

func (p *Page) Insert(ctx context.Context, id uuid.UUID, date time.Time, content string) (model.Page, error) {
	_, err := p.q.InsertPage(ctx, db.InsertPageParams{
		ID:       id,
		Datetime: date,
		Content:  content,
	})
	if err != nil {
		return model.Page{}, fmt.Errorf("db insert: %w", err)
	} else {
		return model.Page{
			ID:       id,
			Datetime: date,
			Content:  content,
		}, nil
	}
}

func (p *Page) Select(ctx context.Context, id uuid.UUID) (model.Page, error) {
	res, err := p.q.SelectPage(ctx, id)
	if err != nil {
		return model.Page{}, fmt.Errorf("db select: %w", err)
	} else {
		return model.Page{
			ID:       res.ID,
			Datetime: res.Datetime,
			Content:  res.Content,
		}, nil
	}
}

func (p *Page) SelectAll(ctx context.Context) ([]model.Page, error) {
	modelPages := []model.Page{}
	res, err := p.q.SelectPages(ctx)
	if err != nil {
		return modelPages, fmt.Errorf("db select all: %w", err)
	} else {
		for _, dbPage := range res {
			temp := model.Page{
				ID:       dbPage.ID,
				Datetime: dbPage.Datetime,
				Content:  dbPage.Content,
			}
			modelPages = append(modelPages, temp)
		}

		return modelPages, nil
	}
}

func (p *Page) Update(ctx context.Context, content string, date time.Time, id uuid.UUID) error {
	err := p.q.UpdatePage(ctx, db.UpdatePageParams{
		Content:  content,
		Datetime: date,
		ID:       id,
	})
	if err != nil {
		return fmt.Errorf("db update: %w", err)
	} else {
		return nil
	}
}
