package main

import "fmt"

func main() {
	var sVar = 101
	// if sVar >= 10 {
	// 	fmt.Print(sVar)
	// }

	// if sVar > 100 {
	// 	fmt.Print("Greater than 100")
	// } else if sVar == 100 {
	// 	fmt.Print("Equals 100")
	// } else {
	// 	fmt.Print("Less than 100")
	// }

	err := getError()

	if err != nil {
		fmt.Println(err.Error())
	}

	// if err := getError(); err != nil {
	// 	fmt.Println(err.Error())
	// }
}
