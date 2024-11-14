package rpc

import (
	"fmt"
)

const (
	DefaultHost = "0.0.0.0"
	DefaultPort = "8085"
)

type RpcConfig struct {
	Host string
	Port string
}

func (c RpcConfig) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
