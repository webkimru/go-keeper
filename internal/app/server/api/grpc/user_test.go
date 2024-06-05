package grpc

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	pb "github.com/webkimru/go-keeper/internal/app/server/api/grpc/proto"
	"github.com/webkimru/go-keeper/internal/app/server/models"
	"github.com/webkimru/go-keeper/internal/app/server/repository/store/inmemory"
	"github.com/webkimru/go-keeper/internal/app/server/service"
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
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Register(context.Background(), tt.in)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantRes, got)
			e, ok := status.FromError(err)
			assert.True(t, ok)
			assert.Equal(t, tt.wantCode, e.Code())
		})
	}
}

func TestUserServer_Login(t *testing.T) {
	nousers := service.NewUserService(inmemory.NewStorage())
	withuser := service.NewUserService(inmemory.NewStorage())
	err := withuser.Add(context.Background(), &models.User{
		Login:    "test",
		Password: "test",
	})
	assert.NoError(t, err)

	tests := []struct {
		name     string
		in       *pb.LoginRequest
		store    *service.UserService
		wantCode codes.Code
		wantRes  *pb.LoginResponse
	}{
		{
			"negative: user doesn't exist",
			&pb.LoginRequest{
				Login:    "test",
				Password: "test",
			},
			nousers,
			codes.Unauthenticated,
			&pb.LoginResponse{},
		},
		{
			"positive: correct cred",
			&pb.LoginRequest{
				Login:    "test",
				Password: "test",
			},
			withuser,
			codes.OK,
			&pb.LoginResponse{},
		},
		{
			"positive: wrong pass",
			&pb.LoginRequest{
				Login:    "test",
				Password: "test123",
			},
			withuser,
			codes.Unauthenticated,
			&pb.LoginResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserServer{
				userService: tt.store,
				jwtManager:  jwtmanager.New("secret", 1),
			}

			got, err := s.Login(context.Background(), tt.in)
			if tt.wantCode == codes.OK {
				tt.wantRes.AccessToken = got.AccessToken
				assert.Equal(t, tt.wantRes, got)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
			e, ok := status.FromError(err)
			assert.True(t, ok)
			assert.Equal(t, tt.wantCode, e.Code())
		})
	}
}
