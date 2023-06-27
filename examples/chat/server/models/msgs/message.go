package msgs

type Message struct {
	Type MessageType `json:"type"`
	Data any         `json:"data"`
}
