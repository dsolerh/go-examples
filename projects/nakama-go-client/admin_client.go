package nkclient

import (
	"context"
	"net/http"
	"sync"
	"time"
)

// AdminClient is a Nakama Console (admin) API client.
//
// An AdminClient is safe for concurrent use; the active AdminSession is
// guarded by an internal mutex.
type AdminClient struct {
	cfg  Config
	http *http.Client

	mu      sync.RWMutex
	session *AdminSession
}

// AdminSession holds the bearer token returned by Console authentication.
type AdminSession struct {
	Token       string
	MFARequired bool
	ExpiresAt   time.Time
}

// Expired reports whether the session token is past its expiry.
func (s *AdminSession) Expired() bool {
	return s != nil && !s.ExpiresAt.IsZero() && time.Now().After(s.ExpiresAt)
}

// Credentials bundles console login fields. MFACode is optional and only
// required when the console user has multi-factor authentication enabled.
type Credentials struct {
	Username string
	Password string
	MFACode  string
}

// NewAdminClient constructs an unauthenticated console AdminClient.
//
// host is the Nakama host without scheme or port. Call Authenticate before
// using any other method. By default the Console port (7351) is used; pass
// WithPort to override.
func NewAdminClient(host string, opts ...Option) (*AdminClient, error) {
	cfg := Config{
		Host:    host,
		Port:    DefaultConsolePort,
		Timeout: DefaultTimeout,
	}
	cfg.Apply(opts...)

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: cfg.Timeout}
	}

	return &AdminClient{
		cfg:  cfg,
		http: httpClient,
	}, nil
}

// Session returns the currently stored session, or nil if Authenticate has
// not been called.
func (c *AdminClient) Session() *AdminSession {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.session
}

// SetSession installs s as the active session. Pass nil to clear it.
func (c *AdminClient) SetSession(s *AdminSession) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.session = s
}

// Authenticate exchanges console credentials for an AdminSession.
//
// On success the session is returned to the caller; install it on the
// AdminClient with SetSession to have it attached to subsequent calls.
func (c *AdminClient) Authenticate(ctx context.Context, creds Credentials) (*AdminSession, error) {
	return nil, ErrNotImplemented
}
