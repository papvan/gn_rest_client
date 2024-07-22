package gorest

import "fmt"

type usersResponse struct {
	Users []User
}

type User struct {
	ID     int    `json: "id"`
	Name   string `json: "name"`
	Email  string `json: "email"`
	Gender string `json: "gender"`
	Status string `json: "status"`
}

func (u User) Info() string {
	return fmt.Sprintf("ID: %d | Name: %s | Email: %s |  Gender: %s |  Status: %s", u.ID, u.Name, u.Email, u.Gender, u.Status)
}
