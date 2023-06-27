package msgs

type DisconnectMessage struct {
	Timestamp int64  `json:"timestamp"`
	Username  string `json:"username"`
}

func NewDisconnectMessage(timestamp int64, username string) *Message {
	return &Message{
		Type: Disconnect,
		Data: &DisconnectMessage{
			Timestamp: timestamp,
			Username:  username,
		},
	}
}
