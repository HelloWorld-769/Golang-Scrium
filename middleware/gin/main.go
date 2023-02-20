package main

import (
	"fmt"
	"log"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing the handler")
	w.Write([]byte("OK"))
}

func middlewareHttp(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Running before handler")
		w.Write([]byte("Hijacking request "))
		fmt.Println(r.URL.Path)
		originalHandler.ServeHTTP(w, r)
		fmt.Println("Running after the handler")
	})
}
func main() {
	mux := http.NewServeMux()
	myHandler := http.HandlerFunc(handler)
	mux.Handle("/", middlewareHttp(myHandler))

	log.Print("Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", mux))

}
