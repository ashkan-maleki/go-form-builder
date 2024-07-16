package domain

type Message interface {
	Name() string
}

type Command interface {
	Message
}

type Event interface {
	Message
}
