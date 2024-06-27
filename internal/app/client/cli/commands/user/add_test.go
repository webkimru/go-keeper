package user_test

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/webkimru/go-keeper/internal/app/client/cli/commands/user"
	"github.com/webkimru/go-keeper/internal/app/client/config"
	"github.com/webkimru/go-keeper/internal/app/client/service/mocks"
	"github.com/webkimru/go-keeper/pkg/logger"
)

func TestNewUserAddCommand(t *testing.T) {
	t.Run("user add command", func(t *testing.T) {
		// mockController
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		m := mocks.NewMockUserStore(ctrl)

		cfg, _ := config.New()
		l, _ := logger.NewZap()
		userService := testUserService(t, m, cfg, l)

		cmd := user.NewUserAddCommand(strings.NewReader("admin\npass"), userService, l)
		err := cmd.Execute()
		assert.NoError(t, err)
	})
}
