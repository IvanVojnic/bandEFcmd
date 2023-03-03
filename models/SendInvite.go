package models

import (
	"github.com/google/uuid"
	"time"
)

type SendInvite struct {
	UsersID *[]uuid.UUID `json:"usersID"`
	Place   string       `json:"place"`
	Date    time.Time    `json:"date"`
}
