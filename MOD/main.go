package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	// Write() is used to write the body of the response, while WriteHeader() is used to set the status code of the response.
	w.WriteHeader(http.StatusOK) // called automatically if not called explicitly
	w.Write([]byte("<h1>Hello World</h1>"))

}
