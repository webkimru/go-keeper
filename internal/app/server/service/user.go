package service

import (
	"context"
	"fmt"
	"github.com/webkimru/go-keeper/internal/app/server/models"
	"github.com/webkimru/go-keeper/pkg/errs"
)

//go:generate mockgen -destination=mocks/mock_user.go -package=mocks github.com/webkimru/go-keeper/internal/app/server/service UserStore

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
func (s *UserService) Add(ctx context.Context, model *models.User) error {
	if field, err := model.Validate("login", "password"); err != nil {
		return fmt.Errorf("UserService - Add - model.Validate(): %w: %s is required", errs.ErrBadRequest, field)
	}

	user, err := models.NewUser(model.Login, model.Password)
	if err != nil {
		return fmt.Errorf("UserService - Add - models.NewUser(): %w", err)
	}

	if err = s.storage.Add(ctx, user); err != nil {
		return fmt.Errorf("UserService - Add - s.storage.Add(): %w", err)
	}

	return nil
}

// Find returns an existing user.
func (s *UserService) Find(ctx context.Context, login, password string) (*models.User, error) {
	model := models.User{Login: login, Password: password}
	if field, err := model.Validate("login", "password"); err != nil {
		return nil, fmt.Errorf("UserService - Find - model.Validate(): %w: %s is required", errs.ErrBadRequest, field)
	}

	user, err := s.storage.Find(ctx, login)
	if err != nil {
		return nil, fmt.Errorf("UserService - Find - s.storage.Find(): %w", err)
	}

	return user, nil
}
