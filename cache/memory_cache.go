package cache

import (
	"fmt"
	"sort"
	"strings"
)

// Memorycache store key value pair
type MemoryCache struct {
	KeyVal map[string]*KeyValue
}

// NewMemoryCache initialize new memory cache
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		KeyVal: map[string]*KeyValue{},
	}
}

func (m *MemoryCache) Get(key string) string {
	attributes, exists := m.KeyVal[key]
	if !exists {
		return fmt.Sprintf("No entry found for %s", key)
	}
	return attributes.String()
}

func (m *MemoryCache) Put(key string, attribtues ...string) error {
	keyValue, err := NewKeyValue(attribtues...)
	if err != nil {
		return fmt.Errorf("invalid attributes")
	}

	m.KeyVal[key] = keyValue

	return nil
}

func (m *MemoryCache) Delete(key string) {
	delete(m.KeyVal, key)
}

func (m *MemoryCache) Search(attributeKey, attributeValue string) string {
	var result []string
	for key, value := range m.KeyVal {
		for attrKey, attrValue := range value.Attributes {
			if attrKey == attributeKey && fmt.Sprint(attrValue.Value) == attributeValue {
				result = append(result, key)
			}
		}
	}

	sort.Strings(result)
	return strings.Join(result, ", ")
}

func (m *MemoryCache) Key() string {
	var keys []string
	for key := range m.KeyVal {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	return strings.Join(keys, ", ")
}
