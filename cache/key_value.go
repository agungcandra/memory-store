package cache

import (
	"fmt"
	"strings"
)

type KeyValue struct {
	Attributes map[string]*Attributes
}

func NewKeyValue(attributes ...string) (*KeyValue, error) {
	if len(attributes)%2 != 0 {
		return nil, fmt.Errorf("invalid attributes length")
	}
	attributeMap := make(map[string]*Attributes)
	for i := 0; i < len(attributes)-1; i += 2 {
		attributeMap[attributes[i]] = NewAttributes(attributes[i], attributes[i+1])
	}

	return &KeyValue{
		Attributes: attributeMap,
	}, nil
}

func (k *KeyValue) String() string {
	var sb strings.Builder
	for k, attribute := range k.Attributes {
		sb.WriteString(k + ": " + fmt.Sprint(attribute.Value) + ", ")
	}

	str := sb.String()
	if len(str) > 0 {
		return str[:len(str)-2]
	}
	return ""
}
