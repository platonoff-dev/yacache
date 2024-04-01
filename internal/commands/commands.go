package commands

import "errors"

const (
	PingCommand = "PING"
	EchoCommand = "ECHO"
)

var (
	ErrNotCommand          = errors.New("not a command")
	ErrInvalidCommandType  = errors.New("invalid type")
	ErrUnknownCommand      = errors.New("unknown command")
	ErrUnexpectedArgument  = errors.New("unexpected argument")
	ErrWrongArgumentNumber = errors.New("wrong argument number")
)

type Command struct {
	Identifier string
	Args       []string
}

func NewCommand(message any) (*Command, error) {
	rawCommand, ok := message.([]string)
	if !ok {
		return nil, ErrNotCommand
	}

	if len(rawCommand) == 0 {
		return nil, ErrNotCommand
	}
	commandName := rawCommand[0]

	switch commandName {
	case PingCommand:
		if len(rawCommand) > 1 {
			return nil, ErrUnexpectedArgument
		}
		return &Command{Identifier: PingCommand}, nil

	case EchoCommand:
		if len(rawCommand) > 2 {
			return nil, ErrWrongArgumentNumber
		}
		arg := rawCommand[1]
		if !ok {
			return nil, ErrInvalidCommandType
		}

		return &Command{Identifier: EchoCommand, Args: []string{arg}}, nil

	default:
		return nil, ErrUnknownCommand
	}
}
