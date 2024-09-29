package cache

import (
	"context"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/redis/go-redis/v9"
)

// get event key id
func getEventKey(redisClient *redis.Client, key Id) (string, error) {

	ctx := context.Background()

	// キャッシュデータを取得する
	event, err := redisClient.Get(ctx, string(key)).Result()

	if err != nil {
		return "", err
	}

	return event, nil

}

// get EventTypeCache

func getEventTypeCache(redisClient *redis.Client) (mapset.Set[EventType], error) {

	ctx := context.Background()

	// キャッシュデータを取得する
	weekly_data, err := redisClient.SMembers(ctx, "EventType").Result()

	if err != nil {
		return nil, err
	}

	eventTypeSet := mapset.NewSet[EventType]()
	for _, eventTypeStr := range weekly_data {
		eventType := EventType(eventTypeStr) // Assuming EventType can be created from string
		eventTypeSet.Add(eventType)
	}
	return eventTypeSet, nil

}

// get weekly cache
func getWeeklyCache(redisClient *redis.Client, key EventType) (mapset.Set[Id], error) {

	ctx := context.Background()

	// キャッシュデータを取得する
	event, err := redisClient.SMembers(ctx, string(key)).Result()

	if err != nil {
		return nil, err
	}

	eventSet := mapset.NewSet[Id]()
	for _, eventStr := range event {
		event := Id(eventStr) // Assuming Id can be created from string
		eventSet.Add(event)
	}
	return eventSet, nil

}

// get Event Cache
func getEventCache(redisClient *redis.Client, key Id) (string, error) {

	ctx := context.Background()

	// キャッシュデータを取得する
	event, err := redisClient.Get(ctx, string(key)).Result()

	if err != nil {
		return "", err
	}

	return event, nil

}
