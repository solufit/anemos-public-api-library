package cache_test

import (
	"testing"

	"sort"

	"github.com/solufit/anemos-public-api-library/internal/cache"
	"github.com/stretchr/testify/assert"
)

func TestGetCache(t *testing.T) {
	client := NewMockRedis(t)

	// Prepare test data
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
		map[string]interface{}{
			"info_objectId": cache.Id("3"),
			"object_type":   cache.EventType("type2"),
			"data":          "event3",
		},
	}

	err := cache.CreateCache(client, anemosData)
	assert.NoError(t, err)

	// Test GetCache function
	result, err := cache.GetCache(client, cache.EventType("type1"))
	assert.NoError(t, err)
	assert.Equal(t, []string{"event1"}, result)

	result, err = cache.GetCache(client, cache.EventType("type2"))
	sort.Strings(result)
	assert.NoError(t, err)
	assert.Equal(t, []string{"event2", "event3"}, result)

	// Test invalid EventType
	_, err = cache.GetCache(client, cache.EventType("invalid_type"))
	assert.Error(t, err)
}
