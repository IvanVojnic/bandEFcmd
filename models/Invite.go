package models

import "github.com/google/uuid"

type Invite struct {
	RoomID uuid.UUID `json:"roomID"`
}
