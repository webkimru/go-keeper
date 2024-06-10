package grpc

import (
	"context"

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
func (s *UserServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.userService.Find(ctx, in.GetLogin())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, errs.MsgNotFound)
	}

	if user == nil || !user.ValidPassword(in.GetPassword()) {
		return nil, status.Errorf(codes.Unauthenticated, errs.MsgInvalidCred)
	}

	token, err := s.jwtManager.BuildJWTString(user.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, errs.MsgInternalServer)
	}

	return &pb.LoginResponse{AccessToken: token}, nil
}

func (s *UserServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var response pb.RegisterResponse

	user := &models.User{
		Login:    in.GetLogin(),
		Password: in.GetPassword(),
	}
	err := s.userService.Add(ctx, user)
	if err != nil {
		response.Error = errs.MsgAlreadyExists
	}

	return &response, nil
}
