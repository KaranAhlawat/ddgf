package model

import "github.com/google/uuid"

type Advice struct {
	ID      uuid.UUID
	Content string
}
