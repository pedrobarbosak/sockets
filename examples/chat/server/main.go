package main

import (
	"log"
	"net/http"
	"time"

	"sockets"
	"sockets/examples/chat/server/models"
	"sockets/examples/chat/server/models/msgs"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	online := models.NewOnlineUsers()
	history := models.NewMessageHistory(
		msgs.NewConnectMessage(time.Now().Unix(), "@hugo"),
		msgs.NewConnectMessage(time.Now().Unix(), "@fernando"),
		msgs.NewNormalMessage(gofakeit.UUID(), time.Now().Unix(), "@hugo", "Heloooo"),
		msgs.NewNormalMessage(gofakeit.UUID(), time.Now().Unix(), "@fernando", "Hi there!"),
	)

	channel := sockets.New("messages")

	// Called every time that a client connects
	channel.OnConnection(func(data any) any {
		id := data.(string)
		online.Add(id)

		msg := msgs.NewConnectMessage(time.Now().Unix(), id)
		history.AddMessage(msg)

		channel.Notify(msg, id)

		return msgs.NewHistoryMessage(history.GetMessages(), online.GetAll())
	})

	// Called every time that a client disconnects
	channel.OnDisconnect(func(data any) {
		id := data.(string)

		online.Delete(id)

		msg := msgs.NewDisconnectMessage(time.Now().Unix(), id)
		history.AddMessage(msg)

		channel.Notify(msg, id)
	})

	// Called every time that a client sends a message
	channel.OnMessage(func(data any, bytes []byte) any {
		id := data.(string)

		msg := msgs.NewNormalMessage(gofakeit.UUID(), time.Now().Unix(), id, string(bytes))
		history.AddMessage(msg)

		channel.Notify(msg, id)
		return nil // todo: confirm msg maybe?
	})

	http.HandleFunc("/msg", func(w http.ResponseWriter, r *http.Request) {
		log.Println(channel.AddConnection(w, r, gofakeit.FirstName()))
	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
