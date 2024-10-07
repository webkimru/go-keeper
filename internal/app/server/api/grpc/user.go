package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/webkimru/go-keeper/internal/app/server/api/grpc/pb"
	"github.com/webkimru/go-keeper/internal/app/server/models"
	"github.com/webkimru/go-keeper/pkg/errs"
)

// UserService is an interface to store users.
type UserService interface {
	Add(ctx context.Context, user *models.User) error
	Find(ctx context.Context, login, password string) (string, error)
}

// UserServer is the server for the authentication.
type UserServer struct {
	userService UserService
	// Must be embedded to have forward compatible implementations
	pb.UnimplementedUserServiceServer
}

// NewUserServer returns a new authentication server.
func NewUserServer(userService UserService) *UserServer {
	return &UserServer{userService: userService}
}

// Login is a unary RPC to login user.
// @Success  0 OK              status & json
// @Failure  3 InvalidArgument status
// @Failure  5 NotFound        status
// @Failure 16 Unauthenticated status
// @Failure 13 Internal        status
func (s *UserServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := s.userService.Find(ctx, in.GetLogin(), in.GetPassword())
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, errs.MsgNotFound)
		}
		if errors.Is(err, errs.ErrBadRequest) {
			return nil, status.Errorf(codes.InvalidArgument, errs.MsgFieldRequiredError(err))
		}
		if errors.Is(err, errs.ErrInvalidCredentials) {
			return nil, status.Errorf(codes.Unauthenticated, errs.MsgInvalidCred)
		}

		return nil, status.Errorf(codes.Internal, errs.MsgInternalServerError(err))
	}

	return &pb.LoginResponse{AccessToken: token}, nil
}

// Register is a unary RPC to creates a new user.
// @Success  0 OK              status & json
// @Failure  3 InvalidArgument status
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
		if errors.Is(err, errs.ErrBadRequest) {
			return nil, status.Errorf(codes.InvalidArgument, errs.MsgFieldRequiredError(err))
		}

		return nil, status.Errorf(codes.Internal, errs.MsgInternalServerError(err))
	}

	return &response, nil
}
