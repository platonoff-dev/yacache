package engine

import (
	"errors"
	"strings"

	"github.com/platonoff-dev/yacache/internal/engine/commands"
)

var (
	ErrEmptyCommand        = errors.New("empty command")
	ErrUnsuppportedCommand = errors.New("unsupported command")
)

type Engine struct {
}

func New() Engine {
	return Engine{}
}

func (e *Engine) ExecuteCommand(command commands.Command) (any, error) {
	if len(command) < 1 {
		return nil, ErrUnsuppportedCommand
	}

	switch strings.ToUpper(command[0]) {
	case commands.PingIdentifier:
		return e.ping(command)
	case commands.EchoIdentifier:
		return e.echo(command)
	default:
		return nil, ErrUnsuppportedCommand
	}
}

func (e *Engine) ping(command commands.Command) (string, error) {
	cmd, err := commands.NewPing(command)
	if err != nil {
		return "", err
	}

	if cmd.Message != "" {
		return cmd.Message, nil
	}

	return "PONG", nil
}

func (e *Engine) echo(command commands.Command) (string, error) {
	cmd, err := commands.NewEcho(command)
	if err != nil {
		return "", err
	}

	return cmd.Message, nil
}
