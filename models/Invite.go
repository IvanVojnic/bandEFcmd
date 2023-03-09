// Package models Invite
package models

import "github.com/google/uuid"

// Invite is an Invite
type Invite struct {
	RoomID uuid.UUID `json:"roomID"`
}
