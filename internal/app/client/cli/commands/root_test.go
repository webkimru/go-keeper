package commands_test

import (
	"context"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/webkimru/go-keeper/internal/app/client/cli/commands"
	"github.com/webkimru/go-keeper/internal/app/client/cli/grpc"
	"github.com/webkimru/go-keeper/internal/app/client/config"
	"github.com/webkimru/go-keeper/internal/app/client/service"
	"github.com/webkimru/go-keeper/internal/app/client/service/mocks"
	"github.com/webkimru/go-keeper/pkg/crypt"
	"github.com/webkimru/go-keeper/pkg/jwtmanager"
	"github.com/webkimru/go-keeper/pkg/logger"
)

func testUserService(t *testing.T, ctrl *gomock.Controller, cfg *config.Config, l *logger.Log) *service.UserService {
	t.Helper()

	gRPC := grpc.NewClient(cfg, l)                     // gRPC client with tcp, host, port
	userClient := grpc.NewUserClient(gRPC.Client, cfg) // gRPC unary user client

	// userService with dependencies:
	// jwtManager saving a token to the app config and reusing it in the commands
	jwtManager := jwtmanager.New(cfg.SecretKey, cfg.TokenExp)
	// userService business logic layer above the commands
	userService := service.NewUserService(
		mocks.NewMockUserStore(ctrl),
		userClient,
		cfg,
		jwtManager,
		l,
	)

	return userService
}

func TestRootCommand(t *testing.T) {
	t.Run("root command", func(t *testing.T) {
		ctx := context.Background()

		// mockController
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// gRPC client with dependencies
		l, _ := logger.NewZap()
		cfg, _ := config.New()

		userService := testUserService(t, ctrl, cfg, l)

		// keyValueService with dependencies:
		// cryptManager to encrypt local key-value data
		cryptManager, _ := crypt.New(cfg.SecretKey)
		keyValueService := service.NewKeyValueService(
			mocks.NewMockKeyValueStore(ctrl),
			cryptManager,
			l,
		)

		cmd := commands.RootCommand(ctx, strings.NewReader(""), userService, keyValueService, cfg, l)
		err := cmd.Execute()
		assert.NoError(t, err)
	})
}
