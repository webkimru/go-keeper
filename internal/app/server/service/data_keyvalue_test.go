package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/webkimru/go-keeper/internal/app/server/models"
	"github.com/webkimru/go-keeper/internal/app/server/service/mocks"
	"github.com/webkimru/go-keeper/pkg/crypt"
	"github.com/webkimru/go-keeper/pkg/errs"
	"testing"
)

func testSetup(t *testing.T) (ctx context.Context, m *mocks.MockKeyValueStore) {
	t.Helper()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m = mocks.NewMockKeyValueStore(ctrl)

	ctx = context.Background()
	ctx = context.WithValue(ctx, "userID", int64(1))

	return ctx, m
}

func TestKeyValueService_Add(t *testing.T) {
	ctx, m := testSetup(t)
	cryptManager, err := crypt.New("secret")
	assert.NoError(t, err)

	// case 3
	ctxErrStore := context.WithValue(ctx, "error", "an error")
	errStore := errors.New("an error")

	m.EXPECT().Add(ctx, gomock.Any()).Return(nil)
	m.EXPECT().Add(ctx, models.KeyValue{Title: ""}).Return(errs.ErrBadRequest)
	m.EXPECT().Add(ctxErrStore, gomock.Any()).Return(errStore)

	tests := []struct {
		name    string
		ctx     context.Context
		model   models.KeyValue
		wantErr bool
	}{
		{"positive: correct data", ctx, models.KeyValue{Title: "title", Key: "key", Value: "value"}, false},
		{"negative: empty data", ctx, models.KeyValue{}, true},
		{"negative: an error", ctxErrStore, models.KeyValue{Title: "title", Key: "key", Value: "value"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &KeyValueService{
				storage:      m,
				cryptManager: cryptManager,
			}
			if err := s.Add(tt.ctx, tt.model); (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}

func TestKeyValueService_Get(t *testing.T) {
	ctx, m := testSetup(t)
	cryptManager, err := crypt.New("secret")
	assert.NoError(t, err)

	ctxPermissionDenied := context.WithValue(ctx, "userID", int64(0))
	errCustom := errors.New("an error")

	m.EXPECT().Get(ctx, int64(1)).Return(&models.KeyValue{
		UserID: 1,
		Key:    "df0c76be5b07aee90dd132c0103722ebb99c60c6a9",
		Value:  "ce0968a45ddae7a989ddcca21f474879ab7001e8",
	}, nil)
	m.EXPECT().Get(ctx, int64(0)).Return(&models.KeyValue{UserID: 1}, errCustom)
	m.EXPECT().Get(ctxPermissionDenied, int64(-2)).Return(&models.KeyValue{UserID: 1}, nil)
	m.EXPECT().Get(ctx, int64(1)).Return(&models.KeyValue{UserID: 1}, nil)
	m.EXPECT().Get(ctx, int64(2)).Return(&models.KeyValue{UserID: 1, Key: "df0c76be5b07aee90dd132c0103722ebb99c60c6a9"}, nil)

	tests := []struct {
		name    string
		ctx     context.Context
		id      int64
		Error   error
		wantErr bool
	}{
		{"positive: correct data", ctx, 1, nil, false},
		{"negative: get storage", ctx, 0, nil, true},
		{"negative: permission denied", ctxPermissionDenied, -2, errs.ErrPermissionDenied, true},
		{"negative: decrypt key", ctx, 1, errCustom, true},
		{"negative: decrypt value", ctx, 2, errCustom, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &KeyValueService{
				storage:      m,
				cryptManager: cryptManager,
			}
			_, err := s.Get(tt.ctx, tt.id)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.Error)
			}
		})
	}
}

func TestKeyValueService_List(t *testing.T) {
	ctx, m := testSetup(t)
	cryptManager, err := crypt.New("secret")
	assert.NoError(t, err)

	errCustom := errors.New("an error")

	m.EXPECT().List(ctx, int64(1), int64(1), int64(0)).Return([]models.KeyValue{
		{
			ID:    1,
			Key:   "df0c76be5b07aee90dd132c0103722ebb99c60c6a9",
			Value: "ce0968a45ddae7a989ddcca21f474879ab7001e8",
		},
	}, nil)
	m.EXPECT().List(ctx, int64(1), int64(0), int64(0)).Return(nil, nil)
	m.EXPECT().List(ctx, int64(1), int64(2), int64(0)).Return(nil, errCustom)
	m.EXPECT().List(ctx, int64(1), int64(3), int64(0)).Return([]models.KeyValue{{ID: 1}}, nil)
	m.EXPECT().List(ctx, int64(1), int64(4), int64(0)).Return([]models.KeyValue{
		{
			ID:  1,
			Key: "df0c76be5b07aee90dd132c0103722ebb99c60c6a9",
		},
	}, nil)

	tests := []struct {
		name    string
		limit   int64
		Error   error
		wantErr bool
	}{
		{"positive: correct data", 1, nil, false},
		{"negative: arg limit 0", 0, errs.ErrBadRequest, true},
		{"negative: store error", 2, errCustom, true},
		{"negative: key encrypt", 3, errCustom, true},
		{"negative: value encrypt", 4, errCustom, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &KeyValueService{
				storage:      m,
				cryptManager: cryptManager,
			}
			_, err := s.List(ctx, 1, tt.limit, 0)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.Error)
			}
		})
	}
}

func TestKeyValueService_Update(t *testing.T) {
	ctx, m := testSetup(t)
	cryptManager, err := crypt.New("secret")
	assert.NoError(t, err)

	errStore := errors.New("an error")
	ctxErrStore := context.WithValue(ctx, "error", "an error")

	m.EXPECT().Update(ctx, gomock.Any()).Return(nil)
	m.EXPECT().Update(ctx, models.KeyValue{}).Return(nil)
	m.EXPECT().Update(ctxErrStore, gomock.Any()).Return(errStore)

	tests := []struct {
		name    string
		ctx     context.Context
		model   models.KeyValue
		wantErr bool
	}{
		{
			"positive: correct data",
			ctx,
			models.KeyValue{
				ID:    1,
				Title: "title",
				Key:   "key",
				Value: "value",
			},
			false,
		},
		{"negative: empty data", ctx, models.KeyValue{}, true},
		{
			"negative: store error",
			ctxErrStore,
			models.KeyValue{
				ID:    1,
				Title: "title",
				Key:   "key",
				Value: "value",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &KeyValueService{
				storage:      m,
				cryptManager: cryptManager,
			}
			err := s.Update(tt.ctx, tt.model)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}
		})
	}

}

func TestKeyValueService_Delete(t *testing.T) {
	ctx, m := testSetup(t)
	cryptManager, err := crypt.New("secret")
	assert.NoError(t, err)

	errStore := errors.New("an error")

	m.EXPECT().Delete(ctx, int64(1), int64(1)).Return(nil)
	m.EXPECT().Delete(ctx, int64(1), int64(0)).Return(errStore)

	tests := []struct {
		name    string
		id      int64
		wantErr bool
	}{
		{"positive: correct data", 1, false},
		{"negative: store error", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &KeyValueService{
				storage:      m,
				cryptManager: cryptManager,
			}
			err := s.Delete(ctx, 1, tt.id)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}
