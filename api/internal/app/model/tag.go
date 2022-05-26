package model

import "github.com/google/uuid"

type Tag struct {
	Tag string    `json:"tag"`
	ID  uuid.UUID `json:"id"`
}
