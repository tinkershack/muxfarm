package dlm

import (
	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
)

type DLMClient struct {
	*redislock.Client
}

func New(address string, db int) (*DLMClient, *redis.Client) {
	rdbClient := redis.NewClient(&redis.Options{
		Addr: address,
		DB:   db,
	})
	dlmc := new(DLMClient)
	dlmc.Client = redislock.New(rdbClient)
	return dlmc, rdbClient
}
