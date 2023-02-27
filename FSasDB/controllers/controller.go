package cont

import (
	"encoding/json"
	"fmt"
	mod "fs-db/models"
	services "fs-db/services"
	"net/http"
	"strconv"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("register Handler working...")

	var user mod.User
	users, _ := services.GetUser()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error from register handler:", err)
		return
	}
	user.Id = len(users) + 1
	err = user.Validate()
	if err != nil {
		w.Write([]byte(fmt.Sprint(err)))
		return
	}
	users = append(users, user)
	services.SavetoFile(users)
	w.Write([]byte("User registered Sucessfully"))

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	users, _ := services.GetUser()
	res, _ := json.Marshal(users)
	w.Write([]byte(res))

}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	temp := r.URL.Query()["id"]
	id, _ := strconv.Atoi(temp[0])

	users, err := services.GetUser()
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	found := false
	for i, val := range users {
		if val.Id == id {
			found = true
			users = append((users)[:i], (users)[i+1:]...)
			var user mod.User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.Id = len(users) + 1
			err = user.Validate()
			if err != nil {
				fmt.Println("shfsjffjsd", err)
				break
			}
			users = append(users, user)
			services.SavetoFile(users)
			return
		}
	}
	if !found {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Deleteing handler working..")

	temp := r.URL.Query()["id"]
	id, _ := strconv.Atoi(temp[0])
	users, _ := services.GetUser()
	for i, val := range users {
		if val.Id == id {
			users = append((users)[:i], (users)[i+1:]...)
			services.SavetoFile(users)
			w.Write([]byte(fmt.Sprint(users)))
			return
		}
	}

}
