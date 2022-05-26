package service

import (
	"context"
	"fmt"

	"github.com/KaranAhlawat/ddgf/internal/app/model"
	repo "github.com/KaranAhlawat/ddgf/internal/repo/postgresql"
	"github.com/google/uuid"
)

type Advice struct {
	r   *repo.Advice
	ctx context.Context
}

func NewAdvice(ctx context.Context, repository *repo.Advice) *Advice {
	return &Advice{
		repository,
		ctx,
	}
}

// Add a new advice with no tag
func (a *Advice) Add(advice *model.Advice) error {
	_, err := a.r.Insert(a.ctx, advice.ID, advice.Content)
	if err != nil {
		return fmt.Errorf("advice add: %w", err)
	}
	return nil
}

// Delete an advice
func (a *Advice) Delete(id uuid.UUID) error {
	err := a.r.Remove(a.ctx, id)
	if err != nil {
		return fmt.Errorf("advice delete: %w", err)
	}
	return nil
}

// Select an advice with all it's tags
func (a *Advice) Get(id uuid.UUID) (model.Advice, error) {
	advice, err := a.r.Select(a.ctx, id)
	if err != nil {
		return model.Advice{}, fmt.Errorf("advice get: %w", err)
	}
	tags, err := a.r.SelectTags(a.ctx, advice.ID)
	if err != nil {
		return model.Advice{}, fmt.Errorf("advice get: %w", err)
	}
	advice.Tags = tags
	return advice, nil
}

// Select all advices with all their tags
func (a *Advice) All() ([]model.Advice, error) {
	advices, err := a.r.SelectAll(a.ctx)
	if err != nil {
		return []model.Advice{}, fmt.Errorf("advice all: %w", err)
	}

	tags, err := a.r.SelectAllTags(a.ctx)
	if err != nil {
		return []model.Advice{}, fmt.Errorf("advice all: %w", err)
	}

	for _, advice := range advices {
		advice.Tags = tags[advice.ID]
	}

	return advices, nil
}

// Update an advice
func (a *Advice) Update(id uuid.UUID, advice *model.Advice) error {
	err := a.r.Update(a.ctx, advice.Content, advice.ID)
	if err != nil {
		return fmt.Errorf("advice update: %w", err)
	}

	return nil
}

// Tag an advice with given tag
func (a *Advice) AddTag(adviceID uuid.UUID, tagID uuid.UUID) error {
	err := a.r.InsertTag(a.ctx, adviceID, tagID)
	if err != nil {
		return fmt.Errorf("advice tag: %w", err)
	}
	return nil
}

// Untag an advice
func (a *Advice) Untag(adviceID uuid.UUID, tagID uuid.UUID) error {
	err := a.r.DeleteTag(a.ctx, adviceID, tagID)
	if err != nil {
		return fmt.Errorf("advice untag: %w", err)
	}
	return nil
}
