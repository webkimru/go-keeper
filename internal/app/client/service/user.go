package service

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/webkimru/go-keeper/internal/app/client/cli/grpc"
	"github.com/webkimru/go-keeper/internal/app/client/config"
	"github.com/webkimru/go-keeper/internal/app/client/models"
	"github.com/webkimru/go-keeper/pkg/errs"
	"github.com/webkimru/go-keeper/pkg/grpcserver/client"
	"github.com/webkimru/go-keeper/pkg/jwtmanager"
	"github.com/webkimru/go-keeper/pkg/logger"
)

// UserStore is an interface to store users.
type UserStore interface {
	Add(ctx context.Context, user *models.User) error
	Get(ctx context.Context, login string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
}

// UserService is a service to save data and do gRPC requests.
type UserService struct {
	gRPC       *client.Client
	cfg        *config.Config
	logger     *logger.Log
	jwtManager *jwtmanager.JWTManager
	storage    UserStore
}

// NewUserService returns a new user service with needed options.
func NewUserService(
	storage UserStore,
	client *client.Client,
	cfg *config.Config,
	jwtManager *jwtmanager.JWTManager,
	l *logger.Log,
) *UserService {
	return &UserService{storage: storage, gRPC: client, cfg: cfg, logger: l, jwtManager: jwtManager}
}

// Add crates a new user in the local storage SQLite and go-keeper server
func (s *UserService) Add(ctx context.Context, model *models.User) error {
	// 1. Puts a new user to the go-keeper server:
	// do request to the server using the client
	gRPClient := grpc.NewUserClient(s.gRPC.Client, s.cfg)
	res, err := gRPClient.Register(model)
	if err != nil {
		return fmt.Errorf("service - Add - userClient.Register(): %w", err)
	}
	if res.Error != "" {
		return errs.ErrAlreadyExists
	}

	// 2, Save to the SQLite local storage
	// make a hashed password for the user
	user, err := models.NewUser(model.Login, model.Password)
	if err != nil {
		return fmt.Errorf("UserService - Add - models.NewUser(): %w", err)
	}
	// create a new status
	user.Status = models.UserStatePending
	// save to the store
	if err = s.storage.Add(ctx, user); err != nil {
		if errors.Is(err, errs.ErrAlreadyExists) {
			return errs.ErrAlreadyExists
		}
		return fmt.Errorf("service - Add - s.storage.Add(): %w", err)
	}

	// 3. Update user status from PENDING to REGISTERED
	if err = s.storage.Update(ctx, user); err != nil {
		return fmt.Errorf("service - Add - s.storage.Update(): %w", err)
	}

	return nil
}

func (s *UserService) Auth(ctx context.Context, login, password string) (string, error) {
	user, err := s.storage.Get(ctx, login)
	if err != nil {
		return "", fmt.Errorf("UserService - Auth - s.storage.Get(): %w", err)
	}

	if user == nil || !user.ValidPassword(password) {
		return "", fmt.Errorf("UserService - Auth - s.storage.Get(): %w", errs.ErrInvalidCredentials)
	}

	token, err := s.jwtManager.BuildJWTString(user.ID)
	if err != nil {
		return "", fmt.Errorf("UserService - Auth - s.jwtManager.BuildJWTString(): %w", errs.ErrInternalServer)
	}

	return token, nil
}
