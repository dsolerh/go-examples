// Package service holds the auth business logic. It knows nothing about HTTP.
package service

import (
	"context"
	"crypto/rand"
	"errors"
	"net/mail"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"microservices/auth/internal/token"
	"microservices/auth/internal/user"
)

var (
	ErrInvalidInput       = errors.New("invalid email or password")
	ErrEmailTaken         = errors.New("email already registered")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorized       = errors.New("unauthorized")
)

// bcrypt only accepts passwords up to 72 bytes.
const (
	minPasswordLen = 8
	maxPasswordLen = 72
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Identity struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

type Auth struct {
	users  user.Repository
	tokens *token.Manager
}

func NewAuth(users user.Repository, tokens *token.Manager) *Auth {
	return &Auth{users: users, tokens: tokens}
}

func (a *Auth) Register(ctx context.Context, email, password string) (Identity, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	if _, err := mail.ParseAddress(email); err != nil || len(password) < minPasswordLen || len(password) > maxPasswordLen {
		return Identity{}, ErrInvalidInput
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return Identity{}, err
	}
	u := user.User{ID: rand.Text(), Email: email, PasswordHash: hash, CreatedAt: time.Now()}
	if err := a.users.Create(ctx, u); err != nil {
		if errors.Is(err, user.ErrAlreadyExists) {
			return Identity{}, ErrEmailTaken
		}
		return Identity{}, err
	}
	return Identity{UserID: u.ID, Email: u.Email}, nil
}

func (a *Auth) Login(ctx context.Context, email, password string) (Tokens, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	u, err := a.users.ByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			return Tokens{}, ErrInvalidCredentials
		}
		return Tokens{}, err
	}
	if bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password)) != nil {
		return Tokens{}, ErrInvalidCredentials
	}
	return a.issue(u)
}

func (a *Auth) Refresh(ctx context.Context, refreshToken string) (Tokens, error) {
	c, err := a.tokens.Parse(refreshToken, token.Refresh)
	if err != nil {
		return Tokens{}, ErrUnauthorized
	}
	u, err := a.users.ByID(ctx, c.Subject)
	if err != nil {
		return Tokens{}, ErrUnauthorized
	}
	return a.issue(u)
}

// Validate is what other services (or an API gateway) call to check
// who is behind an access token.
func (a *Auth) Validate(accessToken string) (Identity, error) {
	c, err := a.tokens.Parse(accessToken, token.Access)
	if err != nil {
		return Identity{}, ErrUnauthorized
	}
	return Identity{UserID: c.Subject, Email: c.Email}, nil
}

func (a *Auth) issue(u user.User) (Tokens, error) {
	access, err := a.tokens.Issue(u.ID, u.Email, token.Access)
	if err != nil {
		return Tokens{}, err
	}
	refresh, err := a.tokens.Issue(u.ID, u.Email, token.Refresh)
	if err != nil {
		return Tokens{}, err
	}
	return Tokens{AccessToken: access, RefreshToken: refresh}, nil
}
