package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// User contains user's information.
type User struct {
	ID        int64
	Login     string
	Password  string
	CreatedAt string
}

// NewUser returns a new user.
func NewUser(login string, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("user model GenerateFromPassword(): %w", err)
	}

	user := &User{
		Login:    login,
		Password: string(hashedPassword),
	}

	return user, nil
}

// ValidPassword checks valid password.
func (user *User) ValidPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}
