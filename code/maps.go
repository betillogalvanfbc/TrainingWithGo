package main

import "fmt"

func main() {
	// var userEmails map[int]string = make(map[int]string)

	// userEmails[1] = "beto@beto.com"
	// userEmails[2] = "beto@mail.com"

	// fmt.Println(userEmails)
	userEmails := map[int]string{
		1: "user1@gmail.com",
		2: "user2@gmail.com",
	}
	userEmails[1] = "user3@gmail.com"

	fmt.Println(userEmails)

	firstEmail, ok := userEmails[4]
	fmt.Println(firstEmail, ok)

	if _, ok := userEmails[4]; ok {
		fmt.Println("email exists!")
	} else {
		fmt.Println("email doesn't exist")
	}

	delete(userEmails, 2)
	fmt.Println(userEmails)
}
