package cache_test

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/redis/go-redis/v9"
	"github.com/solufit/anemos-public-api-library/internal/cache"
	"github.com/stretchr/testify/assert"
)

func NewMockRedis(t *testing.T) *redis.Client {
	t.Helper()

	// redisサーバを作る
	s, err := miniredis.Run()
	if err != nil {
		t.Fatalf("unexpected error while creating test redis server '%#v'", err)
	}
	// *redis.Clientを用意
	client := redis.NewClient(&redis.Options{
		Addr:     s.Addr(),
		Password: "",
		DB:       0,
	})
	return client
}

func TestCreateCache(t *testing.T) {
	client := NewMockRedis(t)

	anemosData := []interface{}{
		map[string]interface{}{
			"info_objectId": cache.Id("1"),
			"object_type":   cache.EventType("type1"),
			"data":          "event1",
		},
		map[string]interface{}{
			"info_objectId": cache.Id("2"),
			"object_type":   cache.EventType("type2"),
			"data":          "event2",
		},
	}

	err := cache.CreateCache(client, anemosData)
	assert.NoError(t, err)

	// Check if the event types are cached
	eventTypes, err := client.SMembers(context.Background(), "EventType").Result()
	assert.NoError(t, err)
	expectedEventTypes := mapset.NewSet("type1", "type2")
	actualEventTypes := mapset.NewSet(eventTypes...)
	assert.True(t, expectedEventTypes.Equal(actualEventTypes))

	// Check if the weekly data is cached
	for _, data := range anemosData {
		objectType := data.(map[string]interface{})["object_type"].(string)
		objectId := data.(map[string]interface{})["info_objectId"].(string)
		members, err := client.SMembers(context.Background(), objectType).Result()
		assert.NoError(t, err)
		assert.Contains(t, members, objectId)
	}

	// Check if the individual events are cached
	for _, data := range anemosData {
		objectId := data.(map[string]interface{})["info_objectId"].(string)
		eventData, err := client.Get(context.Background(), objectId).Result()
		assert.NoError(t, err)
		assert.Equal(t, data.(map[string]interface{})["data"].(string), eventData)
	}
}
