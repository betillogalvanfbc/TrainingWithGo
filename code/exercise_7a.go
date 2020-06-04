package main

import "fmt"

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

var u = User{
	ID:        1,
	FirstName: "Beto",
	LastName:  "Galvan",
	Email:     "beto@galvan.com",
}

func updateEmail(u *User) {
	u.Email = "nw@gmail.com"
	fmt.Println("in update email:", u.Email)
}

func main() {
	fmt.Println("Pointers!")
	updateEmail(&u)
	fmt.Println("Update User:", u)
}
