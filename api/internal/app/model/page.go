package model

import (
	"time"

	"github.com/google/uuid"
)

type Page struct {
	Datetime time.Time `json:"datetime"`
	Content  string    `json:"content"`
	ID       uuid.UUID `json:"id"`
}
