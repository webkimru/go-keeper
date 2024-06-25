package grpc

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	"github.com/webkimru/go-keeper/internal/app/client/config"
	"github.com/webkimru/go-keeper/internal/app/client/models"
	"github.com/webkimru/go-keeper/internal/app/server/api/grpc/pb"
)

// UserClient is the client for the authentication.
type UserClient struct {
	cfg     *config.Config
	service pb.UserServiceClient
}

// NewUserClient returns a new authentication client.
func NewUserClient(cc *grpc.ClientConn, cfg *config.Config) *UserClient {
	return &UserClient{
		cfg:     cfg,
		service: pb.NewUserServiceClient(cc),
	}
}

// Register is a unary RPC request to create a new user in the server.
func (c *UserClient) Register(user *models.User) (*pb.RegisterResponse, error) {
	req := &pb.RegisterRequest{
		Login:    user.Login,
		Password: user.Password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.cfg.GRPC.QueryTimeout)*time.Second)
	defer cancel()

	res, err := c.service.Register(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("grpc - UserClient - Register - c.service.Register(): %w", err)
	}

	return res, nil
}
