package commands

type Set struct {
	Name  string
	Key   string
	Value string
}

func NewSet(command Command) (Set, error) {
	var (
		argCountMin  = 2
		argCouintMax = 2
	)

	set := Set{
		Name: SetIdentifier,
	}

	err := validateArgCount(command[1:], argCountMin, argCouintMax)
	if err != nil {
		return set, err
	}

	set.Key = command[1]
	set.Value = command[2]

	return set, nil
}
