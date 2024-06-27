package service

import (
	"context"
	"github.com/webkimru/go-keeper/internal/app/client/cli/grpc"
	"github.com/webkimru/go-keeper/internal/app/client/config"
	"github.com/webkimru/go-keeper/pkg/logger"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/webkimru/go-keeper/internal/app/client/models"
	"github.com/webkimru/go-keeper/internal/app/client/service/mocks"
)

func testSetupUserService(t *testing.T) (ctx context.Context, m *mocks.MockUserStore) {
	t.Helper()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m = mocks.NewMockUserStore(ctrl)

	ctx = context.Background()

	return ctx, m
}

func TestUserService_Add(t *testing.T) {
	cfg, err := config.New()
	assert.NoError(t, err)
	cfg.GRPC.QueryTimeout = 10
	ctx, m := testSetupUserService(t)

	errStore := errors.New("an error")
	ctxErr := context.WithValue(ctx, models.ContextKey("error"), "an error")
	ctxErrStore := context.WithValue(ctx, models.ContextKey("error"), "a store error")

	l, err := logger.NewZap()
	assert.NoError(t, err)
	gRPC := grpc.NewClient(cfg, l)                     // gRPC client with tcp, host, port
	userClient := grpc.NewUserClient(gRPC.Client, cfg) // gRPC unary user client

	m.EXPECT().Add(ctx, gomock.Any()).Return(nil)
	m.EXPECT().Add(ctxErr, gomock.Any()).Return(nil)
	m.EXPECT().Add(ctxErrStore, gomock.Any()).Return(errStore)

	tests := []struct {
		name    string
		ctx     context.Context
		model   *models.User
		wantErr bool
	}{
		{
			"positive: correct data",
			ctx,
			&models.User{Login: "login", Password: "pass"},
			false,
		},
		{
			"negative: empty login or password",
			ctxErr,
			&models.User{},
			true,
		},
		{
			"negative: store error",
			ctxErrStore,
			&models.User{Login: "login", Password: "pass"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				storage: m,
				cfg:     cfg,
				gRPC:    userClient,
			}
			if err := s.Add(tt.ctx, tt.model); (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}

func TestUserService_Auth(t *testing.T) {
	ctx, m := testSetupUserService(t)

	errStore := errors.New("an error")

	m.EXPECT().Get(ctx, "login").Return(nil, nil)
	m.EXPECT().Get(ctx, "").Return(nil, errStore)
	m.EXPECT().Get(ctx, "login2").Return(nil, errStore)

	tests := []struct {
		name    string
		ctx     context.Context
		login   string
		wantErr bool
	}{
		{
			"positive: correct data",
			ctx,
			"login",
			false,
		},
		{
			"negative: empty data",
			ctx,
			"",
			true,
		},
		{
			"negative: store error",
			ctx,
			"login2",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				storage: m,
			}
			_, err := s.Auth(tt.ctx, tt.login, "pass")
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}
		})
	}

}
