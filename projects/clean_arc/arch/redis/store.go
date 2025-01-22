package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host string
	Port uint16
	Pwd  string
	DB   int
}

type Store interface {
	GetInstance() *store
	Connect()
	Disconnect()
}

var _ Store = (*store)(nil)

type store struct {
	*redis.Client
	context context.Context
}

// Connect implements Store.
func (r *store) Connect() {
	log.Println("connecting to redis")
	pong, err := r.Ping(r.context).Result()
	if err != nil {
		log.Fatal("could not connect to redis", err)
	}
	log.Println("connected to Redis:", pong)
}

// Disconnect implements Store.
func (r *store) Disconnect() {
	log.Println("disconnecting redis...")
	err := r.Close()
	if err != nil {
		log.Fatal("error disconnecting redis", err)
	}
	log.Println("disconnected redis")
}

// GetInstance implements Store.
func (r *store) GetInstance() *store {
	return r
}

func NewStore(context context.Context, config *Config) Store {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Pwd,
		DB:       config.DB,
	})
	return &store{
		context: context,
		Client:  client,
	}
}
