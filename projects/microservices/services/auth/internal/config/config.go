// Package config loads the auth service settings from environment variables.
package config

import (
	"os"
	"time"
)

type Config struct {
	Addr       string
	JWTSecret  string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}

func Load() Config {
	return Config{
		Addr:       env("AUTH_ADDR", ":8081"),
		JWTSecret:  env("AUTH_JWT_SECRET", "dev-secret-change-me"),
		AccessTTL:  duration("AUTH_ACCESS_TTL", 15*time.Minute),
		RefreshTTL: duration("AUTH_REFRESH_TTL", 7*24*time.Hour),
	}
}

func env(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return def
}

func duration(key string, def time.Duration) time.Duration {
	if d, err := time.ParseDuration(os.Getenv(key)); err == nil {
		return d
	}
	return def
}
