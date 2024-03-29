package protocol

type Query interface {
	Command() string
	Args() []string
}
