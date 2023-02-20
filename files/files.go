package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Ninja struct {
	Name   string
	Weapon string
	Level  int
}

func createFile() {
	content := "hello World!"
	file, err := os.Create("text.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	length, err := file.WriteString(content)
	fmt.Println(length)
}
func fun() {
	sFrom := `{
		"name":"John",
		"weapon":"star2",
		"level":2
	}`

	check := json.Valid([]byte(sFrom))

	// a, _ := json.Marshal(sFrom)

	// fmt.Println("a", a)

	fmt.Println(check)
	var s Ninja
	err := json.Unmarshal([]byte(sFrom), &s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
	// file, err := os.OpenFile("text.txt", os.O_APPEND|os.O_CREATE, 0600)
	// defer file.Close()
	// len, _ := fmt.Fprint(file, s)
	// fmt.Print(len)
}
func appendToFile(fileName string) {
	fmt.Println("Updating to file...")
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, 0444)

	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// _, err := os.Lstat("text.txt")
	// fmt.Println(fs.Mode().Perm())

	file.WriteString("File appended successfully")

	readFile("text.txt")
}

func readFile(filename string) {

	fmt.Println("reading from file...")
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text extracted from file:", string(databyte))

}
func readUsingOpen(filename string) {

	fmt.Println("Reading from file using Open method..")
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	b1 := make([]byte, 1024)
	len, err := file.Read(b1)
	fmt.Println("length of read bytes is :", len)

	fmt.Println(string(b1))

}

func main() {
	// fun()
	createFile()
	readFile("text.txt")
	readUsingOpen("text.txt")

}
