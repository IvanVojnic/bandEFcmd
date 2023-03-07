// Package models SendInvite
package models

import (
	"time"

	"github.com/google/uuid"
)

// SendInvite is a SendInvite
type SendInvite struct {
	UsersID []*uuid.UUID `json:"usersID"`
	Place   string       `json:"place"`
	Date    time.Time    `json:"date"`
}
