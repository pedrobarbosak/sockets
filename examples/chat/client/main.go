package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"sockets/examples/chat/client/models"

	"github.com/gorilla/websocket"
	"github.com/pedrobarbosak/errors"
)

func inputHandler(conn *websocket.Conn) {
	reader := bufio.NewReader(os.Stdin)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Println(errors.New("failed to read:", err))
			return
		}

		if err = conn.WriteMessage(websocket.TextMessage, []byte(strings.TrimSpace(msg))); err != nil {
			log.Println(errors.New("failed to write to websocket:", err))
			return
		}
	}
}

func receiveHandler(conn *websocket.Conn) {
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			if err == websocket.ErrCloseSent {
				return
			}

			log.Println(errors.New("error receiving msg:", err))
			continue
		}

		var msg models.Message
		if err = json.Unmarshal(data, &msg); err != nil {
			log.Println(errors.New("failed to unmarshal msg:", err))
		}

		fmt.Println(msg.String())
	}
}

func main() {
	// connect to server
	socketUrl := "ws://localhost:8080" + "/msg"
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Panicln(errors.New(err))
	}

	defer conn.Close()

	// launch handlers
	go inputHandler(conn)
	go receiveHandler(conn)

	// Terminate gracefully
	interrupt := make(chan os.Signal) // Channel to listen for interrupt signal to terminate gracefully

	signal.Notify(interrupt, os.Interrupt) // Notify the interrupt channel for SIGINT

	<-interrupt

	err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println(errors.New("error during closing websocket:", err))
		return
	}
}
