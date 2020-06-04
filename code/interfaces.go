package main

//Describer interface describes the struct being passed in
type Describer interface {
	describers()
}

//User is a single user type
type User struct {
	ID                         int
	Firstname, LastName, Email string
}

// Group represents a set of users
type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

func main() {

}
