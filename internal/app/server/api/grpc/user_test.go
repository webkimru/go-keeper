package grpc

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	pb "github.com/webkimru/go-keeper/internal/app/server/api/grpc/proto"
	"github.com/webkimru/go-keeper/internal/app/server/repository/store/inmemory"
	"github.com/webkimru/go-keeper/pkg/jwtmanager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestUserServer_Register(t *testing.T) {
	s := &UserServer{
		userService: inmemory.NewStorage(),
		jwtManager:  jwtmanager.New("secret", 1),
	}
	tests := []struct {
		name     string
		in       *pb.RegisterRequest
		wantCode codes.Code
		wantRes  *pb.RegisterResponse
	}{
		{
			"positive: new user",
			&pb.RegisterRequest{
				Login:    "test",
				Password: "test",
			},
			codes.OK,
			&pb.RegisterResponse{},
		},
		{
			"negative: new user",
			&pb.RegisterRequest{
				Login:    "test",
				Password: "test",
			},
			codes.OK,
			&pb.RegisterResponse{Error: fmt.Sprintf("%v", ErrAlreadyExists)},
		},
	}
	for _, tt := range tests {
		tt := tt // fix: in goroutines always runs last test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := s.Register(context.Background(), tt.in)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantRes, got)
			e, ok := status.FromError(err)
			assert.True(t, ok)
			assert.Equal(t, tt.wantCode, e.Code())
		})
	}
}
