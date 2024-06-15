package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/webkimru/go-keeper/internal/app/server/api/grpc/pb"
	"github.com/webkimru/go-keeper/internal/app/server/models"
	"github.com/webkimru/go-keeper/pkg/errs"
	"github.com/webkimru/go-keeper/pkg/jwtmanager"
)

// UserService is an interface to store users.
type UserService interface {
	Add(ctx context.Context, user *models.User) error
	Find(ctx context.Context, login string) (*models.User, error)
}

// UserServer is the server for the authentication.
type UserServer struct {
	userService UserService
	jwtManager  *jwtmanager.JWTManager
	// Must be embedded to have forward compatible implementations
	pb.UnimplementedUserServiceServer
}

// NewUserServer returns a new authentication server.
func NewUserServer(userService UserService, jwtManager *jwtmanager.JWTManager) *UserServer {
	return &UserServer{userService: userService, jwtManager: jwtManager}
}

// Login is a unary RPC to login user.
// @Success  0 OK              status & json
// @Failure  5 NotFound        status
// @Failure 16 Unauthenticated status
// @Failure 13 Internal        status
func (s *UserServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.userService.Find(ctx, in.GetLogin())
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, errs.MsgNotFound)
		}

		return nil, status.Errorf(codes.Internal, errs.MsgInternalServerError(err))
	}

	if user == nil || !user.ValidPassword(in.GetPassword()) {
		return nil, status.Errorf(codes.Unauthenticated, errs.MsgInvalidCred)
	}

	token, err := s.jwtManager.BuildJWTString(user.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, errs.MsgInternalServerError(err))
	}

	return &pb.LoginResponse{AccessToken: token}, nil
}

// Register is a unary RPC to creates a new user.
// @Success  0 OK              status & json
// @Failure 16 OK              json {"error": "already exists"}
// @Failure 13 Internal        status
func (s *UserServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var response pb.RegisterResponse

	user := &models.User{
		Login:    in.GetLogin(),
		Password: in.GetPassword(),
	}

	err := s.userService.Add(ctx, user)
	if err != nil {
		if errors.Is(err, errs.ErrAlreadyExists) {
			response.Error = errs.MsgAlreadyExists
			return &response, nil
		}

		return nil, status.Errorf(codes.Internal, errs.MsgInternalServerError(err))
	}

	return &response, nil
}
