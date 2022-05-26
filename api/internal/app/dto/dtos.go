package dto

import "github.com/google/uuid"

type PagePostDTO struct {
	Content string    `json:"content"`
	ID      uuid.UUID `json:"id"`
}

type AdvicePostDTO struct {
	Content string    `json:"content"`
	ID      uuid.UUID `json:"id"`
}

type TagPostDTO struct {
	Tag string    `json:"tag"`
	ID  uuid.UUID `json:"id"`
}
