package engine

import (
	"errors"
	"fmt"

	"github.com/platonoff-dev/yacache/internal/commands"
)

var (
	ErrUnsuppportedCommand = errors.New("unsupported command")

	SupportedCommands = []string{
		commands.EchoCommand,
		commands.PingCommand,
	}
)

func ping() string {
	return "+PONG\r\n"
}

func echo(text string) string {
	return fmt.Sprintf("$%d\r\n%s\r\n", len(text), text)
}

func ExecuteCommand(command *commands.Command) (string, error) {
	switch command.Identifier {
	case commands.PingCommand:
		return ping(), nil
	case commands.EchoCommand:
		return echo(command.Args[0]), nil
	default:
		return "", ErrUnsuppportedCommand
	}
}
