// Package utils used to verify tokens
package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// SigningKeyAccess is a secret key for tokens
const SigningKeyAccess = "al5jkvkls83l9cw6l"

// SigningKeyRefresh is a secret key for tokens
const SigningKeyRefresh = "jkvf7834lkjbas98"

// TokenATDuration is a duration of at life
const TokenATDuration = 100 * time.Minute

type tokenClaims struct {
	jwt.StandardClaims
	UserID uuid.UUID `json:"user_id"`
}

// ParseToken used to parse tokens with claims
func ParseToken(tokenToParse string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenToParse, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(SigningKeyAccess), nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if ok && !token.Valid {
		return uuid.UUID{}, fmt.Errorf("invalid Token")
	}

	return claims.UserID, nil
}

// GenerateToken used to generate tokens with id
func GenerateToken(id uuid.UUID, tokenDuration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})
	return token.SignedString([]byte(SigningKeyAccess))
}

// IsAuthorized used to check is user authorized with the tokens
func IsAuthorized(requestToken string) (bool, error) {
	_, err := ParseToken(requestToken)
	if err != nil {
		return false, err
	}
	return true, nil
}

// IsTokenExpired used to check is token expired
func IsTokenExpired(requestToken string) bool {
	token, err := jwt.ParseWithClaims(requestToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(SigningKeyAccess), nil
	})
	v, _ := err.(*jwt.ValidationError)
	tokenExpired := false

	if tk := token.Claims.(jwt.StandardClaims); v.Errors == jwt.ValidationErrorExpired && tk.VerifyExpiresAt(time.Now().Unix(), tokenExpired) {
		return true
	}
	return tokenExpired
}
