package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Received message:", string(message))

		err = conn.WriteMessage(websocket.TextMessage, []byte("message"))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
func main() {
	http.HandleFunc("/", echo)
	http.ListenAndServe(":8081", nil)
}
