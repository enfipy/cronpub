package usecase

import (
	"github.com/enfipy/cronpub/src/services/bot"

	"github.com/gomodule/redigo/redis"
)

type botUsecase struct {
	pool *redis.Pool
}

func NewUsecase(pool *redis.Pool) bot.Usecase {
	return &botUsecase{
		pool: pool,
	}
}
