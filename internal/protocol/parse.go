package protocol

import (
	"bytes"
	"fmt"
	"strconv"
)

type UnknownTypeIdentifierError struct {
	Identifier byte
}

func (e *UnknownTypeIdentifierError) Error() string {
	return fmt.Sprintf("unknown type identifier: %c", e.Identifier)
}

type parser struct {
	data []byte
	pos  int
}

func NewParser(data []byte) parser {
	return parser{
		data: data,
		pos:  0,
	}
}

func (p *parser) peek(n int) []byte {
	return p.data[p.pos : p.pos+n-1]
}

func (p *parser) parseLength() (int, error) {
	currentPos := p.pos
	for !bytes.Equal(p.data[currentPos:currentPos+2], []byte("\r\n")) {
		currentPos += 1
	}

	length, err := strconv.Atoi(string(p.data[p.pos:currentPos]))
	if err != nil {
		return 0, err
	}

	p.pos = currentPos + 2
	return length, nil
}

func (p *parser) parseArray() ([]interface{}, error) {
	length, err := p.parseLength()
	if err != nil {
		return nil, err
		// TODO add details
	}

	arr := make([]interface{}, length)

	for i := 0; i < length; i++ {
		msg, err := p.ParseMessage()
		if err != nil {
			return nil, err
		}

		arr[i] = msg
	}

	return arr, nil
}

func (p *parser) parseBulkString() (string, error) {
	length, err := p.parseLength()
	if err != nil {
		return "", err
		// TODO add details
	}

	str := string(p.data[p.pos : p.pos+length])
	p.pos += length + 2
	return str, nil
}

func (p *parser) ParseMessage() (interface{}, error) {
	dataTypeidentifier := p.data[p.pos]
	p.pos += 1

	switch dataTypeidentifier {
	case ArrayIdentifier:
		arr, err := p.parseArray()
		return arr, err
	case BulkStringIdentifier:
		str, err := p.parseBulkString()
		return str, err
	default:
		return nil, &UnknownTypeIdentifierError{p.data[p.pos-1]}
	}
}
