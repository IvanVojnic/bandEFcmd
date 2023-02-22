package service

import (
	"cmdMS/internal/handler"
	"cmdMS/internal/utils"
	"cmdMS/models"
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/google/uuid"
)

const salt = "s53d42fg98gh7j6kkbver"

// Authorization interface consists of methos to communicate with user repo
type Authorization interface {
	SignUp(ctx context.Context, user *models.User) error
	SignIn(ctx context.Context, user *models.User) error
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
	user.Password = generatePasswordHash(user.Password)
	err := s.repo.SignUp(ctx, user)
	if err != nil {
		return fmt.Errorf("error create auth user %w", err)
	}
	return err
}

// GetUserVerified used to get user
func (s *AuthService) GetUser(ctx context.Context, id uuid.UUID) (models.User, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	return user, err
}

// SignInUser used to sign in user
func (s *AuthService) SignIn(ctx context.Context, user *models.User) (bool, handler.Tokens, error) {
	hashedPass := generatePasswordHash(user.Password)
	err := s.repo.SignIn(ctx, user)
	if err != nil {
		return false, handler.Tokens{}, fmt.Errorf("error while sign in query %w", err)
	}
	if user.Password == hashedPass {
		rt, errRT := utils.GenerateToken(user.UserID, utils.TokenRTDuration)
		if errRT != nil {
			return false, handler.Tokens{}, fmt.Errorf("error while creating rt token, %s", errRT)
		}
		at, errAT := utils.GenerateToken(user.UserID, utils.TokenATDuretion)
		if errAT != nil {
			return false, handler.Tokens{}, fmt.Errorf("error while creating at token, %s", errRT)
		}
		errUpdateRT := s.repo.UpdateRefreshToken(ctx, rt, user.UserID)
		if errUpdateRT != nil {
			return false, handler.Tokens{}, fmt.Errorf("error while updating, %s", errUpdateRT)
		}
		return true, handler.Tokens{AccessToken: at, RefreshToken: rt}, nil
	}
	return false, handler.Tokens{}, nil
}

func (s *AuthService) UpdateRefreshToken(context.Context, string, uuid.UUID) error {
	return nil
}

// generatePasswordHash used to generate hash password
func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
