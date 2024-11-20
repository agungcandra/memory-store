package cache_test

import (
	"testing"

	"github.com/agungcandra/durian/cache"
	"github.com/stretchr/testify/assert"
)

func TestMemoryCache(t *testing.T) {
	m := cache.NewMemoryCache()

	assert.NotNil(t, m)

	t.Run("put", func(t *testing.T) {
		m.Put("key1", "boolean", "true", "string", "Hello world")
	})

	t.Run("get", func(t *testing.T) {
		result := m.Get("key1")
		assert.Equal(t, "boolean: true, string: Hello world", result)
	})

	t.Run("search", func(t *testing.T) {
		result := m.Search("boolean", "true")
		assert.Equal(t, "key1", result)
	})

	t.Run("keys", func(t *testing.T) {
		result := m.Key()
		assert.Equal(t, "key1", result)
	})

	t.Run("delete", func(t *testing.T) {
		m.Delete("key1")
		result := m.Get("key1")
		assert.Equal(t, "No entry found for key1", result)
	})
}
