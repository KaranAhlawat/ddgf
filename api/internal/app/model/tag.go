package model

import "github.com/google/uuid"

type Tag struct {
	ID  uuid.UUID
	Tag string
}
