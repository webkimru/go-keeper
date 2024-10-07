package data_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/webkimru/go-keeper/internal/app/client/cli/commands/data"
	"github.com/webkimru/go-keeper/internal/app/client/models"
	"github.com/webkimru/go-keeper/internal/app/client/service"
	"github.com/webkimru/go-keeper/pkg/crypt"
)

func TestNewKeyValueListCommand(t *testing.T) {
	ctx, m, cfg, l := setupTest(t)
	ctx = context.WithValue(ctx, models.ContextKey("userID"), int64(1))

	m.EXPECT().List(ctx, int64(1), int64(100), int64(0)).Return([]models.KeyValue{
		{
			ID:    1,
			Key:   "df0c76be5b07aee90dd132c0103722ebb99c60c6a9",
			Value: "ce0968a45ddae7a989ddcca21f474879ab7001e8",
		},
	}, nil)

	// keyValueService with dependencies:
	// cryptManager to encrypt local key-value data
	cryptManager, _ := crypt.New(cfg.SecretKey)
	keyValueService := service.NewKeyValueService(
		m,
		cryptManager,
		l,
	)
	tests := []struct {
		name    string
		ctx     context.Context
		wantErr bool
	}{
		{"positive: correct data", ctx, false},
		{"negative: permission denied", context.Background(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := data.NewKeyValueListCommand(tt.ctx, keyValueService, l)
			err := cmd.Execute()
			assert.NoError(t, err)
		})
	}
}
