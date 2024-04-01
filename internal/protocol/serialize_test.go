package protocol

import (
	"context"
	"errors"
	"testing"

	"gotest.tools/assert"
)

func TestSerialize(t *testing.T) {
	tests := []struct {
		name     string
		data     interface{}
		expected []byte
		err      error
	}{
		{
			name:     "SerializeInt",
			data:     42,
			expected: []byte(":42\r\n"),
			err:      nil,
		},
		{
			name:     "SerializeFloat",
			data:     3.14,
			expected: []byte(",3.140000\r\n"),
			err:      nil,
		},
		{
			name:     "SerializeBoolTrue",
			data:     true,
			expected: []byte("#t\r\n"),
			err:      nil,
		},
		{
			name:     "SerializeBoolFalse",
			data:     false,
			expected: []byte("#f\r\n"),
			err:      nil,
		},
		{
			name:     "SerializeString",
			data:     "hello",
			expected: []byte("+hello\r\n"),
			err:      nil,
		},
		{
			name:     "SerializeByteSlice",
			data:     []byte("world"),
			expected: []byte("$5\r\nworld\r\n"),
			err:      nil,
		},
		{
			name:     "SerializeError",
			data:     ErrNotSerizlizableType,
			expected: []byte("-type is not serializable\r\n"),
			err:      nil,
		},
		{
			name:     "SerializeNil",
			data:     nil,
			expected: []byte("_\r\n"),
			err:      nil,
		},
		{
			name:     "SerializeUnknown",
			data:     context.TODO(),
			expected: []byte{},
			err:      ErrNotSerizlizableType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Serialize(tt.data)

			assert.DeepEqual(t, tt.expected, result)
			assert.Equal(t, errors.Is(tt.err, err), true)
		})
	}
}
