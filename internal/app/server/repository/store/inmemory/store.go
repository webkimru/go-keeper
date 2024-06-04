package inmemory

import (
	"context"
	"fmt"
	"github.com/webkimru/go-keeper/internal/app/server/models"
	"sync"
)

// Storage contains users in memory.
type Storage struct {
	m     sync.RWMutex
	users map[string]*models.User
}

// NewStorage returns a new storage for users.
func NewStorage() *Storage {
	return &Storage{users: make(map[string]*models.User)}
}

// Add creates a new user.
func (s *Storage) Add(ctx context.Context, user *models.User) error {
	s.m.Lock()
	defer s.m.Unlock()

	_, exist := s.users[user.Login]
	if exist {
		return fmt.Errorf("user already exists")
	}

	s.users[user.Login] = user

	return nil
}

// Find look for the user.
func (s *Storage) Find(ctx context.Context, login string) (*models.User, error) {
	s.m.RLock()
	defer s.m.RUnlock()

	_, exist := s.users[login]
	if !exist {
		return nil, fmt.Errorf("user is not found")
	}

	return s.users[login], nil
}
