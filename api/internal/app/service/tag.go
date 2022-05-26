package service

import (
	"context"
	"fmt"

	"github.com/KaranAhlawat/ddgf/internal/app/model"
	repo "github.com/KaranAhlawat/ddgf/internal/repo/postgresql"
	"github.com/google/uuid"
)

type Tag struct {
	r   *repo.Tag
	ctx context.Context
}

func NewTag(ctx context.Context, repository *repo.Tag) *Tag {
	return &Tag{
		repository,
		ctx,
	}
}

// Add a tag
func (t *Tag) Add(tag *model.Tag) error {
	_, err := t.r.Insert(t.ctx, tag.ID, tag.Tag)
	if err != nil {
		return fmt.Errorf("tag add: %w", err)
	}
	return nil
}

// Delete a tag
func (t *Tag) Delete(id uuid.UUID) error {
	err := t.r.Remove(t.ctx, id)
	if err != nil {
		return fmt.Errorf("tag delete: %w", err)
	}
	return nil
}

// Get a tag
func (t *Tag) Get(id uuid.UUID) (model.Tag, error) {
	tag, err := t.r.Select(t.ctx, id)
	if err != nil {
		return model.Tag{}, fmt.Errorf("tag get: %w", err)
	}
	return tag, nil
}

// Get all tags
func (t *Tag) All() ([]model.Tag, error) {
	tags, err := t.r.SelectAll(t.ctx)
	if err != nil {
		return []model.Tag{}, fmt.Errorf("tag all: %w", err)
	}
	return tags, nil
}

// Get all advices for a given tag
func (t *Tag) ListAdvices(id uuid.UUID) ([]model.Advice, error) {
	advices, err := t.r.SelectAdvices(t.ctx, id)
	if err != nil {
		return []model.Advice{}, fmt.Errorf("tag list-ad: %w", err)
	}
	tmpIds := []uuid.UUID{}

	for _, advice := range advices {
		tmpIds = append(tmpIds, advice.ID)
	}

	m, err := t.r.SelectTagsForList(t.ctx, tmpIds)
	if err != nil {
		return []model.Advice{}, fmt.Errorf("tag list-ad: %w", err)
	}

	for _, advice := range advices {
		advice.Tags = m[advice.ID]
	}

	return advices, nil
}
