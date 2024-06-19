// Package jwtmanager allows creating a new token with user ID claim ang getting it out the token.
package jwtmanager

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// JWTManager config.
type JWTManager struct {
	secretKey string
	tokenExp  int
}

// UserClaims is a custom JWT claims.
type UserClaims struct {
	jwt.RegisteredClaims
	UserID int64
}

// New returns a new token.
func New(secretKey string, tokenExp int) *JWTManager {
	return &JWTManager{secretKey: secretKey, tokenExp: tokenExp}
}

// BuildJWTString create a new token with algorithm sign HS256 and custom claim.
func (j JWTManager) BuildJWTString(userID int64) (string, error) {
	// Create a new token with algorithm sign HS256 and Claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			// Creation date
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(j.tokenExp))),
		},
		// Custom claim
		UserID: userID,
	})

	// Create token string
	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GetUserID returns a valid user identifier or -1.
func (j JWTManager) GetUserID(tokenString string) int64 {
	claims := &UserClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		// Check header of the algorithm
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return -1
	}

	if !token.Valid {
		return -1
	}

	return claims.UserID
}
