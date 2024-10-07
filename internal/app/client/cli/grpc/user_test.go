package grpc

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/webkimru/go-keeper/internal/app/client/config"
	"github.com/webkimru/go-keeper/internal/app/client/models"
	"github.com/webkimru/go-keeper/internal/app/server/api/grpc/pb"
	"github.com/webkimru/go-keeper/pkg/logger"
)

func TestUserClient_Register(t *testing.T) {
	cfg, err := config.New()
	assert.NoError(t, err)

	tests := []struct {
		name    string
		cfg     *config.Config
		user    *models.User
		want    *pb.RegisterResponse
		wantErr bool
	}{
		{"negative: connection error", cfg, &models.User{}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, err := logger.NewZap()
			assert.NoError(t, err)
			gRPC := NewClient(cfg, l)

			c := &UserClient{
				cfg:     cfg,
				service: pb.NewUserServiceClient(gRPC.Client),
			}
			got, err := c.Register(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}
