package models

type MessageType string

const (
	Normal     MessageType = "normal"
	Connect    MessageType = "connect"
	Disconnect MessageType = "disconnect"
	History    MessageType = "history"
)
