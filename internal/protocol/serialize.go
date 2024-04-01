package protocol

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	NullValue = fmt.Sprintf("%c%s", NullIdentifier, terminator)

	ErrNotSerizlizableType = errors.New("type is not serializable")
)

func Serialize(data interface{}) ([]byte, error) {
	var result string
	var err error

	if data == nil {
		return []byte(NullValue), err
	}

	switch v := data.(type) {
	case int:
		value := strconv.FormatInt(int64(v), 10)
		result = fmt.Sprintf("%c%s%s", IntegerIdentifier, value, terminator)
	case float64, float32:
		result = fmt.Sprintf("%c%f%s", FloatIdentifier, v, terminator)
	case bool:
		value := 'f'
		if v {
			value = 't'
		}
		result = fmt.Sprintf("%c%c%s", BoolIdentifier, value, terminator)
	case string:
		result = fmt.Sprintf("%c%s%s", SimpleStringIdentifier, v, terminator)
	case []byte:
		result = fmt.Sprintf("%c%d%s%s%s", BulkStringIdentifier, len(v), terminator, v, terminator)
	case error:
		result = fmt.Sprintf("%c%s%s", ErrorIdentifier, v, terminator)
	default:
		err = ErrNotSerizlizableType
	}

	return []byte(result), err
}
