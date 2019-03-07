package helpers

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

func InitRedis(redisAddress, redisNetwork string) *redis.Pool {
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(redisNetwork, redisAddress)
			if err != nil {
				return nil, err
			}
			return conn, err
		},

		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
	return pool
}
