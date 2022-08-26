package domain

type Event[Message any] struct {
	Message *Message
}
