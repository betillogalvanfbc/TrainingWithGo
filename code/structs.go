package main

import "fmt"

//User is a user type
// type User struct {
// 	ID       int
// 	Email    string
// 	FirsName string
// 	LatName  string
// }

//User is a user type
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

func describeGroup(g Group) string {
	if len(g.users) > 2 {
		g.spaceAvailable = false
	}
	desc := fmt.Sprint("This user group has %d. The newest user is %s %s.Accepting New Users: %t", len(g.users), g.newestUser.FirsName, g.newestUser.LastName, g.spaceAvailable)
	return desc
}

///fun describeGroup
// => "This user has 19 users. The newest user is joe Smith.Accepting New Users:"

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

	fmt.Println(describeUser(u))
	fmt.Println(describeGroup(g))
	fmt.Println(g)
	//fmt.Println(u.Email)

}
