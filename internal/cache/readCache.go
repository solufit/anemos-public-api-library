package cache

import (
	"context"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/redis/go-redis/v9"
)

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

// Get Cached data from Redis Cache database
// This function return json as a string
func GetCache(redisClient *redis.Client, key EventType) ([]string, error) {

	//イベントタイプが正しいか確認する
	eventTypes, err := getEventTypeCache(redisClient)

	if err != nil {
		return nil, fmt.Errorf("failed to get EventType cache: %v", err)
	}

	if !eventTypes.Contains(key) {
		return nil, fmt.Errorf("invalid EventType: %v", key)
	}

	//所属するイベントIDを取得する
	eventIds, err := getWeeklyCache(redisClient, key)

	if err != nil {
		return nil, fmt.Errorf("failed to get weekly cache: %v", err)
	}

	//イベントIDに対応するイベントを取得する

	channelErr := make(chan error)
	channelResult := make(chan string)

	for _, eventId := range eventIds.ToSlice() {
		go func(eventId Id) {
			defer close(channelResult)
			defer close(channelErr)

			event, err := getEventCache(redisClient, eventId)
			if err != nil {
				channelErr <- err
				return
			}
			channelResult <- event
		}(eventId)
	}

	// データをうけとる
	var errResult []error
	var event []string

	// イベントデータを取得
	for result := range channelResult {
		event = append(event, result)
	}
	for err := range channelErr {
		errResult = append(errResult, err)
	}

	// エラー処理
	for err := range channelErr {
		if err != nil {
			errResult = append(errResult, err)
		}

	}

	if len(errResult) > 0 {

		return nil, fmt.Errorf("failed to get event cache: %v", errResult)
	}

	return event, nil

}
