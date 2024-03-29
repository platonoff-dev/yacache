package protocol

import (
	"testing"

	"gotest.tools/assert"
)

func TestParser_ParseMessage(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected interface{}
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
			expected: []interface{}{"PING", "PONG"},
			err:      nil,
		},
		{
			name:     "EmptyArray",
			data:     []byte("*0\r\n"),
			expected: []interface{}{},
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
			p := NewParser(tt.data)

			result, err := p.ParseMessage()

			assert.DeepEqual(t, tt.expected, result)
			assert.DeepEqual(t, tt.err, err)
		})
	}
}
