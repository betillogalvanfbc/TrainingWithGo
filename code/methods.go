package main

import (
	"fmt"
)

type User struct {
	ID                        int
	FirsName, LastName, Email string
}

// Group represents a set of users
type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirsName, u.LastName, u.Email)
	return desc
}

func (u *User) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirsName, u.LastName, u.Email)
	return desc
}

func main() {
	u := User{ID: 1, FirsName: "Beto", LastName: "Galvan", Email: "beto@galvan.com"}
	desc := describeUser(user)
	//desc := user.describe()
	desc := user.describe() 
	fmt.Println(desc)

}
