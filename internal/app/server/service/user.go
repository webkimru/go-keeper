package service

import (
	"context"
	"fmt"
	"github.com/webkimru/go-keeper/internal/app/server/models"
	"github.com/webkimru/go-keeper/pkg/errs"
	"github.com/webkimru/go-keeper/pkg/jwtmanager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//go:generate mockgen -destination=mocks/mock_user.go -package=mocks github.com/webkimru/go-keeper/internal/app/server/service UserStore

// UserStore is an interface to store users.
type UserStore interface {
	Add(ctx context.Context, user *models.User) error
	Find(ctx context.Context, login string) (*models.User, error)
}

// UserService contains a user storage.
type UserService struct {
	jwtManager *jwtmanager.JWTManager
	storage    UserStore
}

// NewUserService return a new user service with storage.
func NewUserService(storage UserStore, jwtManager *jwtmanager.JWTManager) *UserService {
	return &UserService{storage: storage, jwtManager: jwtManager}
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
func (s *UserService) Find(ctx context.Context, login, password string) (string, error) {
	model := models.User{Login: login, Password: password}
	if field, err := model.Validate("login", "password"); err != nil {
		return "", fmt.Errorf("UserService - Find - model.Validate(): %w: %s is required", errs.ErrBadRequest, field)
	}

	user, err := s.storage.Find(ctx, login)
	if err != nil {
		return "", fmt.Errorf("UserService - Find - s.storage.Find(): %w", err)
	}

	if user == nil || !user.ValidPassword(password) {
		return "", fmt.Errorf("UserService - Find - user.ValidPassword(): %w", errs.ErrInvalidCredentials)
	}

	token, err := s.jwtManager.BuildJWTString(user.ID)
	if err != nil {
		return "", status.Errorf(codes.Internal, errs.MsgInternalServerError(err))
	}

	return token, nil
}
