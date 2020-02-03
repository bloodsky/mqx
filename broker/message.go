package main

type Message struct {
	Text string
	Visibility bool
}

func NewMessage(Text string, Visibility bool) *Message {
	return &Message{
		Text:    Text,
		Visibility: Visibility,
	}
}
