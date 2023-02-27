package main

import (
	"fmt"
	cont "jwt-token/controllers"
	"log"
	"net/http"
)

// // creating sscret key used for signing the token

func main() {

	fmt.Println("Running sever on port 8080")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	mux.HandleFunc("/login", cont.LoginHandler)
	mux.HandleFunc("/welcome", cont.MiddleWare(http.HandlerFunc(cont.WelcomeHandler)))
	mux.HandleFunc("/refresh", cont.RefreshHandler)
	mux.HandleFunc("/logout", cont.LogoutHandler)
	mux.HandleFunc("/register", cont.RegisterHandler)

	//websocket request..
	mux.HandleFunc("/ws", cont.CreatingConn)

	go cont.StartListner()

	log.Fatal(http.ListenAndServe(":8080", mux))
}
