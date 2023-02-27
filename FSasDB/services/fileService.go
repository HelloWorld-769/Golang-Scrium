package services

import (
	"encoding/json"
	"fmt"
	mod "fs-db/models"
	"io/ioutil"
	"os"
)

const datafile = "user.json"

func GetUser() ([]mod.User, error) {

	file, err := os.Open(datafile)
	if err != nil {
		fmt.Println("Opening file error:", err)
		return nil, err
	}

	defer file.Close()

	data, _ := ioutil.ReadAll(file)
	var users []mod.User
	json.Unmarshal(data, &users)
	return users, nil
}

func SavetoFile(users []mod.User) error {
	data, err := json.Marshal(users)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(datafile, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
