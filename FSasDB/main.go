package main

import (
	"fmt"
	cont "fs-db/controllers"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Server Running on port 8000")
	mux := http.NewServeMux()
	mux.HandleFunc("/register", cont.RegisterHandler)
	mux.HandleFunc("/getUsers", cont.GetUserHandler)
	mux.HandleFunc("/update", cont.UpdateHandler)
	mux.HandleFunc("/delete", cont.DeleteHandler)

	log.Fatal(http.ListenAndServe(":8000", mux))
}
