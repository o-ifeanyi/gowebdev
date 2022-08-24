package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

func StoreUsers(m map[string]User) {
	f, err := os.Create("users")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(m)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func LoadUsers() map[string]User {
	m := map[string]User{}
	f, err := os.Open("users")
	if err != nil {
		return m
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&m)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return m
	}
	return m
}
