package cache

import (
	"context"
	"time"

	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/redis/go-redis/v9"
)

// Anemosイベントタイプ
type EventType string

// AnemosイベントID
type Id string

func createEventTypeCache(redisClient *redis.Client, weekly_data mapset.Set[EventType], channel chan error) {

	ctx := context.Background()

	var weeklyDataSlice []string
	for _, data := range weekly_data.ToSlice() {
		weeklyDataSlice = append(weeklyDataSlice, string(string(data)))
	}

	// 一週間分のキャッシュデータを作成する
	_, err := redisClient.SAdd(ctx, "EventType", weeklyDataSlice).Result()

	if err != nil {
		channel <- err
		return
	}

	// キャッシュデータ作成完了を通知する
	channel <- nil

}

func createWeeklyCache(redisClient *redis.Client, key EventType, event mapset.Set[Id], channel chan error) {

	ctx := context.Background()
	var eventSlice []string
	for _, id := range event.ToSlice() {
		eventSlice = append(eventSlice, string(string(id)))
	}

	// 一週間分のキャッシュデータを作成する
	_, err := redisClient.SAdd(ctx, string(key), eventSlice).Result()

	if err != nil {
		channel <- err
		return
	}

	// キャッシュデータ作成完了を通知する
	channel <- nil

}

func createEventCache(redisClient *redis.Client, key Id, event string, channel chan error) {

	ctx := context.Background()

	// 一週間分のキャッシュデータを作成する
	day := 7
	expire := time.Duration(day*24) * time.Hour
	_, err := redisClient.Set(ctx, string(key), interface{}(event), expire).Result()

	if err != nil {
		channel <- err
		return
	}

	// キャッシュデータ作成完了を通知する
	channel <- nil

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
		dataMap := data.(map[string]interface{})
		var object_id Id = Id(dataMap["info_objectId"].(string))
		var object_type EventType = EventType(dataMap["object_type"].(string))

		// objectidをキーにして、データを保存する
		// convert data to string
		stringData := fmt.Sprintf("%v", data)
		target_data[object_id] = stringData

		// イベントタイプをキーにして、objectidを保存する
		if _, exists := weekly_data[object_type]; !exists {
			weekly_data[object_type] = mapset.NewSet[Id]()
		}
		weekly_data[object_type].Add(object_id)

		// object_typeを保存する
		object_types.Add(object_type)
	}

	channel := make(chan error)

	go createEventTypeCache(redisClient, object_types, channel)

	for key, data := range weekly_data {
		go createWeeklyCache(redisClient, key, data, channel)
	}

	for key, data := range target_data {
		go createEventCache(redisClient, key, data, channel)
	}

	var errorSlice []error

	for i := 0; i < len(weekly_data)+len(target_data)+1; i++ {
		err := <-channel
		if err != nil {
			errorSlice = append(errorSlice, err)
		}
	}

	if len(errorSlice) > 0 {
		return fmt.Errorf("error: %v", errorSlice)
	}

	return nil
}
