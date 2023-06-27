package msgs

type NormalMessage struct {
	ID        string `json:"id"`
	Timestamp int64  `json:"timestamp"`
	Username  string `json:"username"`
	Content   string `json:"content"`
}

func NewNormalMessage(id string, timestamp int64, username string, content string) *Message {
	return &Message{
		Type: Normal,
		Data: &NormalMessage{
			ID:        id,
			Timestamp: timestamp,
			Username:  username,
			Content:   content,
		},
	}
}
