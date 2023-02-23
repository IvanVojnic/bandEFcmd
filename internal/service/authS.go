package service

import (
	"cmdMS/internal/handler"
	"cmdMS/models"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

// Authorization interface consists of methos to communicate with user repo
type Authorization interface {
	SignUp(ctx context.Context, user *models.User) error
	SignIn(ctx context.Context, user *models.User) (handler.Tokens, error)
	GetUserByID(context.Context, uuid.UUID) (models.User, error)
	UpdateRefreshToken(context.Context, string, uuid.UUID) error
}

// AuthService is wrapper for user repo
type AuthService struct {
	repo Authorization
}

// NewAuthService used to init AS
func NewAuthService(repo Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// SignUp used to
func (s *AuthService) SignUp(ctx context.Context, user *models.User) error {
	errPass := generatePasswordHash(user)
	if errPass != nil {
		return fmt.Errorf("cannot hash password, %s", errPass)
	}
	err := s.repo.SignUp(ctx, user)
	if err != nil {
		return fmt.Errorf("error create auth user %w", err)
	}
	return nil
}

// GetUserVerified used to get user
func (s *AuthService) GetUser(ctx context.Context, id uuid.UUID) (models.User, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return models.User{}, fmt.Errorf("error while getting user, %s", err)
	}
	return user, nil
}

// SignIn used to sign in user
func (s *AuthService) SignIn(ctx context.Context, user *models.User) (handler.Tokens, error) {
	tokens, err := s.repo.SignIn(ctx, user)
	if err != nil {
		return handler.Tokens{}, fmt.Errorf("error while sign in query %w", err)
	}
	return tokens, nil
}

func (s *AuthService) UpdateRefreshToken(context.Context, string, uuid.UUID) error {
	return nil
}

// generatePasswordHash used to generate hash password
func generatePasswordHash(user *models.User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(bytes)
	return err
}
