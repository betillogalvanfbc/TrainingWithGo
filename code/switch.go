package main

import "fmt"

var city string

func main() {

	// switch city {

	// case "Tijuana":
	// 	fmt.Println("You live in Tijuana")

	// case "Rosarito":
	// 	fmt.Println("You live in Rosarito")

	// case "Mexicali":
	// 	fmt.Println("You live in Mexicali")

	// default:
	// 	fmt.Println("Don't exists in the scope!")
	// }

	var i int = 100

	switch {
	case i != 10:
		fmt.Println("Does not equal 10")
		fallthrough
	case i > 10:
		fmt.Println("Greater than 10")
	case i < 10:
		fmt.Println("Less than 10")
	default:
		fmt.Println("Is 10")
	}

}
