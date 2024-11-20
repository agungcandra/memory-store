package cache_test

import (
	"testing"

	"github.com/agungcandra/durian/cache"
	"github.com/stretchr/testify/assert"
)

func TestNewKeyValue(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		keyValue, err := cache.NewKeyValue(
			"boolean", "true",
			"string", "Hello world",
		)

		assert.Nil(t, err)
		assert.Equal(t, &cache.KeyValue{
			Attributes: map[string]*cache.Attributes{
				"boolean": {
					Key:      "boolean",
					DataType: cache.Boolean,
					Value:    "true",
				},
				"string": {
					Key:      "string",
					DataType: cache.String,
					Value:    "Hello world",
				},
			},
		}, keyValue)
	})

	t.Run("failed", func(t *testing.T) {
		keyValue, err := cache.NewKeyValue(
			"boolean", "true",
			"string",
		)

		assert.Equal(t, "invalid attributes length", err.Error())
		assert.Nil(t, keyValue)
	})
}

func TestKeyValue_String(t *testing.T) {
	keyValue, err := cache.NewKeyValue(
		"boolean", "true",
		"string", "Hello world",
	)
	assert.Nil(t, err)
	assert.Equal(t, "boolean: true, string: Hello world", keyValue.String())

}
