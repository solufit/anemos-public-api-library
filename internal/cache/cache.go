package cache

import (
	"context"

	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/redis/go-redis/v9"
)

// Anemosイベントタイプ
type EventType string

// AnemosイベントID
type Id string

func createEventTypeCache(redisClient *redis.Client, weekly_data *mapset.Set[EventType], channel *chan error) {

	ctx := context.Background()

	// 一週間分のキャッシュデータを作成する
	_, err := redisClient.SAdd(ctx, "EventType", weekly_data).Result()

	if err != nil {
		*channel <- err
		return
	}

	// キャッシュデータ作成完了を通知する
	*channel <- nil

}

func createWeeklyCache(redisClient *redis.Client, key EventType, event mapset.Set[Id], channel *chan error) {

	ctx := context.Background()

	// 一週間分のキャッシュデータを作成する
	_, err := redisClient.SAdd(ctx, string(key), event).Result()

	if err != nil {
		*channel <- err
		return
	}

	// キャッシュデータ作成完了を通知する
	*channel <- nil

}

func createEventCache(redisClient *redis.Client, key Id, event string, channel *chan error) {

	ctx := context.Background()

	// 一週間分のキャッシュデータを作成する
	_, err := redisClient.Set(ctx, string(key), event, 0).Result()

	if err != nil {
		*channel <- err
		return
	}

	// キャッシュデータ作成完了を通知する
	*channel <- nil

}

// APIから取得した情報をもとに、Cacheを作成する
func CreateCache(redisClient *redis.Client, anemosData []interface{}) error {

	// Redisに保存するデータ
	var target_data = make(map[Id]string)

	// 一週間分のキャッシュデータインデックスを作成する
	var weekly_data = make(map[EventType]mapset.Set[Id])

	// Object Typeの一覧
	var object_types = mapset.NewSet[EventType]()

	// 一週間分のキャッシュデータを作成する
	for _, data := range anemosData {
		var object_id Id = data.(map[string]interface{})["info_objectId"].(Id)
		var object_type EventType = data.(map[string]interface{})["object_type"].(EventType)

		//objectidをキーにして、データを保存する
		target_data[object_id] = data.(string)

		//イベントタイプをキーにして、objectidを保存する
		weekly_data[object_type].Add(object_id)

		//object_typeを保存する
		object_types.Add(object_type)

	}

	channel := make(chan error)

	go createEventTypeCache(redisClient, &object_types, &channel)

	for key, data := range weekly_data {
		go createWeeklyCache(redisClient, key, data, &channel)
	}

	for key, data := range target_data {
		go createEventCache(redisClient, key, data, &channel)
	}

	var errorSlice []error

	for c := range channel {
		if c != nil {
			errorSlice = append(errorSlice, c)
		}
	}

	if len(errorSlice) > 0 {
		return fmt.Errorf("error: %v", errorSlice)
	}

	return nil
}
