package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type hotdog int

func handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println(r.Method)
	w.Write([]byte("HelloWorld"))
}
func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
func main() {
	mux := httprouter.New()
	mux.GET("/", handler)
	mux.GET("/sas/:name", hello)
	log.Fatal(http.ListenAndServe(":8080", mux))

}
