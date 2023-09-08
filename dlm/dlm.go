package dlm

import (
	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
)

type dlmClient struct {
	*redislock.Client
}

func New(address string, db int) (*dlmClient, *redis.Client) {
	rdbClient := redis.NewClient(&redis.Options{
		Addr: address,
		DB:   db,
	})
	dlmc := new(dlmClient)
	dlmc.Client = redislock.New(rdbClient)
	return dlmc, rdbClient
}
