package commands

type Get struct {
	Name string
	Key  string
}

func NewGet(command Command) (Get, error) {
	var (
		argCountMin = 1
		argCountMax = 1
	)

	get := Get{
		Name: GetIdentifier,
	}

	err := validateArgCount(command[1:], argCountMin, argCountMax)
	if err != nil {
		return get, err
	}

	get.Key = command[1]
	return get, nil
}
