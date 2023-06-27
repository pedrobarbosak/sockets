package msgs

type HistoryMessage struct {
	Messages []*Message `json:"messages"`
	Online   []string   `json:"online"`
}

func NewHistoryMessage(msgs []*Message, users []string) *Message {
	return &Message{
		Type: History,
		Data: &HistoryMessage{
			Messages: msgs,
			Online:   users,
		},
	}
}
