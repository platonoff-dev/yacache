package protocol

import (
	"testing"

	"gotest.tools/assert"
)

func TestParser_ParseMessage(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected any
		err      error
	}{
		{
			name:     "BasicString",
			data:     []byte("$4\r\nPING\r\n"),
			expected: "PING",
			err:      nil,
		},
		{
			name:     "EmptyString",
			data:     []byte("$0\r\n\r\n"),
			expected: "",
			err:      nil,
		},
		{
			name:     "BasicArray",
			data:     []byte("*2\r\n$4\r\nPING\r\n$4\r\nPONG\r\n"),
			expected: []any{"PING", "PONG"},
			err:      nil,
		},
		{
			name:     "EmptyArray",
			data:     []byte("*0\r\n"),
			expected: []any{},
			err:      nil,
		},
		{
			name:     "UnexpectedType",
			data:     []byte("_some\r\n"),
			expected: nil,
			err:      &UnknownTypeIdentifierError{Identifier: '_'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.data)

			assert.DeepEqual(t, tt.expected, result)
			assert.DeepEqual(t, tt.err, err)
		})
	}
}
