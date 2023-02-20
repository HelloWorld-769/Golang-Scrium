package main

import (
	"fmt"
	"log"
	"net/http"
)

// middleware 1
func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received %s request for %s\n", r.Method, r.URL.Path)
		fmt.Println("before middleware 1")
		next.ServeHTTP(w, r)
		fmt.Println("After middleware 1")
	})
}

// middleware 2
func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before middlware 2")
		next.ServeHTTP(w, r)
		fmt.Println("After middleware 2")
	})
}

//final handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to home page..!")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.Handle("/chainMid", logRequest(authenticate(http.HandlerFunc(homeHandler))))
	log.Fatal(http.ListenAndServe(":4000", mux))
}
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Go Middleware"))
}
