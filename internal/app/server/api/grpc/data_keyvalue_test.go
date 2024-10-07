package grpc

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/webkimru/go-keeper/internal/app/server/api/grpc/mocks"
	"github.com/webkimru/go-keeper/internal/app/server/api/grpc/pb"
	"github.com/webkimru/go-keeper/internal/app/server/models"
	"github.com/webkimru/go-keeper/pkg/errs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func testSetup(t *testing.T) (ctx context.Context, m *mocks.MockKeyValueService) {
	t.Helper()

	// create controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// create a stub-store
	m = mocks.NewMockKeyValueService(ctrl)

	ctx = context.Background()
	ctx = context.WithValue(ctx, models.ContextKey("userID"), int64(1))

	return ctx, m
}

func TestKeyValueServer_DelKeyValue(t *testing.T) {
	ctx, m := testSetup(t)

	// conditions after run s.DelKeyValue()
	m.EXPECT().Delete(ctx, int64(1), int64(1)).Return(nil)
	m.EXPECT().Delete(ctx, int64(1), int64(0)).Return(errs.ErrNotFound)
	m.EXPECT().Delete(ctx, int64(1), int64(-1)).Return(errs.ErrInternalServer)

	tests := []struct {
		name     string
		in       *pb.DelKeyValueRequest
		wantCode codes.Code
		wantRes  *pb.DelKeyValueResponse
		wantErr  bool
	}{
		{
			"positive: item exists",
			&pb.DelKeyValueRequest{Id: 1},
			codes.OK,
			&pb.DelKeyValueResponse{},
			false,
		},
		{
			"negative: item not exists",
			&pb.DelKeyValueRequest{Id: 0},
			codes.NotFound,
			&pb.DelKeyValueResponse{},
			true,
		},
		{
			"negative: internal server error",
			&pb.DelKeyValueRequest{Id: -1},
			codes.Internal,
			&pb.DelKeyValueResponse{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &KeyValueServer{keyValueService: m}
			_, err := s.DelKeyValue(ctx, tt.in)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}

			e, ok := status.FromError(err)
			assert.True(t, ok)
			assert.Equal(t, tt.wantCode, e.Code())
		})
	}
}

func TestKeyValueServer_AddKeyValue(t *testing.T) {
	ctx, m := testSetup(t)

	m.EXPECT().Add(ctx, models.KeyValue{UserID: 1}).Return(errs.ErrBadRequest)
	m.EXPECT().Add(ctx, models.KeyValue{UserID: 1, Title: "title", Key: "key", Value: "value"}).Return(nil)
	m.EXPECT().Add(ctx, models.KeyValue{UserID: 1, Title: "title2", Key: "key2", Value: "value2"}).Return(errs.ErrAlreadyExists)
	m.EXPECT().Add(ctx, models.KeyValue{UserID: 1, Title: "title3", Key: "key3", Value: "value3"}).Return(errs.ErrInternalServer)

	tests := []struct {
		name     string
		in       *pb.AddKeyValueRequest
		wantCode codes.Code
		wantRes  *pb.AddKeyValueResponse
		wantErr  bool
	}{
		{
			"positive: correct data",
			&pb.AddKeyValueRequest{
				Data: &pb.KeyValue{
					Title: "title",
					Key:   "key",
					Value: "value",
				},
			},
			codes.OK,
			&pb.AddKeyValueResponse{},
			false,
		},
		{
			"negative: already exists",
			&pb.AddKeyValueRequest{
				Data: &pb.KeyValue{
					Title: "title2",
					Key:   "key2",
					Value: "value2",
				},
			},
			codes.AlreadyExists,
			&pb.AddKeyValueResponse{},
			true,
		},
		{
			"negative: internal server error",
			&pb.AddKeyValueRequest{
				Data: &pb.KeyValue{
					Title: "title3",
					Key:   "key3",
					Value: "value3",
				},
			},
			codes.Internal,
			&pb.AddKeyValueResponse{},
			true,
		},
		{
			"negative: empty data",
			&pb.AddKeyValueRequest{Data: nil},
			codes.InvalidArgument,
			&pb.AddKeyValueResponse{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &KeyValueServer{keyValueService: m}
			_, err := s.AddKeyValue(ctx, tt.in)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}

			e, ok := status.FromError(err)
			assert.True(t, ok)
			assert.Equal(t, tt.wantCode, e.Code())
		})
	}
}

func TestKeyValueServer_GetKeyValue(t *testing.T) {
	ctx, m := testSetup(t)

	m.EXPECT().Get(ctx, int64(1)).Return(&models.KeyValue{
		Title: "title",
		Key:   "key",
		Value: "value",
	}, nil)
	m.EXPECT().Get(ctx, int64(0)).Return(nil, errs.ErrNotFound)
	m.EXPECT().Get(ctx, int64(-1)).Return(nil, errs.ErrInternalServer)
	m.EXPECT().Get(ctx, int64(-2)).Return(nil, errs.ErrPermissionDenied)

	tests := []struct {
		name     string
		in       *pb.GetKeyValueRequest
		wantCode codes.Code
		wantRes  *pb.GetKeyValueResponse
		wantErr  bool
	}{
		{
			"positive: correct data",
			&pb.GetKeyValueRequest{Id: 1},
			codes.OK,
			&pb.GetKeyValueResponse{
				Data: &pb.KeyValue{
					Title: "title",
					Key:   "key",
					Value: "value",
				},
			},
			false,
		},
		{
			"negative: not found",
			&pb.GetKeyValueRequest{Id: 0},
			codes.NotFound,
			nil,
			true,
		},
		{
			"negative: internal server error",
			&pb.GetKeyValueRequest{Id: -1},
			codes.Internal,
			nil,
			true,
		},
		{
			"negative: permission denied",
			&pb.GetKeyValueRequest{Id: -2},
			codes.PermissionDenied,
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &KeyValueServer{keyValueService: m}
			res, err := s.GetKeyValue(ctx, tt.in)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.wantRes, res)

			e, ok := status.FromError(err)
			assert.True(t, ok)
			assert.Equal(t, tt.wantCode, e.Code())
		})
	}

}

func TestKeyValueServer_UpdateKeyValue(t *testing.T) {
	ctx, m := testSetup(t)

	m.EXPECT().Update(ctx, models.KeyValue{ID: 1, UserID: 1, Title: "title", Key: "key", Value: "value"}).Return(nil)
	m.EXPECT().Update(ctx, models.KeyValue{ID: 0, UserID: 1}).Return(errs.ErrNotFound)
	m.EXPECT().Update(ctx, models.KeyValue{ID: -1, UserID: 1}).Return(errs.ErrInternalServer)
	m.EXPECT().Update(ctx, models.KeyValue{ID: -2, UserID: 1}).Return(errs.ErrPermissionDenied)
	m.EXPECT().Update(ctx, models.KeyValue{ID: -3, UserID: 1}).Return(errs.ErrBadRequest)

	tests := []struct {
		name     string
		in       *pb.UpdateKeyValueRequest
		wantCode codes.Code
		wantRes  *pb.UpdateKeyValueResponse
		wantErr  bool
	}{
		{
			"positive: correct data",
			&pb.UpdateKeyValueRequest{
				Id: int64(1),
				Data: &pb.KeyValue{
					Title: "title",
					Key:   "key",
					Value: "value",
				},
			},
			codes.OK,
			&pb.UpdateKeyValueResponse{},
			false,
		},
		{
			"negative: not found",
			&pb.UpdateKeyValueRequest{Id: int64(0)},
			codes.NotFound,
			&pb.UpdateKeyValueResponse{},
			true,
		},
		{
			"negative: internal server error",
			&pb.UpdateKeyValueRequest{Id: int64(-1)},
			codes.Internal,
			&pb.UpdateKeyValueResponse{},
			true,
		},
		{
			"negative: permission denied",
			&pb.UpdateKeyValueRequest{Id: int64(-2)},
			codes.PermissionDenied,
			&pb.UpdateKeyValueResponse{},
			true,
		},
		{
			"negative: invalid argument",
			&pb.UpdateKeyValueRequest{Id: int64(-3)},
			codes.InvalidArgument,
			&pb.UpdateKeyValueResponse{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &KeyValueServer{keyValueService: m}
			_, err := s.UpdateKeyValue(ctx, tt.in)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}

			e, ok := status.FromError(err)
			assert.True(t, ok)
			assert.Equal(t, tt.wantCode, e.Code())
		})
	}
}

func TestKeyValueServer_ListKeyValue(t *testing.T) {
	ctx, m := testSetup(t)

	m.EXPECT().List(ctx, int64(1), int64(1), int64(0)).Return([]models.KeyValue{{
		ID:    1,
		Title: "title",
		Key:   "key",
		Value: "value",
	}}, nil)
	m.EXPECT().List(ctx, int64(1), int64(-1), int64(0)).Return(nil, errs.ErrInternalServer)
	m.EXPECT().List(ctx, int64(1), int64(-3), int64(0)).Return(nil, errs.ErrBadRequest)

	tests := []struct {
		name     string
		in       *pb.ListKeyValueRequest
		wantCode codes.Code
		wantRes  *pb.ListKeyValueResponse
		wantErr  bool
	}{
		{
			"positive: correct list",
			&pb.ListKeyValueRequest{Limit: 1, Offset: 0},
			codes.OK,
			&pb.ListKeyValueResponse{
				Data: []*pb.KeyValue{{
					Id:    1,
					Title: "title",
					Key:   "key",
					Value: "value",
				}},
			},
			false,
		},
		{
			"negative: internal server error",
			&pb.ListKeyValueRequest{Limit: -1, Offset: 0},
			codes.Internal,
			&pb.ListKeyValueResponse{},
			true,
		},
		{
			"negative: invalid argument",
			&pb.ListKeyValueRequest{Limit: -3, Offset: 0},
			codes.InvalidArgument,
			&pb.ListKeyValueResponse{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &KeyValueServer{keyValueService: m}
			_, err := s.ListKeyValue(ctx, tt.in)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}

			e, ok := status.FromError(err)
			assert.True(t, ok)
			assert.Equal(t, tt.wantCode, e.Code())
		})
	}

}
