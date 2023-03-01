package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const apiURL = "https://newsdata.io/api/1/news?apikey=pub_18051652e0aec2dcf1759b169d73ad51d01d8&language=en"

// using go routines
func ApiHandler1(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query().Get("z")

	newApi := apiURL + "&" + params

	respChan := make(chan []byte)
	errChan := make(chan error)

	go func() {
		res, err := http.Get(newApi)
		if err != nil {
			errChan <- err
			return
		}

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			errChan <- err
			return
		}

		respChan <- body
	}()

	select {
	case resp := <-respChan:
		w.Write(resp)
	case err := <-errChan:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	case <-time.After(2 * time.Second):
		http.Error(w, "API call timed out", http.StatusGatewayTimeout)
	}
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query().Get("z")

	newApi := apiURL + "&" + params

	res, err := http.Get(newApi)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(body))

}
func main() {
	port := ":8080"
	fmt.Println("Listening on port:", port)
	mux := http.NewServeMux()
	mux.HandleFunc("/api", ApiHandler1)
	log.Fatal(http.ListenAndServe(port, mux))
}
