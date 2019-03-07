package usecase

import (
	"github.com/enfipy/cronpub/src/helpers"
	"github.com/enfipy/cronpub/src/models"
	"github.com/enfipy/cronpub/src/services/bot"

	"github.com/enfipy/locker"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

type botUsecase struct {
	pool   *redis.Pool
	locker *locker.Locker
}

func NewUsecase(pool *redis.Pool, locker *locker.Locker) bot.Usecase {
	return &botUsecase{
		pool:   pool,
		locker: locker,
	}
}

func (ucs *botUsecase) SavePost(post *models.Post) {
	conn := ucs.pool.Get()
	defer conn.Close()

	key := "posts"
	ucs.locker.Lock(key)
	defer ucs.locker.Unlock(key)

	data, err := post.EncodeBinary()
	helpers.PanicOnError(err)

	_, err = conn.Do("SET", uuid.New(), data)
	helpers.PanicOnError(err)
}

func (ucs *botUsecase) GetRandomPost() *models.Post {
	conn := ucs.pool.Get()
	defer conn.Close()

	key := "posts"
	ucs.locker.RLock(key)
	defer ucs.locker.RUnlock(key)

	res, err := conn.Do("RANDOMKEY")
	helpers.PanicOnError(err)

	if res == nil {
		return nil
	}

	id, err := uuid.ParseBytes(res.([]byte))
	helpers.PanicOnError(err)

	res, err = conn.Do("GET", id)
	helpers.PanicOnError(err)

	post := new(models.Post)
	post.DecodeBinary(res.([]byte))

	return post
}
