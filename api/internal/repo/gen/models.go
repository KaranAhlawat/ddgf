// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type Page struct {
	ID       uuid.UUID `json:"id"`
	Datetime time.Time `json:"datetime"`
	Content  string    `json:"content"`
}
