package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Product struct {
	Id       string
	Name     string
	Price    float64
	Quantity int
	Status   bool
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", welcomeText)
	mux.HandleFunc("/html", Index)
	log.Print("Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
func welcomeText(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
func Index(res http.ResponseWriter, req *http.Request) {
	data := map[string]interface{}{
		"fullname": "john doe",
		"product": Product{
			Id:       "1",
			Name:     "iPhone X",
			Price:    1.99,
			Quantity: 5,
			Status:   true,
		},
	}

	//func help to add the functionality to the template

	tmp, _ := template.New("index.html").Funcs(template.FuncMap{
		"upper": func(name string) string {
			return strings.ToUpper(name)
		},
		"displayStatus": func(status bool) string {
			if status {
				return "Show"
			}
			return "Hide"
		},
		"total": func(prod Product) float64 {
			return prod.Price * float64(prod.Quantity)
		},
		"checkStock": func(prod Product) string {
			if prod.Quantity == 0 {
				return "Out of stock"
			} else {
				return "In stock"
			}
		},
	}).ParseFiles("index.html")
	tmp.Execute(res, data)
}
