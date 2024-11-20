package cache_test

import (
	"testing"

	"github.com/agungcandra/durian/cache"
	"github.com/stretchr/testify/assert"
)

func TestNewAttributes(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		attributes := cache.NewAttributes("boolean", "true")
		assert.Equal(t, &cache.Attributes{
			Key:      "boolean",
			DataType: cache.Boolean,
			Value:    "true",
		}, attributes)
	})

	t.Run("false", func(t *testing.T) {
		attributes := cache.NewAttributes("boolean", "false")
		assert.Equal(t, &cache.Attributes{
			Key:      "boolean",
			DataType: cache.Boolean,
			Value:    "false",
		}, attributes)
	})

	t.Run("int", func(t *testing.T) {
		attributes := cache.NewAttributes("int", "20000")
		assert.Equal(t, &cache.Attributes{
			Key:      "int",
			DataType: cache.Integer,
			Value:    "20000",
		}, attributes)
	})

	t.Run("float", func(t *testing.T) {
		attributes := cache.NewAttributes("float", "20000.20")
		assert.Equal(t, &cache.Attributes{
			Key:      "float",
			DataType: cache.Decimal,
			Value:    "20000.20",
		}, attributes)
	})

	t.Run("string", func(t *testing.T) {
		attributes := cache.NewAttributes("string", "Hello")
		assert.Equal(t, &cache.Attributes{
			Key:      "string",
			DataType: cache.String,
			Value:    "Hello",
		}, attributes)
	})
}
