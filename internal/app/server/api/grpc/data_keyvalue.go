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

// KeyValueService is an interface to store data.
type KeyValueService interface {
	Add(ctx context.Context, model models.KeyValue) error
	Get(ctx context.Context, id int64) (*models.KeyValue, error)
	List(ctx context.Context) ([]models.KeyValue, error)
	Update(ctx context.Context, model models.KeyValue) error
	Delete(ctx context.Context, id int64) error
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

		return nil, status.Errorf(codes.Internal, errs.MsgInternalServer)
	}

	return nil, nil
}

func (s *KeyValueServer) GetKeyValue(ctx context.Context, in *pb.GetKeyValueRequest) (*pb.GetKeyValueResponse, error) {
	data, err := s.keyValueService.Get(ctx, in.GetId())
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, errs.MsgNotFound)
		}

		return nil, status.Errorf(codes.Internal, errs.MsgInternalServer)
	}

	return &pb.GetKeyValueResponse{
		Data: &pb.KeyValue{
			Title: data.Title,
			Key:   data.Key,
			Value: data.Value,
		},
	}, nil
}
