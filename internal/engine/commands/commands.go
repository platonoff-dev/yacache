package commands

import "errors"

const (
	EchoIdentifier = "ECHO"
	PingIdentifier = "PING"
	GetIdentifier  = "GET"
	SetIdentifier  = "SET"
)

var (
	ErrWrongNumberOfArguments = errors.New("wrong number of arguments")
)

type Command = []string

func validateArgCount(args []string, min, max int) error {
	if len(args) < min || len(args) > max {
		return ErrWrongNumberOfArguments
	}

	return nil
}

func getArg(command Command, n int) (string, bool) {
	if len(command)-1 < n {
		return "", false
	}

	return command[n], true
}
