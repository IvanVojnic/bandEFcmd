package repository

import (
	"cmdMS/models"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

// UserMS has an internal grpc object
type UserMS struct {
	conn *grpc.ClientConn
}

// NewUserMS used to init UsesAP
func NewUserMS(conn *grpc.ClientConn) *UserMS {
	return &UserMS{conn: conn}
}

// SignUp used to create user
func (r *UserMS) SignUp(ctx context.Context, user *models.User) error {

	return nil
}

// UpdateRefreshToken used to update rt
func (r *UserMS) UpdateRefreshToken(ctx context.Context, rt string, id uuid.UUID) error {

	return nil
}

// GetUserByID used to get user by ID
func (r *UserMS) GetUserByID(ctx context.Context, userID uuid.UUID) (models.User, error) {
	user := models.User{}

	return user, nil
}

// SignInUser used to sign in user
func (r *UserMS) SignIn(ctx context.Context, user *models.User) error {

	return nil
}
