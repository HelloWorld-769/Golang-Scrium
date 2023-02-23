package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
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

func postMiddleWare(c *gin.Context) {

	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("Request body: ", string(reqBody))
	var newAlbum album
	err := json.Unmarshal(reqBody, &newAlbum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.Set("body", &newAlbum)
	c.Next()

}

func main() {
	r := gin.Default()
	//tu use middleware on all the router pass the middleware in r.Use()
	//eg-> e.Use(postMiddleWare)
	r.GET("/", welcomeHandler)
	r.GET("/get", getAlbums)

	//r.POST("/post", postMiddleWare, postAlbum)
	r.POST("/post", handleBody)
	r.POST("/form", handleForm)
	r.GET("/fun", fun)
	r.Run(":3000")
}
func welcomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {

	//passing data from middleware to function post albums and then retreiving it using c.Get("key")
	reqBody, exists := c.Get("body")
	if !exists {
		fmt.Println("value do not exist")
	}

	albums = append(albums, *reqBody.(*album))
	c.JSON(http.StatusOK, albums)
	return

}

func handleBody(c *gin.Context) {
	var newAlbum album
	err := c.BindJSON(&newAlbum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Println("Body is:", newAlbum)
	c.JSON(http.StatusOK, newAlbum)

}
func handleForm(c *gin.Context) {
	formRes := c.PostForm("username")
	fmt.Println(formRes)

}
func fun(c *gin.Context) {
	var temp album
	err := c.BindQuery(&temp)
	if err != nil {
		fmt.Println("error is ", err)
	}
	fmt.Println(temp)
}
