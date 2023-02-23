package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs"

func main() {
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)   //https
	fmt.Println(result.Host)     //lco.dev:3000
	fmt.Println(result.Path)     //learn
	fmt.Println(result.RawQuery) // "courses":"reactjs"
	fmt.Println(result.Port())   //3000

	qParams := result.Query()
	fmt.Println(qParams)
	// fmt.Println(qParams["coursename"]) //[reactjs]

	//Create a url form its different parts...

	// partsOfUrl := &url.URL{
	// 	Scheme:   "http",
	// 	Host:     "loc.dev",
	// 	Path:     "/learn",
	// 	RawQuery: "coursename=reactjs&paymneid=5fd5fd",
	// }

	// fmt.Println(partsOfUrl.String())

	// u := &url.URL{
	// 	Scheme: "https",
	// 	User:   url.UserPassword("user", "password"),
	// 	Host:   "example.com",
	// 	Path:   "foo/bar",
	// }
	// fmt.Println(u.Redacted())
	// use, a := url.UserPassword("me", "newerPassword").Password()
	// fmt.Println(use, a)

}
