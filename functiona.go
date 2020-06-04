package main

import "fmt"

//DECLARE FUNCTION AGE
// func printAge(age int) int {
// 	fmt.Println(age)
// 	return age
// }

// func printAge() (ageBob int, ageJohn int) {
// 	ageBob = 20
// 	ageJohn = 21
// 	return
// }

// func main() {
// 	age1, age2 := printAge()
// 	fmt.Println(age1)
// 	fmt.Println(age2)
// }

func printAge(ages ...int) {
	for _, value := range ages {
		fmt.Println(value)
	}
}

func main() {
	printAge(15, 30, 10)
}
