// Package service user auth service
package service

import (
	"cmdMS/models"
	"context"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Authorization interface consists of methos to communicate with user repo
type Authorization interface {
	SignUp(ctx context.Context, user *models.User) error
	SignIn(ctx context.Context, user *models.User) (models.Tokens, error)
	UpdateRefreshToken(context.Context, string, uuid.UUID) error
}

// cost used to generate password
const cost = 14

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

// SignIn used to sign in user
func (s *AuthService) SignIn(ctx context.Context, user *models.User) (models.Tokens, error) {
	tokens, err := s.repo.SignIn(ctx, user)
	if err != nil {
		return models.Tokens{}, fmt.Errorf("error while sign in query %w", err)
	}
	return tokens, nil
}

// UpdateRefreshToken used to update refresh token
func (s *AuthService) UpdateRefreshToken(context.Context, string, uuid.UUID) error {
	return nil
}

// generatePasswordHash used to generate hash password
func generatePasswordHash(user *models.User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), cost)
	user.Password = string(bytes)
	return err
}
