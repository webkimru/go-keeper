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

//go:generate mockgen -destination=mocks/mock_keyvalue.go -package=mocks github.com/webkimru/go-keeper/internal/app/server/api/grpc KeyValueService

// KeyValueService is an interface to store data.
type KeyValueService interface {
	Add(ctx context.Context, model models.KeyValue) error
	Get(ctx context.Context, id int64) (*models.KeyValue, error)
	List(ctx context.Context, UserID, limit, offset int64) ([]models.KeyValue, error)
	Update(ctx context.Context, model models.KeyValue) error
	Delete(ctx context.Context, UserID, id int64) error
}

// KeyValueServer is the server for data.
type KeyValueServer struct {
	keyValueService KeyValueService
	// Must be embedded to have forward compatible implementations
	pb.UnimplementedKeyValueServiceServer
}

// NewKeyValueServer returns a new data server.
func NewKeyValueServer(keyValueService KeyValueService) *KeyValueServer {
	return &KeyValueServer{keyValueService: keyValueService}
}

// AddKeyValue saves data to the store.
// @Success  0 OK               status & json
// @Failure  3 InvalidArgument  status
// @Failure  6 AlreadyExists    status
// @Failure  7 PermissionDenied status
// @Failure 13 Internal         status
func (s *KeyValueServer) AddKeyValue(ctx context.Context, in *pb.AddKeyValueRequest) (*pb.AddKeyValueResponse, error) {
	userID := s.getContextUserID(ctx)
	if userID == -1 {
		return nil, status.Errorf(codes.PermissionDenied, errs.MsgPermissionDenied)
	}

	data := &models.KeyValue{
		UserID: userID,
		Title:  in.GetData().GetTitle(),
		Key:    in.GetData().GetKey(),
		Value:  in.GetData().GetValue(),
	}

	if err := s.keyValueService.Add(ctx, *data); err != nil {
		if errors.Is(err, errs.ErrAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, errs.MsgAlreadyExists)
		}
		if errors.Is(err, errs.ErrBadRequest) {
			return nil, status.Errorf(codes.InvalidArgument, errs.MsgFieldRequiredError(err))
		}

		return nil, status.Errorf(codes.Internal, errs.MsgInternalServerError(err))
	}

	return &pb.AddKeyValueResponse{}, nil
}

// GetKeyValue get data from the store.
// @Success  0 OK               status & json
// @Failure  5 NotFound         status
// @Failure  7 PermissionDenied status
// @Failure 13 Internal         status
func (s *KeyValueServer) GetKeyValue(ctx context.Context, in *pb.GetKeyValueRequest) (*pb.GetKeyValueResponse, error) {
	data, err := s.keyValueService.Get(ctx, in.GetId())
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, errs.MsgNotFound)
		}
		if errors.Is(err, errs.ErrPermissionDenied) {
			return nil, status.Errorf(codes.PermissionDenied, errs.MsgPermissionDenied)
		}

		return nil, status.Errorf(codes.Internal, errs.MsgInternalServerError(err))
	}

	return &pb.GetKeyValueResponse{
		Data: &pb.KeyValue{
			Title: data.Title,
			Key:   data.Key,
			Value: data.Value,
		},
	}, nil
}

// UpdateKeyValue updates data in the store.
// @Success  0 OK               status & json
// @Failure  3 InvalidArgument  status
// @Failure  5 NotFound         status
// @Failure  7 PermissionDenied status
// @Failure 13 Internal         status
func (s *KeyValueServer) UpdateKeyValue(ctx context.Context, in *pb.UpdateKeyValueRequest) (*pb.UpdateKeyValueResponse, error) {
	userID := s.getContextUserID(ctx)
	if userID == -1 {
		return nil, status.Errorf(codes.PermissionDenied, errs.MsgPermissionDenied)
	}

	data := &models.KeyValue{
		ID:     in.GetId(),
		UserID: userID,
		Title:  in.GetData().GetTitle(),
		Key:    in.GetData().GetKey(),
		Value:  in.GetData().GetValue(),
	}

	err := s.keyValueService.Update(ctx, *data)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, errs.MsgNotFound)
		}
		if errors.Is(err, errs.ErrPermissionDenied) {
			return nil, status.Errorf(codes.PermissionDenied, errs.MsgPermissionDenied)
		}
		if errors.Is(err, errs.ErrBadRequest) {
			return nil, status.Errorf(codes.InvalidArgument, errs.MsgFieldRequiredError(err))
		}

		return nil, status.Errorf(codes.Internal, errs.MsgInternalServerError(err))
	}

	return &pb.UpdateKeyValueResponse{}, nil
}

// ListKeyValue updates data in the store.
// @Success  0 OK               status & json
// @Failure  3 InvalidArgument  status
// @Failure  7 PermissionDenied status
// @Failure 13 Internal         status
func (s *KeyValueServer) ListKeyValue(ctx context.Context, in *pb.ListKeyValueRequest) (*pb.ListKeyValueResponse, error) {
	userID := s.getContextUserID(ctx)
	if userID == -1 {
		return nil, status.Errorf(codes.PermissionDenied, errs.MsgPermissionDenied)
	}

	data, err := s.keyValueService.List(ctx, userID, in.GetLimit(), in.GetOffset())
	if err != nil {
		if errors.Is(err, errs.ErrBadRequest) {
			return nil, status.Errorf(codes.InvalidArgument, errs.MsgFieldRequiredError(err))
		}

		return nil, status.Errorf(codes.Internal, errs.MsgInternalServerError(err))
	}

	n := 0
	slice := make([]*pb.KeyValue, len(data))
	for _, item := range data {
		// prepare data for the unary response *pb.ListKeyValueResponse
		slice[n] = &pb.KeyValue{
			Id:    int32(item.ID),
			Title: item.Title,
			Key:   item.Key,
			Value: item.Value,
		}
		n++
	}

	return &pb.ListKeyValueResponse{
		Count: int32(len(data)),
		Data:  slice,
	}, nil
}

// DelKeyValue deletes data in the store.
// @Success  0 OK               status & json
// @Failure  5 NotFound         status
// @Failure  7 PermissionDenied status
// @Failure 13 Internal         status
func (s *KeyValueServer) DelKeyValue(ctx context.Context, in *pb.DelKeyValueRequest) (*pb.DelKeyValueResponse, error) {
	userID := s.getContextUserID(ctx)
	if userID == -1 {
		return nil, status.Errorf(codes.PermissionDenied, errs.MsgPermissionDenied)
	}

	if err := s.keyValueService.Delete(ctx, userID, in.GetId()); err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, errs.MsgNotFound)
		}
		return nil, status.Errorf(codes.Internal, errs.MsgInternalServerError(err))
	}

	return &pb.DelKeyValueResponse{}, nil
}

// getContextUserID is a helper method to get user ID out of ctx.Value as int64
func (s *KeyValueServer) getContextUserID(ctx context.Context) int64 {
	id := ctx.Value(models.ContextKey("userID"))

	switch id := id.(type) {
	case int64:
		return id

	default:
		return -1
	}
}
