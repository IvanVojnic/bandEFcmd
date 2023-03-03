package service

import (
	"cmdMS/models"
	"context"
	"github.com/google/uuid"
)

// UserComm interface consists of methods to users actions
type UserComm interface {
	AcceptFriendsRequest(ctx context.Context, userSenderID uuid.UUID, userReceiverID uuid.UUID) error
	DeclineFriendsRequest(ctx context.Context, userSenderID uuid.UUID, userID uuid.UUID) error
	GetFriends(ctx context.Context, userID uuid.UUID) ([]models.User, error)
	SendFriendsRequest(ctx context.Context, userSender uuid.UUID, userReceiver uuid.UUID) error
	FindUser(ctx context.Context, userEmail string) (models.User, error)
	GetRequest(ctx context.Context, userID uuid.UUID) ([]models.User, error)
}

// UserCommSrv wrapper for UserCommP repo
type UserCommSrv struct {
	repo UserComm
}

// NewUserCommSrv used to init UserCommP
func NewUserCommSrv(repo UserComm) *UserCommSrv {
	return &UserCommSrv{repo: repo}
}

// AcceptFriendsRequest used to accept friends request
func (s *UserCommSrv) AcceptFriendsRequest(ctx context.Context, userSenderID uuid.UUID, userReceiverID uuid.UUID) error {
	return s.repo.AcceptFriendsRequest(ctx, userSenderID, userReceiverID)
}

// DeclineFriendsRequest used to accept friends request
func (s *UserCommSrv) DeclineFriendsRequest(ctx context.Context, userSenderID uuid.UUID, userID uuid.UUID) error {
	return s.repo.DeclineFriendsRequest(ctx, userSenderID, userID)
}

// GetFriends used send users friends
func (s *UserCommSrv) GetFriends(ctx context.Context, userID uuid.UUID) ([]models.User, error) {
	return s.repo.GetFriends(ctx, userID)
}

// SendFriendsRequest used users requests
func (s *UserCommSrv) SendFriendsRequest(ctx context.Context, userSender uuid.UUID, userReceiver uuid.UUID) error {
	return s.repo.SendFriendsRequest(ctx, userSender, userReceiver)
}

// FindUser used find user by email
func (s *UserCommSrv) FindUser(ctx context.Context, userEmail string) (models.User, error) {
	return s.repo.FindUser(ctx, userEmail)
}

// GetRequest used send request to the user
func (s *UserCommSrv) GetRequest(ctx context.Context, userID uuid.UUID) ([]models.User, error) {
	return s.repo.GetRequest(ctx, userID)
}
