package models

import (
	"github.com/google/uuid"
	"time"
)

// Room is a Room
type Room struct {
	ID            uuid.UUID `json:"id" db:"id"`
	UserCreatorID uuid.UUID `json:"userCreatorID" db:"idUserCreator"`
	Place         string    `json:"place" db:"place"`
	Date          time.Time `json:"date" db:"date"`
}
