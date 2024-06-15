package service

import (
	"context"
	"github.com/webkimru/go-keeper/internal/app/server/models"
)

// UserStore is an interface to store users.
type UserStore interface {
	Add(ctx context.Context, user *models.User) error
	Find(ctx context.Context, login string) (*models.User, error)
}

// UserService contains a user storage.
type UserService struct {
	storage UserStore
}

// NewUserService return a new user service with storage.
func NewUserService(storage UserStore) *UserService {
	return &UserService{storage: storage}
}

// Add puts a user to the storage.
func (s *UserService) Add(ctx context.Context, u *models.User) error {
	user, err := models.NewUser(u.Login, u.Password)
	if err != nil {
		return err
	}

	if err = s.storage.Add(ctx, user); err != nil {
		return err
	}

	return nil
}

// Find looks for a user in the storage.
func (s *UserService) Find(ctx context.Context, login string) (*models.User, error) {
	user, err := s.storage.Find(ctx, login)
	if err != nil {
		return nil, err
	}

	return user, nil
}
