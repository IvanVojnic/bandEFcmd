package repository

import (
	"cmdMS/models"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// UserPostgres has an internal db object
type UserPostgres struct {
	db *pgxpool.Pool
}

// NewUserAuthPostgres used to init UsesAP
func NewUserPostgres(db *pgxpool.Pool) *UserPostgres {
	return &UserPostgres{db: db}
}

// SignUp used to create user
func (r *UserPostgres) SignUp(ctx context.Context, user *models.User) error {

	return nil
}

// UpdateRefreshToken used to update rt
func (r *UserPostgres) UpdateRefreshToken(ctx context.Context, rt string, id uuid.UUID) error {

	return nil
}

// GetUserByID used to get user by ID
func (r *UserPostgres) GetUserByID(ctx context.Context, userID uuid.UUID) (models.User, error) {
	user := models.User{}

	return user, nil
}

// SignInUser used to sign in user
func (r *UserPostgres) SignIn(ctx context.Context, user *models.User) error {

	return nil
}
