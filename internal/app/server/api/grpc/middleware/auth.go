package middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/webkimru/go-keeper/internal/app/server/models"
	"github.com/webkimru/go-keeper/pkg/jwtmanager"
)

type AuthInterceptor struct {
	jwtManager     *jwtmanager.JWTManager
	accessiblePath map[string]struct{}
}

func NewAuthInterceptor(j *jwtmanager.JWTManager) *AuthInterceptor {
	return &AuthInterceptor{
		j,
		map[string]struct{}{
			"/kim.gokeeper.UserService/Login":    {},
			"/kim.gokeeper.UserService/Register": {},
		},
	}
}

func (u *AuthInterceptor) UnaryAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var token string

	if _, ok := u.accessiblePath[info.FullMethod]; ok {
		return handler(ctx, req)
	}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get("authorization")
		if len(values) > 0 {
			token = values[0]
		}
	}
	if len(token) == 0 {
		return nil, status.Error(codes.Unauthenticated, "Missing token")
	}

	userID := u.jwtManager.GetUserID(token)
	if userID == -1 {
		return nil, status.Error(codes.Unauthenticated, "Invalid token")
	}

	ctx = context.WithValue(ctx, models.ContextKey("userID"), userID)

	return handler(ctx, req)
}
