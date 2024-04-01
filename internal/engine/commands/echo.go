package commands

type Echo struct {
	Name    string
	Message string
}

func NewEcho(command Command) (Echo, error) {
	echo := Echo{
		Name: EchoIdentifier,
	}

	err := validateArgCount(command[1:], 1, 1)
	if err != nil {
		return echo, err
	}

	if message, ok := getArg(command, 1); ok {
		echo.Message = message
	}

	return echo, nil
}
