package myredis

import (
	"github.com/go-redis/redis/v8"
)

type RedisStore struct {
	Db *redis.Client
}

func NewStore(db *redis.Client) RedisStore {
	return RedisStore{
		Db: db,
	}
}

func GetRedisStore(redisAddr string) RedisStore {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return NewStore(client)
}
