package nkclient

import (
	"errors"
	"net/http"
	"time"
)

// ErrNotImplemented is returned by stub methods that have not yet been wired
// up to a real transport.
var ErrNotImplemented = errors.New("nkclient: not implemented")

// Default endpoint values matching Nakama's out-of-the-box ports.
const (
	DefaultHost        = "127.0.0.1"
	DefaultAPIPort     = 7350
	DefaultConsolePort = 7351
	DefaultTimeout     = 10 * time.Second
)

// Config holds the connection parameters shared by user and admin clients.
//
// A Config is normally built by passing Options to a client constructor; it
// may also be constructed directly when the caller wants to keep the values
// in a single struct (e.g. loaded from a config file).
type Config struct {
	// Host is the Nakama server hostname or IP (no scheme, no port).
	Host string
	// Port is the TCP port to dial.
	Port int
	// UseSSL toggles between http:// and https:// for the REST transport.
	UseSSL bool
	// Timeout is the per-request timeout applied to HTTP calls.
	Timeout time.Duration
	// HTTPClient overrides the underlying http.Client. When nil a client
	// configured from Timeout is used.
	HTTPClient *http.Client
}

// Option mutates a Config. Constructors accept a variadic list of Options
// after their required positional arguments.
type Option func(*Config)

// WithHost overrides Config.Host.
func WithHost(host string) Option {
	return func(c *Config) { c.Host = host }
}

// WithPort overrides Config.Port.
func WithPort(port int) Option {
	return func(c *Config) { c.Port = port }
}

// WithSSL toggles TLS on the transport.
func WithSSL(useSSL bool) Option {
	return func(c *Config) { c.UseSSL = useSSL }
}

// WithTimeout sets the per-request timeout.
func WithTimeout(d time.Duration) Option {
	return func(c *Config) { c.Timeout = d }
}

// WithHTTPClient injects a caller-managed *http.Client. When supplied, the
// client's own Timeout is honored and Config.Timeout is ignored.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Config) { c.HTTPClient = hc }
}

// Apply runs the options against c in order and returns c for chaining.
func (c *Config) Apply(opts ...Option) *Config {
	for _, opt := range opts {
		if opt != nil {
			opt(c)
		}
	}
	return c
}
