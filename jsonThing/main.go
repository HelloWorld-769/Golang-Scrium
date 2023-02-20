package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to golang server....")
	//getRequest()
	//postRequestJson()
	//postRequestForm()
	//EncodeStringtoJson()
	DecodeJson()

}
func getRequest() {
	const myurl = "http://localhost:3000/get"

	res, err := http.Get(myurl)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(content))
}

func postRequestJson() {
	const myurl = "http://localhost:3000/post"

	responseBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price":"200000"
		}
	`)

	res, err := http.Post(myurl, "application/json", responseBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(content))
}

func postRequestForm() {

	const myurl = "http://localhost:3000/postform"

	data := url.Values{
		"name":       {"John Doe"},
		"occupation": {"gardener"},
	}

	res, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(content))

}

func EncodeStringtoJson() {
	values := map[string]string{
		"name":       "John Doe",
		"occupation": "gardener",
	}

	jsonFormat, err := json.Marshal(values)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonFormat))
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"name":"John Doe",
		"occupation":"gardener"
	}
	
	`)
	mp := make(map[string]string)
	check := json.Valid(jsonDataFromWeb)

	if check {
		fmt.Println("Json valid")
		json.Unmarshal(jsonDataFromWeb, &mp)
		fmt.Println(mp)

	} else {
		fmt.Println("Json invalid")
	}
}
