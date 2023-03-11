package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	//gorm.Model
	Id    int    `json:"id" gorm:"primary key ;default:auto_random()"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var DB *gorm.DB

func init() {
	fmt.Println("Connecting to database...")
	dsn := "host=localhost user=postgres password=Test@123 dbname=testdb sslmode=disable port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in connecting the database: ", err)
	}
	db.AutoMigrate(&User{})
	DB = db
}
func CreateHandlder(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	result := DB.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Return response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)

}
func GetHandlder(w http.ResponseWriter, r *http.Request) {
	users := &[]User{}
	DB.Find(users)
	// if err != nil {
	// 	panic(err)
	// }

	json.NewEncoder(w).Encode(&users)
}
func DeleteHandlder(w http.ResponseWriter, r *http.Request) {
	user := User{}
	id := 2

	result := DB.Delete(&user, id)
	fmt.Println(result)
	w.Write([]byte("User Deleted Succesfully"))
	json.NewEncoder(w).Encode(&user)
}

func UpdateHandlder(w http.ResponseWriter, r *http.Request) {
	var user User
	id := 1
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	result := DB.Model(&User{}).Where("id=?", id).Updates(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Return response
	w.Write([]byte("Sucess"))
}
func main() {
	fmt.Println("Listening to server at port: 8080")
	mux := http.NewServeMux()
	mux.HandleFunc("/create", CreateHandlder)
	mux.HandleFunc("/update", UpdateHandlder)
	mux.HandleFunc("/delete", DeleteHandlder)
	mux.HandleFunc("/get", GetHandlder)
	//mux.HandleFunc("")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
