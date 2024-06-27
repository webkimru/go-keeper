package data_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/webkimru/go-keeper/internal/app/client/config"
	"github.com/webkimru/go-keeper/internal/app/client/service/mocks"
	"github.com/webkimru/go-keeper/pkg/logger"
)

func setupTest(t *testing.T) (context.Context, *mocks.MockKeyValueStore, *config.Config, *logger.Log) {
	t.Helper()

	ctx := context.Background()

	cfg, _ := config.New()
	l, _ := logger.NewZap()

	// mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockKeyValueStore(ctrl)

	return ctx, m, cfg, l
}
