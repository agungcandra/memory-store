package cache

import "strconv"

type Attributes struct {
	Key      string
	DataType DataType
	Value    string
}

type DataType string

const (
	Boolean DataType = "bool"
	Integer DataType = "int"
	Decimal DataType = "decimal"
	String  DataType = "string"
)

// NewAttributes initialize new attributes
func NewAttributes(key string, value string) *Attributes {
	datatype := String
	switch {
	case IsBoolean(value):
		datatype = Boolean
	case IsInt(value):
		datatype = Integer
	case IsFloat(value):
		datatype = Decimal
	}

	return &Attributes{
		Key:      key,
		DataType: datatype,
		Value:    value,
	}
}

func IsBoolean(value string) bool {
	return value == "false" || value == "true"
}

func IsInt(value string) bool {
	_, err := strconv.ParseInt(value, 10, 64)
	return err == nil
}

func IsFloat(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}
