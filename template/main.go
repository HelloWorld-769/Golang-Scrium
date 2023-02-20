package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type secret struct {
	Name    string
	Secret  string
	Message string
}

func main() {
	fmt.Println("Template testing...")

	secret := secret{
		"John",
		"Keep it safe",
		"Dont Tell anyone",
	}
	filePath := "C:/Users/prajw/Desktop/go-workspace/src/template/temp.txt"

	t, err := template.New("temp.txt").ParseFiles(filePath)
	if err != nil {
		log.Fatal(err)

	}
	err1 := t.Execute(os.Stdout, secret)

	if err != nil {
		log.Fatal(err1)
	}

}
