package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//wesocket basic...
/*var upgrader = websocket.Upgrader{
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
}*/

//updgrades http connection to websocket connection
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//Simple chat application in golang
func main() {
	fmt.Println("Working with simple chat app using sockets..")
	clients := make(map[*websocket.Conn]bool)

	broadcast := make(chan []byte)

	connection := make(chan *websocket.Conn)

	connCount := 0

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		isOpen := true
		if err != nil {
			fmt.Println(err)
			return
		}

		connCount++

		if connCount > 2 {
			fmt.Println("Conncection exceeds 2")
			connCount--
			conn.Close()
			return
		}
		fmt.Println(connCount)
		clients[conn] = true
		fmt.Println(clients)

		for isOpen {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				isOpen = false
				connCount--
				delete(clients, conn)
				fmt.Println("Error from main routine", err)

				break
			}
			fmt.Println("Message sent:", string(msg))
			broadcast <- msg
			connection <- conn
		}
	})

	//Another go routine to recieve the message a wrte it to client
	go func() {

		for {
			msg := <-broadcast
			conn := <-connection
			for client := range clients {
				if client == conn {
					continue
				}
				err := client.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					fmt.Println("Error from routine:", err)
					client.Close()
					delete(clients, client)
				}
			}
		}

	}()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
