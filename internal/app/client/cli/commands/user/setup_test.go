package user_test

import (
	"testing"

	"github.com/webkimru/go-keeper/internal/app/client/cli/grpc"
	"github.com/webkimru/go-keeper/internal/app/client/config"
	"github.com/webkimru/go-keeper/internal/app/client/service"
	"github.com/webkimru/go-keeper/pkg/jwtmanager"
	"github.com/webkimru/go-keeper/pkg/logger"
)

func testUserService(t *testing.T, db service.UserStore, cfg *config.Config, l *logger.Log) *service.UserService {
	t.Helper()

	// gRPC client with tcp, host, port
	gRPC := grpc.NewClient(cfg, l)
	// gRPC unary user client
	userClient := grpc.NewUserClient(gRPC.Client, cfg)

	// userService with dependencies:
	// jwtManager saving a token to the app config and reusing it in the commands
	jwtManager := jwtmanager.New(cfg.SecretKey, cfg.TokenExp)
	// userService business logic layer above the commands
	userService := service.NewUserService(db, userClient, cfg, jwtManager, l)

	return userService
}
