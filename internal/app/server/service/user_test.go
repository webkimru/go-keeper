package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/webkimru/go-keeper/internal/app/server/models"
	"github.com/webkimru/go-keeper/internal/app/server/service/mocks"
	"testing"
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
	ctx, m := testSetupUserService(t)

	errStore := errors.New("an error")

	ctxErr := context.WithValue(ctx, models.ContextKey("error"), "an error")
	ctxErrStore := context.WithValue(ctx, models.ContextKey("error"), "a store error")

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
			}
			if err := s.Add(tt.ctx, tt.model); (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}

func TestUserService_Find(t *testing.T) {
	ctx, m := testSetupUserService(t)

	m.EXPECT().Find(ctx, "login").Return(&models.User{
		ID:       1,
		Login:    "login",
		Password: "pass",
	}, nil)

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				storage: m,
			}
			_, err := s.Find(tt.ctx, tt.login, "pass")
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}
		})
	}

}
