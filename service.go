package sockets

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/pedrobarbosak/errors"
)

type Channel struct {
	upgrader websocket.Upgrader
	log      *log.Logger

	mutex   sync.RWMutex
	members []*Conn

	onConnection func(data any) (msgToSend any)
	onDisconnect func(data any)
	onMessage    func(data any, msg []byte) (msgToSend any)
	onBeforeSend func(data any, msg any) (msgToSend any)
}

func (c *Channel) Notify(data any, except ...any) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	exceptions := map[any]struct{}{}
	for _, e := range except {
		exceptions[e] = struct{}{}
	}

	for _, conn := range c.members {
		if _, exists := exceptions[conn.data]; !exists {
			conn.msgsToSend <- data
		}
	}
}

func (c *Channel) AddConnection(w http.ResponseWriter, r *http.Request, data any) error {
	conn, err := c.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return errors.New(err)
	}

	cc := newConn(conn, data)
	c.addMember(cc)

	go c.handleConn(cc)

	return nil
}

func (c *Channel) handleConn(conn *Conn) {
	defer conn.Close()
	defer c.removeMember(conn)
	defer c.onDisconnect(conn.data)

	if msg := c.onConnection(conn.data); msg != nil {
		fmt.Println("MSG:", msg)
		if err := conn.WriteMessage(msg); err != nil {
			c.log.Println(err)
			return
		}
	}

	for {
		select {
		case msg := <-conn.msgsToSend:

			msg = c.onBeforeSend(conn.data, msg)
			if msg != nil {
				if err := conn.WriteMessage(msg); err != nil {
					c.log.Println(err)
					return
				}
			}

		case msg := <-conn.receivedMsgs:
			if msg.error != nil {
				closeErr, ok := msg.error.(*websocket.CloseError)
				if !ok {
					c.log.Println(errors.New("unidentified error:", msg.error))
					return
				}

				c.log.Println(closeErr.Error())
				return
			}

			if toSent := c.onMessage(conn.data, msg.data); toSent != nil {
				if err := conn.WriteMessage(toSent); err != nil {
					c.log.Println(err)
					return
				}
			}
		}
	}
}

func New(name string) Channel {
	logger := log.New(os.Stdout, fmt.Sprintf("<channel: %s> ", name), 0)

	return Channel{
		upgrader: websocket.Upgrader{},
		log:      logger,
		mutex:    sync.RWMutex{},
		members:  nil,
		onConnection: func(data any) (msgToSend any) {
			logger.Println("<onConnection> new connection with id:", data)
			return nil
		},
		onDisconnect: func(data any) {
			logger.Println("<onDisconnect> disconnection with id:", data)
		},
		onMessage: func(data any, msg []byte) (msgToSend any) {
			logger.Println("<onMessage> connection with id", data, "received msg:", string(msg))
			return nil
		},
		onBeforeSend: func(data any, msg any) (msgToSend any) {
			logger.Println("<onBeforeSend> sending", msg, "to connection with id:", data)
			return msg
		},
	}
}
