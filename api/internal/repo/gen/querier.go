// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	DeletePage(ctx context.Context, id uuid.UUID) error
	InsertPage(ctx context.Context, arg InsertPageParams) (Page, error)
	SelectPage(ctx context.Context, id uuid.UUID) (Page, error)
	SelectPages(ctx context.Context) ([]Page, error)
	UpdatePage(ctx context.Context, arg UpdatePageParams) error
}

var _ Querier = (*Queries)(nil)
