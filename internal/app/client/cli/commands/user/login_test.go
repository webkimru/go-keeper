package user_test

import (
	"context"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/webkimru/go-keeper/internal/app/client/cli/commands/user"
	"github.com/webkimru/go-keeper/internal/app/client/config"
	"github.com/webkimru/go-keeper/internal/app/client/service/mocks"
	"github.com/webkimru/go-keeper/pkg/logger"
)

func TestNewUserLoginCommand(t *testing.T) {
	ctx := context.Background()

	// mockDB
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockUserStore(ctrl)

	m.EXPECT().Get(ctx, "admin").Return(nil, nil)

	cfg, _ := config.New()
	l, _ := logger.NewZap()
	userService := testUserService(t, m, cfg, l)

	tests := []struct {
		name    string
		creds   string
		wantErr bool
	}{
		{"positive: correct data", "admin\npass", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := user.NewUserLoginCommand(strings.NewReader(tt.creds), userService, l)
			err := cmd.Execute()
			assert.NoError(t, err)
		})
	}
}
