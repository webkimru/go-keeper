package inmemory

import (
	"context"
	"sync"

	"github.com/webkimru/go-keeper/internal/app/server/models"
	"github.com/webkimru/go-keeper/pkg/errs"
)

// UserStorage contains users in memory.
type UserStorage struct {
	m     sync.RWMutex
	users map[string]*models.User
}

// NewUserStorage returns a new storage for users.
func NewUserStorage() *UserStorage {
	return &UserStorage{users: make(map[string]*models.User)}
}

// Add creates a new user.
func (s *UserStorage) Add(ctx context.Context, user *models.User) error {
	s.m.Lock()
	defer s.m.Unlock()

	_, exist := s.users[user.Login]
	if exist {
		return errs.ErrAlreadyExists
	}

	user.ID = int64(len(s.users)) + 1
	s.users[user.Login] = user

	return nil
}

// Find look for the user.
func (s *UserStorage) Find(ctx context.Context, login string) (*models.User, error) {
	s.m.RLock()
	defer s.m.RUnlock()

	_, exist := s.users[login]
	if !exist {
		return nil, errs.ErrNotFound
	}

	return s.users[login], nil
}
