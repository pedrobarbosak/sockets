package msgs

type ConnectMessage struct {
	Timestamp int64  `json:"timestamp"`
	Username  string `json:"username"`
}

func NewConnectMessage(timestamp int64, username string) *Message {
	return &Message{
		Type: Connect,
		Data: &ConnectMessage{
			Timestamp: timestamp,
			Username:  username,
		},
	}
}
