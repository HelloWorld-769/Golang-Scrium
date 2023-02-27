package cont

import (
	"fmt"
	mod "jwt-token/models"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
)

//upgrader function..
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

//gloabl Variables.....
var clients = make(map[*websocket.Conn]bool)
var connCount = 0
var broadcast = make(chan []byte)
var currentConnection = make(chan *websocket.Conn)

//Creating a  websocket
func CreatingConn(w http.ResponseWriter, r *http.Request) {

	//pass token as parms in url string..
	token := r.URL.Query()["token"]
	tokenString := ""
	fmt.Println("Header are :", r.Header)
	for _, v := range token {
		tokenString += v
	}

	claims := &mod.Claims{}

	parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !parsedToken.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Creating connection error", err)
		return
	}
	conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprint("Welcome ", claims.Username)))

	connCount++
	if connCount > 2 {
		fmt.Println("Connection limit exceeds 2")
		connCount--
		conn.Close()
		return
	}
	clients[conn] = true
	fmt.Println(clients)
	Communicating(conn)
}

// Reading the incoming message
func Communicating(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			connCount--
			delete(clients, conn)
			fmt.Println("Error from read function:", err)
			break
		}
		fmt.Println("Reading message:", string(msg))
		broadcast <- msg
		// fmt.Println("Message in broadcast", string(<-broadcast))
		currentConnection <- conn
	}
}

//writing to the user
func StartListner() {
	for {
		msg := <-broadcast
		conn := <-currentConnection
		for client := range clients {
			if client == conn {
				continue
			}
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				fmt.Println("Error writing the message ", err)
				client.Close()
				delete(clients, client)
				return
			}
		}
	}
}
