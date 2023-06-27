package sockets

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gorilla/websocket"
	"github.com/pedrobarbosak/errors"
)

type receivedMsg struct {
	mType int
	data  []byte
	error error
}

type Conn struct {
	id   string
	conn *websocket.Conn
	data any

	receivedMsgs chan *receivedMsg
	msgsToSend   chan any
}

func (c *Conn) Close() error {
	return c.conn.Close()
}

func (c *Conn) WriteMessage(data any) error {
	if err := c.conn.WriteJSON(data); err != nil {
		return errors.New("failed to send message:", err)
	}

	return nil
}

func (c *Conn) listener() {
	for {
		mType, message, err := c.conn.ReadMessage()
		c.receivedMsgs <- &receivedMsg{
			mType: mType,
			data:  message,
			error: err,
		}
	}
}

func newConn(conn *websocket.Conn, data any) *Conn {
	c := &Conn{
		id:           gofakeit.UUID(),
		conn:         conn,
		data:         data,
		receivedMsgs: make(chan *receivedMsg),
		msgsToSend:   make(chan any),
	}

	go c.listener()

	return c
}
