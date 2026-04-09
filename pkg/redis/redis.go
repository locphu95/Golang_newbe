package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Addr     string
	Password string
	Client   *redis.Client
}

func NewRedis(addr, password string) *Redis {
	return &Redis{Addr: addr, Password: password}
}

func (r *Redis) Connect() error {
	r.Client = redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       0,
	})
	return r.Client.Ping(context.Background()).Err()
}

func (r *Redis) Close() error {
	if r.Client != nil {
		return r.Client.Close()
	}
	return nil
}
