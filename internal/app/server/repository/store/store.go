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
