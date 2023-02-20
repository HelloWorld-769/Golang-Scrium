package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Person struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Age   int    `json:"age"`
	User  User   `json:"user"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing the handler")
	w.Write([]byte("OK"))
}

func middlewareHttp(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Running before handler")
		w.Write([]byte("Hijacking request "))

		fmt.Println(r.URL.Path)
		originalHandler.ServeHTTP(w, r)
		fmt.Println("Running after the handler")
	})
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	fmt.Println("Form data parsing:", username, password)
}

func handleQueryParams(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	username := query.Get("username")
	pass := query.Get("password")

	fmt.Println("Query Params", username, pass)
}

func isvalidatePass(password string) bool {

	if len(password) < 8 {
		fmt.Println("Password is too short")
		return false
	}
	hasUpperCase := false
	hasLowerCase := false
	hasNumbers := false
	hasSpecial := false

	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUpperCase = true
		} else if char >= 'a' && char <= 'z' {
			hasLowerCase = true
		} else if char >= '0' && char <= '9' {
			hasNumbers = true
		} else if char >= '!' && char <= '/' {
			hasSpecial = true
		} else if char >= ':' && char <= '@' {
			hasSpecial = true
		}
	}

	if !hasUpperCase {
		fmt.Println("Password do not contain upperCase Character")
		return false
	}

	if !hasLowerCase {
		fmt.Println("Password do not contain LowerCase Character")
		return false
	}

	if !hasNumbers {
		fmt.Println("Password do not contain any numbers")
		return false
	}

	if !hasSpecial {
		fmt.Println("Password do not contain any special character")
		return false
	}
	return true
}
func middleware(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Println("Error from middleware:", err)
		}
		fmt.Println("Value form middleware:", user)
		if !isvalidatePass(user.Password) {
			http.Error(w, "password do not match the requirements", http.StatusBadRequest)
			return
		}
		//used to pass vLUE FORM MIDDLEWARE TO HANDLER...
		ctx := context.WithValue(r.Context(), "mystruct", &user)
		originalHandler.ServeHTTP(w, r.WithContext(ctx))
	})
}
func authUsingBcrypt(w http.ResponseWriter, r *http.Request) {
	s := "Hello*1234"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	fmt.Println("Hashed password:", string(bs))
	myUser := r.Context().Value("mystruct").(*User)
	fmt.Println("My user:", myUser.Password)
	password := myUser.Password

	fmt.Println("Value from original handler:", password)

	err = bcrypt.CompareHashAndPassword(bs, []byte(password))

	if err != nil {
		http.Error(w, "Password Mismatch", http.StatusBadRequest)
		return
	}
	w.Write([]byte("Authenticated"))

}

var people []Person

func main() {
	mux := http.NewServeMux()
	myHandler := http.HandlerFunc(handler)
	mux.Handle("/", middlewareHttp(myHandler))
	mux.HandleFunc("/formSubmit", handleForm)
	mux.HandleFunc("/QuerySubmit", handleQueryParams)

	mux.Handle("/passwordAuth", middleware(http.HandlerFunc(authUsingBcrypt)))
	log.Print("Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", mux))

}
