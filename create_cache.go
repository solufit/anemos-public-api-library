package libanemos

import (
	"github.com/redis/go-redis/v9"
	"github.com/solufit/anemos-public-api-library/internal/cache"
)

func CreateCache(client *redis.Client, anemosData []interface{}) error {
	return cache.CreateCache(client, anemosData)
}
