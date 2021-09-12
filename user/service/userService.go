package service

import (
	"log"
)
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
func CreateUser(user string) string {
	log.Println(user)
	return "succesfully"
}

func GetUser(user string) string {

	return "succesfully"
}

func UpdateUser(user *User) string {
	return "succesfully"
}

func DeleteUser(user *User) string {
	return "succesfully"
}
	
