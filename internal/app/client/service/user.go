package service

import (
	"context"
	"fmt"
	"github.com/webkimru/go-keeper/internal/app/client/config"
	"github.com/webkimru/go-keeper/pkg/errs"

	"github.com/webkimru/go-keeper/internal/app/client/cli/grpc"
	"github.com/webkimru/go-keeper/internal/app/client/models"
	"github.com/webkimru/go-keeper/pkg/grpcserver/client"
	"github.com/webkimru/go-keeper/pkg/logger"
)

// UserStore is an interface to store users.
type UserStore interface {
	Add(ctx context.Context, user *models.User) error
	Get(ctx context.Context, login string) (*models.User, error)
}

// UserService is a service to save data and do gRPC requests.
type UserService struct {
	cfg     *config.Config
	logger  *logger.Log
	gRPC    *client.Client
	storage UserStore
}

// NewUserService returns a new user service with needed options.
func NewUserService(storage UserStore, client *client.Client, cfg *config.Config, l *logger.Log) *UserService {
	return &UserService{storage: storage, gRPC: client, cfg: cfg, logger: l}
}

// Add crates a new user in the local storage SQLite and go-keeper server
func (s *UserService) Add(ctx context.Context, user *models.User) error {
	// 1, save to the local storage - SQLite
	if err := s.storage.Add(ctx, user); err != nil {
		return fmt.Errorf("service - Add - s.storage.Add(): %w", err)
	}

	// 2. put a new user to the go-keeper server
	gRPClient := grpc.NewUserClient(s.gRPC.Client, s.cfg)
	res, err := gRPClient.Register(user)
	if err != nil {
		return fmt.Errorf("service - Add - userClient.Register(): %w", err)
	}
	if res != "" {
		errs.MsgCLI(res)
	}

	return nil
}

func (s *UserService) Get(ctx context.Context, login string) (*models.User, error) {

	return nil, nil
}
