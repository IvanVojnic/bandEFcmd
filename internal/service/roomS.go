package service

import (
	"cmdMS/models"
	"context"
	"github.com/google/uuid"
	"time"
)

type RoomInvite interface {
	SendInvite(ctx context.Context, userCreatorID uuid.UUID, usersID *[]uuid.UUID, place string, date time.Time) error
	AcceptInvite(ctx context.Context, userID uuid.UUID, roomID uuid.UUID) error
	DeclineInvite(ctx context.Context, userID uuid.UUID, roomID uuid.UUID) error
	GetRooms(ctx context.Context, user uuid.UUID) (*[]models.Room, error)
	GetRoomUsers(ctx context.Context, roomID uuid.UUID) (*[]models.User, error)
}

// RoomInviteService is wrapper for user repo
type RoomInviteService struct {
	repo RoomInvite
}

// NewRoomInviteService used to init AS
func NewRoomInviteService(repo RoomInvite) *RoomInviteService {
	return &RoomInviteService{repo: repo}
}

func (s *RoomInviteService) SendInvite(ctx context.Context, userCreatorID uuid.UUID, usersID *[]uuid.UUID, place string, date time.Time) error {
	return s.repo.SendInvite(ctx, userCreatorID, usersID, place, date)
}

func (s *RoomInviteService) AcceptInvite(ctx context.Context, userID uuid.UUID, roomID uuid.UUID) error {
	return s.repo.AcceptInvite(ctx, userID, roomID)
}

func (s *RoomInviteService) DeclineInvite(ctx context.Context, userID uuid.UUID, roomID uuid.UUID) error {
	return s.repo.DeclineInvite(ctx, userID, roomID)
}

func (s *RoomInviteService) GetRooms(ctx context.Context, user uuid.UUID) (*[]models.Room, error) {
	return s.repo.GetRooms(ctx, user)
}

func (s *RoomInviteService) GetRoomUsers(ctx context.Context, roomID uuid.UUID) (*[]models.User, error) {
	return s.repo.GetRoomUsers(ctx, roomID)
}
