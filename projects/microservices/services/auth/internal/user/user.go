// Package user holds the user model and its storage.
package user

import (
	"context"
	"errors"
	"sync"
	"time"
)

var (
	ErrNotFound      = errors.New("user not found")
	ErrAlreadyExists = errors.New("user already exists")
)

type User struct {
	ID           string
	Email        string
	PasswordHash []byte
	CreatedAt    time.Time
}

// Repository is the storage contract. Swap MemoryRepository for a
// Postgres implementation later without touching the service layer.
type Repository interface {
	Create(ctx context.Context, u User) error
	ByEmail(ctx context.Context, email string) (User, error)
	ByID(ctx context.Context, id string) (User, error)
}

type MemoryRepository struct {
	mu      sync.RWMutex
	byID    map[string]User
	byEmail map[string]string
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		byID:    make(map[string]User),
		byEmail: make(map[string]string),
	}
}

func (r *MemoryRepository) Create(_ context.Context, u User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.byEmail[u.Email]; ok {
		return ErrAlreadyExists
	}
	r.byID[u.ID] = u
	r.byEmail[u.Email] = u.ID
	return nil
}

func (r *MemoryRepository) ByEmail(_ context.Context, email string) (User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	id, ok := r.byEmail[email]
	if !ok {
		return User{}, ErrNotFound
	}
	return r.byID[id], nil
}

func (r *MemoryRepository) ByID(_ context.Context, id string) (User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.byID[id]
	if !ok {
		return User{}, ErrNotFound
	}
	return u, nil
}
