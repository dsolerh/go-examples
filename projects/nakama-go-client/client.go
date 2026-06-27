package nkclient

import (
	"net/http"
	"time"
)

// Client is a Nakama player-facing API client.
//
// A Client is safe for concurrent use; the active Session is guarded by an
// internal mutex.
type Client struct {
	cfg       Config
	serverKey string
	http      *http.Client
}

// Session holds the credentials returned by an authentication call.
type Session struct {
	Token        string
	RefreshToken string
	UserID       string
	Username     string
	Vars         map[string]string
	Created      bool
	ExpiresAt    time.Time
	RefreshAt    time.Time
}

// Expired reports whether the session's access token is past its expiry.
func (s *Session) Expired() bool {
	return s != nil && !s.ExpiresAt.IsZero() && time.Now().After(s.ExpiresAt)
}

// AuthenticateRequest are the common knobs accepted by every authentication
// flow. Provider-specific credentials are passed as the first positional
// argument to each Authenticate* method.
type AuthenticateRequest struct {
	// Create controls whether the server should create a new account when
	// no matching identity is found.
	Create bool
	// Username is the desired username when Create is true. Ignored
	// otherwise.
	Username string
	// Vars are arbitrary session variables passed to server runtime hooks.
	Vars map[string]string
}

// NewClient constructs a player-facing Client.
//
// serverKey is the Nakama server key. host is the Nakama host without
// scheme or port. Remaining knobs are supplied via Option.
func NewClient(serverKey, host string, opts ...Option) (*Client, error) {
	cfg := Config{
		Host:    host,
		Port:    DefaultAPIPort,
		Timeout: DefaultTimeout,
	}
	cfg.Apply(opts...)

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: cfg.Timeout}
	}

	return &Client{
		cfg:       cfg,
		serverKey: serverKey,
		http:      httpClient,
	}, nil
}

