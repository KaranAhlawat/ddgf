package model

import (
	"time"

	"github.com/google/uuid"
)

type Page struct {
	Datetime time.Time
	Content  string
	ID       uuid.UUID
}
