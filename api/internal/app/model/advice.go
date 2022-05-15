package model

import "github.com/google/uuid"

type Advice struct {
	Content string
	ID      uuid.UUID
}
