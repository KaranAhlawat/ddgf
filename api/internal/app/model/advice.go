package model

import "github.com/google/uuid"

type Advice struct {
	Content string
	Tags    []Tag
	ID      uuid.UUID
}
