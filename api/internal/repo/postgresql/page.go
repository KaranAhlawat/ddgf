package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/KaranAhlawat/ddgf/internal/app/model"
	db "github.com/KaranAhlawat/ddgf/internal/repo/gen"
	"github.com/google/uuid"
)

type PageRepo struct {
	q *db.Queries
}

func NewPageRepo(database *sql.DB) *PageRepo {
	return &PageRepo{
		q: db.New(database),
	}
}

func (p *PageRepo) Delete(ctx context.Context, id uuid.UUID) error {
	err := p.q.DeletePage(ctx, id)
	return err
}

func (p *PageRepo) Create(ctx context.Context, id uuid.UUID, datetime time.Time, content string) (model.Page, error) {
	_, err := p.q.InsertPage(ctx, db.InsertPageParams{
		ID:       id,
		Datetime: datetime,
		Content:  content,
	})
	if err != nil {
		return model.Page{}, err
	}

	return model.Page{
		ID:       id,
		Datetime: datetime,
		Content:  content,
	}, nil
}

func (p *PageRepo) Find(ctx context.Context, id uuid.UUID) (model.Page, error) {
	res, err := p.q.SelectPage(ctx, id)
	if err != nil {
		return model.Page{}, err
	}

	return model.Page{
		ID:       res.ID,
		Datetime: res.Datetime,
		Content:  res.Content,
	}, nil
}

func (p *PageRepo) All(ctx context.Context) ([]model.Page, error) {
	modelPages := []model.Page{}
	res, err := p.q.SelectPages(ctx)
	if err != nil {
		return modelPages, err
	}

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

func (p *PageRepo) Update(ctx context.Context, content string, datetime time.Time, id uuid.UUID) error {
	err := p.q.UpdatePage(ctx, db.UpdatePageParams{
		Content:  content,
		Datetime: datetime,
		ID:       id,
	})
	return err
}
