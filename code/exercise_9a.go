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

func (u User) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirsName, u.LastName, u.Email)
	return desc

}

func (g *Group) describe() string {
	if len(g.users) > 2 {
		g.spaceAvailable = false
	}
	desc := fmt.Sprint("This user group has %d. The newest user is %s %s.Accepting New Users: %t", len(g.users), g.newestUser.FirsName, g.newestUser.LastName, g.spaceAvailable)
	return desc

}
func main() {
	u := User{ID: 1, FirsName: "Beto", LastName: "Galvan", Email: "beto@galvan.com"}
	u2 := User{ID: 2, FirsName: "Beto", LastName: "Galvan", Email: "beto@galvan.com"}
	u3 := User{ID: 3, FirsName: "Beto", LastName: "Galvan", Email: "beto@galvan.com"}

	g := Group{
		role:           "admin",
		users:          []User{u, u2, u3},
		newestUser:     u2,
		spaceAvailable: true,
	}

	fmt.Println(g.describe())
	fmt.Println(u.describe())
	fmt.Println(g)
}
