package service

import (
	"context"
	"fmt"

	"github.com/KaranAhlawat/ddgf/internal/app/model"
	repo "github.com/KaranAhlawat/ddgf/internal/repo/postgresql"
	"github.com/google/uuid"
)

type Page struct {
	r   *repo.Page
	ctx context.Context
}

func NewPage(ctx context.Context, repository *repo.Page) *Page {
	return &Page{
		repository,
		ctx,
	}
}

// Add a new page
func (p *Page) Add(page *model.Page) error {
	_, err := p.r.Insert(p.ctx, page.ID, page.Datetime, page.Content)
	if err != nil {
		return fmt.Errorf("page add: %w", err)
	}
	return nil
}

// Delete a page
func (p *Page) Delete(id uuid.UUID) error {
	err := p.r.Remove(p.ctx, id)
	if err != nil {
		return fmt.Errorf("page delete: %w", err)
	}
	return nil
}

// Update a page
func (p *Page) Update(id uuid.UUID, page *model.Page) error {
	err := p.r.Update(p.ctx, page.Content, page.Datetime, id)
	if err != nil {
		return fmt.Errorf("page update: %w", err)
	}
	return nil
}

// Get a page
func (p *Page) Get(id uuid.UUID) (model.Page, error) {
	page, err := p.r.Select(p.ctx, id)
	if err != nil {
		return model.Page{}, fmt.Errorf("page get: %w", err)
	}
	return page, nil
}

// Get all pages
func (p *Page) All() ([]model.Page, error) {
	pages, err := p.r.SelectAll(p.ctx)
	if err != nil {
		return []model.Page{}, fmt.Errorf("page all: %w", err)
	}
	return pages, nil
}
