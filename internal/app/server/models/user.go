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

// NewUser returns a new user with hashed password.
func NewUser(login string, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("User -  NewUser - bcrypt.GenerateFromPassword(): %w", err)
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

func (user *User) Validate(required ...string) (string, error) {
	for _, field := range required {
		if !user.valid(field) {
			return field, fmt.Errorf("field %s is required", field)
		}
	}

	return "", nil
}

func (user *User) valid(field string) bool {
	switch field {
	case "login":
		if user.Login == "" {
			return false
		}
	case "password":
		if user.Password == "" {
			return false
		}
	}

	return true
}
