package data_test

import (
	"context"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/webkimru/go-keeper/internal/app/client/cli/commands/data"
	"github.com/webkimru/go-keeper/internal/app/client/config"
	"github.com/webkimru/go-keeper/internal/app/client/service"
	"github.com/webkimru/go-keeper/internal/app/client/service/mocks"
	"github.com/webkimru/go-keeper/pkg/crypt"
	"github.com/webkimru/go-keeper/pkg/logger"
)

func TestNewKeyValueDelCommand(t *testing.T) {
	t.Run("key-value del command", func(t *testing.T) {
		ctx := context.Background()

		cfg, _ := config.New()
		l, _ := logger.NewZap()

		// mock
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// keyValueService with dependencies:
		// cryptManager to encrypt local key-value data
		cryptManager, _ := crypt.New(cfg.SecretKey)
		keyValueService := service.NewKeyValueService(
			mocks.NewMockKeyValueStore(ctrl),
			cryptManager,
			l,
		)

		cmd := data.NewKeyValueDelCommand(ctx, strings.NewReader("1"), keyValueService, l)
		err := cmd.Execute()
		assert.NoError(t, err)
	})
}
