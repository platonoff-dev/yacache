package commands

const (
	argCountMax = 1
	argCountMin = 0
)

type Ping struct {
	Name    string
	Message string
}

func NewPing(command Command) (Ping, error) {
	ping := Ping{
		Name: PingIdentifier,
	}

	err := validateArgCount(command[1:], argCountMin, argCountMax)
	if err != nil {
		return ping, err
	}

	if message, ok := getArg(command, 1); ok {
		ping.Message = message
	}

	return ping, nil
}
