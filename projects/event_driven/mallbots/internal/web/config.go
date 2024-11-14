package web

import (
	"fmt"
)

const (
	DefaultHost = "0.0.0.0"
	DefaultPort = "8080"
)

type WebConfig struct {
	Host string
	Port string
}

func (c WebConfig) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
