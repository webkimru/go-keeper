package grpc

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/webkimru/go-keeper/internal/app/server/api/grpc/pb"
	"github.com/webkimru/go-keeper/internal/app/server/models"
	"github.com/webkimru/go-keeper/pkg/crypt"
	"github.com/webkimru/go-keeper/pkg/errs"
)

// KeyValueService is an interface to store data.
type KeyValueService interface {
	Add(ctx context.Context, model models.KeyValue) error
	Get(ctx context.Context, id int64) (*models.KeyValue, error)
	List(ctx context.Context, limit, offset int64) ([]models.KeyValue, error)
	Update(ctx context.Context, model models.KeyValue) error
	Delete(ctx context.Context, id int64) error
}

// KeyValueServer is the server for data.
type KeyValueServer struct {
	keyValueService KeyValueService
	cryptManager    *crypt.Crypt
	// Must be embedded to have forward compatible implementations
	pb.UnimplementedKeyValueServiceServer
}

// NewKeyValueServer returns a new data server.
func NewKeyValueServer(keyValueService KeyValueService, cryptManager *crypt.Crypt) *KeyValueServer {
	return &KeyValueServer{keyValueService: keyValueService, cryptManager: cryptManager}
}

// AddKeyValue saves data to the store.
// @Success  0 OK              status & json
// @Failure  3 InvalidArgument status
// @Failure  6 AlreadyExists   status
// @Failure 13 Internal        status
func (s *KeyValueServer) AddKeyValue(ctx context.Context, in *pb.AddKeyValueRequest) (*pb.AddKeyValueResponse, error) {
	data := &models.KeyValue{
		UserID: (ctx.Value("userID")).(int64),
		Title:  in.GetData().GetTitle(),
		Key:    in.GetData().GetKey(),
		Value:  in.GetData().GetValue(),
	}

	if err := s.keyValueService.Add(ctx, *data); err != nil {
		if errors.Is(err, errs.ErrAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, errs.MsgAlreadyExists)
		}
		if errors.Is(err, errs.ErrBadRequest) {
			return nil, status.Errorf(codes.InvalidArgument, s.fieldMessage(errs.MsgFieldRequired, err))
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

// fieldMessage do user-friendly errors of form fields.
//
//	example: field key is required
//	example: field value is required
//	instead: field is required
func (s *KeyValueServer) fieldMessage(mess string, err error) string {
	field := strings.Split(err.Error(), ":")

	return strings.Replace(mess, "field", fmt.Sprintf("field %s", field[0]), 1)
}

// UpdateKeyValue updates data in the store.
// @Success  0 OK               status & json
// @Failure  3 InvalidArgument  status
// @Failure  5 NotFound         status
// @Failure  7 PermissionDenied status
// @Failure 13 Internal         status
func (s *KeyValueServer) UpdateKeyValue(ctx context.Context, in *pb.UpdateKeyValueRequest) (*pb.UpdateKeyValueResponse, error) {
	data := &models.KeyValue{
		ID:     in.GetId(),
		UserID: (ctx.Value("userID")).(int64),
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
			return nil, status.Errorf(codes.InvalidArgument, s.fieldMessage(errs.MsgFieldRequired, err))
		}

		return nil, status.Errorf(codes.Internal, errs.MsgInternalServerError(err))
	}

	return &pb.UpdateKeyValueResponse{}, nil
}

// ListKeyValue updates data in the store.
// @Success  0 OK               status & json
// @Failure  3 InvalidArgument  status
// @Failure 13 Internal         status
func (s *KeyValueServer) ListKeyValue(ctx context.Context, in *pb.ListKeyValueRequest) (*pb.ListKeyValueResponse, error) {
	data, err := s.keyValueService.List(ctx, in.GetLimit(), in.GetOffset())
	if err != nil {
		if errors.Is(err, errs.ErrBadRequest) {
			return nil, status.Errorf(codes.InvalidArgument, s.fieldMessage(errs.MsgFieldRequired, err))
		}

		return nil, status.Errorf(codes.Internal, errs.MsgInternalServerError(err))
	}

	var count int32
	var slice []*pb.KeyValue
	for i, item := range data {
		count = int32(i)
		// decrypt
		if item.Key, err = s.Decrypt(item.Key); err != nil {
			return nil, err
		}
		if item.Value, err = s.Decrypt(item.Value); err != nil {
			return nil, err
		}
		// prepare data for the unary response *pb.ListKeyValueResponse
		slice = append(slice, &pb.KeyValue{
			Id:    int32(item.ID),
			Title: item.Title,
			Key:   item.Key,
			Value: item.Value,
		})
	}

	return &pb.ListKeyValueResponse{
		Count: count + 1, // for i - start from 0 => +1
		Data:  slice,
	}, nil
}

// DelKeyValue deletes data in the store.
// @Success  0 OK               status & json
// @Failure  3 InvalidArgument  status
// @Failure 13 Internal         status
func (s *KeyValueServer) DelKeyValue(ctx context.Context, in *pb.DelKeyValueRequest) (*pb.DelKeyValueResponse, error) {
	if err := s.keyValueService.Delete(ctx, in.GetId()); err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, errs.MsgNotFound)
		}
		return nil, status.Errorf(codes.Internal, errs.MsgInternalServerError(err))
	}

	return &pb.DelKeyValueResponse{}, nil
}

// Decrypt decrypts fields.
func (s *KeyValueServer) Decrypt(field string) (string, error) {
	decrypted, err := s.cryptManager.Decrypt(field)
	if err != nil {
		return "", fmt.Errorf("KeyValueServer - ListKeyValue - s.cryptManager.Decrypt(%s): %w", field, err)
	}

	return decrypted, nil
}
