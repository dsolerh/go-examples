// Package token issues and verifies signed JWTs.
package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Kind string

const (
	Access  Kind = "access"
	Refresh Kind = "refresh"
)

var ErrInvalid = errors.New("invalid token")

type Claims struct {
	Email string `json:"email"`
	Kind  Kind   `json:"kind"`
	jwt.RegisteredClaims
}

type Manager struct {
	secret     []byte
	accessTTL  time.Duration
	refreshTTL time.Duration
	now        func() time.Time
}

func NewManager(secret string, accessTTL, refreshTTL time.Duration) *Manager {
	return &Manager{
		secret:     []byte(secret),
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
		now:        time.Now,
	}
}

func (m *Manager) Issue(userID, email string, kind Kind) (string, error) {
	ttl := m.accessTTL
	if kind == Refresh {
		ttl = m.refreshTTL
	}
	now := m.now()
	claims := Claims{
		Email: email,
		Kind:  kind,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			Issuer:    "auth-service",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(m.secret)
}

// Parse checks the signature, expiry and kind of a token.
func (m *Manager) Parse(raw string, want Kind) (Claims, error) {
	var c Claims
	_, err := jwt.ParseWithClaims(raw, &c, func(t *jwt.Token) (any, error) {
		return m.secret, nil
	},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithIssuer("auth-service"),
		jwt.WithTimeFunc(m.now),
	)
	if err != nil {
		return Claims{}, fmt.Errorf("%w: %v", ErrInvalid, err)
	}
	if c.Kind != want {
		return Claims{}, fmt.Errorf("%w: expected %s token", ErrInvalid, want)
	}
	return c, nil
}
