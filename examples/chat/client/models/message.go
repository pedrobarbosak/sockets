package models

import (
	"fmt"
	"time"
)

type Message struct {
	Type MessageType `json:"type"`
	Data MessageBody `json:"data"`
}

func (m Message) normalMsg() string {
	return fmt.Sprintf("%s (%s): %s", m.Data.Username, time.Unix(m.Data.Timestamp, 64).Format("02-01-2006 15:04"), m.Data.Content)
}

func (m Message) String() string {
	switch m.Type {

	case Connect:
		return fmt.Sprintf("[+] %s is now online ...", m.Data.Username)

	case Disconnect:
		return fmt.Sprintf("[-] %s disconnected ...", m.Data.Username)

	case History:
		msgs := ""
		for _, msg := range m.Data.Messages {
			msgs += fmt.Sprintf("\t%s\n", msg.String())
		}

		return fmt.Sprintf("[=] Users online: %v\n", m.Data.Online) +
			fmt.Sprintf("[=] Previous messages:\n%s", msgs)

	case Normal:
		return fmt.Sprintf("[!] %s (%s): %s", m.Data.Username, time.Unix(m.Data.Timestamp, 64).Format("02-01-2006 15:04"), m.Data.Content)

	default:
		return ""
	}
}

type MessageBody struct {
	ID        string     `json:"id"`
	Timestamp int64      `json:"timestamp"`
	Username  string     `json:"username"`
	Content   string     `json:"content"`
	Messages  []*Message `json:"messages"`
	Online    []string   `json:"online"`
}
