package middleware

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/webkimru/go-keeper/pkg/jwtmanager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"reflect"
	"testing"
)

func TestAuthInterceptor_UnaryAuthInterceptor(t *testing.T) {
	handler := func(ctx context.Context, req interface{}) (interface{}, error) { return nil, nil }
	jwtManager := jwtmanager.New("secret", 1)
	token, err := jwtManager.BuildJWTString(1)
	assert.NoError(t, err)
	// valid token
	md := metadata.New(map[string]string{"authorization": token})
	ctxWithValidToken := metadata.NewIncomingContext(context.Background(), md)
	// wrong token
	md = metadata.New(map[string]string{"authorization": "wrongToken"})
	ctxWithWrongToken := metadata.NewIncomingContext(context.Background(), md)

	tests := []struct {
		name       string
		accessPath map[string]struct{}
		actualPath *grpc.UnaryServerInfo
		ctx        context.Context
		want       interface{}
		wantErr    bool
	}{
		{
			"positive: accessible path",
			map[string]struct{}{
				"/kim.gokeeper.UserService/Login": {},
			},
			&grpc.UnaryServerInfo{
				FullMethod: "/kim.gokeeper.UserService/Login",
			},
			context.Background(),
			nil,
			false,
		},
		{
			"positive: correct token",
			map[string]struct{}{
				"/kim.gokeeper.UserService/Login": {},
			},
			&grpc.UnaryServerInfo{
				FullMethod: "/kim.gokeeper.UserService/Register",
			},
			ctxWithValidToken,
			nil,
			false,
		},
		{
			"negative: missing token",
			map[string]struct{}{
				"/kim.gokeeper.UserService/Login": {},
			},
			&grpc.UnaryServerInfo{
				FullMethod: "/kim.gokeeper.UserService/Register",
			},
			context.Background(),
			nil,
			true,
		},
		{
			"negative: wrong token",
			map[string]struct{}{
				"/kim.gokeeper.UserService/Login": {},
			},
			&grpc.UnaryServerInfo{
				FullMethod: "/kim.gokeeper.UserService/Register",
			},
			ctxWithWrongToken,
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &AuthInterceptor{
				jwtManager:     jwtManager,
				accessiblePath: tt.accessPath,
			}
			got, err := u.UnaryAuthInterceptor(tt.ctx, nil, tt.actualPath, handler)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnaryAuthInterceptor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnaryAuthInterceptor() got = %v, want %v", got, tt.want)
			}
		})
	}
}
