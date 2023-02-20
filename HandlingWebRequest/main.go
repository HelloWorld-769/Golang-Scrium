package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	fmt.Println("Handling web requests in go..")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type is: %T\n", res)
	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(content))

}
