package model

import (
	"time"

	"github.com/google/uuid"
)

type Page struct {
	ID       uuid.UUID
	Datetime time.Time
	Content  string
}
