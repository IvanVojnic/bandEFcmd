// Package service room invite service
package service

import (
	"cmdMS/models"
	"context"
	"time"

	"github.com/google/uuid"
)

// RoomInvite interface define rooms ad invites methods
type RoomInvite interface {
	SendInvite(ctx context.Context, userCreatorID uuid.UUID, usersID []*uuid.UUID, place string, date time.Time) error
	AcceptInvite(ctx context.Context, userID, roomID uuid.UUID) error
	DeclineInvite(ctx context.Context, userID, roomID uuid.UUID) error
	GetRooms(ctx context.Context, user uuid.UUID) ([]*models.Room, error)
	GetRoomUsers(ctx context.Context, roomID uuid.UUID) ([]*models.User, error)
}

// RoomInviteService is wrapper for user repo
type RoomInviteService struct {
	repo RoomInvite
}

// NewRoomInviteService used to init AS
func NewRoomInviteService(repo RoomInvite) *RoomInviteService {
	return &RoomInviteService{repo: repo}
}

// SendInvite used to send invite by repo
func (s *RoomInviteService) SendInvite(ctx context.Context, userCreatorID uuid.UUID, usersID []*uuid.UUID, place string, date time.Time) error {
	return s.repo.SendInvite(ctx, userCreatorID, usersID, place, date)
}

// AcceptInvite used to accept invite by repo
func (s *RoomInviteService) AcceptInvite(ctx context.Context, userID, roomID uuid.UUID) error { // nolint:dupl, gocritic
	return s.repo.AcceptInvite(ctx, userID, roomID)
}

// DeclineInvite used to decline invite by repo
func (s *RoomInviteService) DeclineInvite(ctx context.Context, userID, roomID uuid.UUID) error { // nolint:dupl, gocritic
	return s.repo.DeclineInvite(ctx, userID, roomID)
}

// GetRooms used to get rooms by repo
func (s *RoomInviteService) GetRooms(ctx context.Context, user uuid.UUID) ([]*models.Room, error) {
	return s.repo.GetRooms(ctx, user)
}

// GetRoomUsers used to get users from current room by repo
func (s *RoomInviteService) GetRoomUsers(ctx context.Context, roomID uuid.UUID) ([]*models.User, error) {
	return s.repo.GetRoomUsers(ctx, roomID)
}
