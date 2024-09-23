package cache

import (
	"github.com/redis/go-redis/v9"
)

// Anemosイベントタイプ
type eventType string

// AnemosイベントID
type id string

// APIから取得した情報をもとに、Cacheを作成する
func CreateCache(redisClient *redis.Client, anemosData []interface{}) error {

	// Redisに保存するデータ
	var target_data = make(map[id]string)

	// 一週間分のキャッシュデータインデックスを作成する
	var weekly_data = make(map[eventType][]id)

	// Object Typeの一覧
	var object_types []eventType

	// 一週間分のキャッシュデータを作成する
	for _, data := range anemosData {
		var object_id id = data.(map[string]interface{})["info_objectId"].(id)
		var object_type eventType = data.(map[string]interface{})["object_type"].(eventType)
		//objectidをキーにして、データを保存する
		target_data[object_id] = data.(string)

		//イベントタイプをキーにして、objectidを保存する
		weekly_data[object_type] = append(weekly_data[object_type], object_id)

	}

	return nil
}
