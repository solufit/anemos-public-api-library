package cache

import (
	"context"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/redis/go-redis/v9"
)

// Anemosイベントタイプ
type eventType string

// AnemosイベントID
type id string

func createEventTypeCache(redisClient *redis.Client, weekly_data *mapset.Set[eventType], channel *chan error) {

	// 一週間分のキャッシュデータを作成する
	ctx := context.Background()
	_, err := redisClient.SAdd(ctx, "EventType", weekly_data).Result()

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
	var target_data = make(map[id]string)

	// 一週間分のキャッシュデータインデックスを作成する
	var weekly_data = make(map[eventType]mapset.Set[id])

	// Object Typeの一覧
	var object_types = mapset.NewSet[eventType]()

	// 一週間分のキャッシュデータを作成する
	for _, data := range anemosData {
		var object_id id = data.(map[string]interface{})["info_objectId"].(id)
		var object_type eventType = data.(map[string]interface{})["object_type"].(eventType)

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
	}

	return nil
}
