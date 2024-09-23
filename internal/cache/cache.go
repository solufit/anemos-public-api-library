package cache

import (
	"github.com/redis/go-redis/v9"
)

// Anemosにおいて、Cacheを扱うための構造体
type AnemosCache struct {
	redisClient *redis.Client
}

// AnemosCacheのコンストラクタ
func NewAnemosCache(redisOption *redis.Options) *AnemosCache {

	redisClient := redis.NewClient(redisOption)
	return &AnemosCache{redisClient: redisClient}
}
