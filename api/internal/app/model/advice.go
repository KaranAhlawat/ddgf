package model

import "github.com/google/uuid"

type Advice struct {
	Content string    `json:"content"`
	Tags    []Tag     `json:"tags"`
	ID      uuid.UUID `json:"id"`
}
