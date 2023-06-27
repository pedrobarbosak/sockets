package models

import (
	"sync"

	"sockets/examples/chat/server/models/msgs"
)

type MessageHistory struct {
	messages []*msgs.Message
	mutex    sync.RWMutex
}

func (h *MessageHistory) AddMessage(msg *msgs.Message) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	h.messages = append(h.messages, msg)
}

func (h *MessageHistory) GetMessages() []*msgs.Message {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	return h.messages
}

func NewMessageHistory(msgs ...*msgs.Message) *MessageHistory {
	return &MessageHistory{
		messages: msgs,
		mutex:    sync.RWMutex{},
	}
}
