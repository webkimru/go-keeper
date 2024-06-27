package data_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/webkimru/go-keeper/internal/app/client/cli/commands/data"
	"github.com/webkimru/go-keeper/internal/app/client/service"
	"github.com/webkimru/go-keeper/pkg/crypt"
)

func TestNewKeyValueAddCommand(t *testing.T) {
	t.Run("key-value add command", func(t *testing.T) {
		ctx, m, cfg, l := setupTest(t)

		// keyValueService with dependencies:
		// cryptManager to encrypt local key-value data
		cryptManager, _ := crypt.New(cfg.SecretKey)
		keyValueService := service.NewKeyValueService(m, cryptManager, l)

		cmd := data.NewKeyValueAddCommand(ctx, strings.NewReader("title\nkey\nvalue"), keyValueService, l)
		err := cmd.Execute()
		assert.NoError(t, err)
	})
}
