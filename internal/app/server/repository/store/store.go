package store

import (
	"context"
	"github.com/webkimru/go-keeper/internal/app/server/models"
)

// UserStore is an interface to store users.
type UserStore interface {
	Add(ctx context.Context, user *models.User) error
	Find(ctx context.Context, login string) (*models.User, error)
}

// KeyValueStore is an interface to store data.
type KeyValueStore interface {
	Add(ctx context.Context, kv models.KeyValue) error
	Get(ctx context.Context, id int64) (*models.KeyValue, error)
	List(ctx context.Context) ([]models.KeyValue, error)
	Update(ctx context.Context, model models.KeyValue) error
	Delete(ctx context.Context, id int64) error
}
