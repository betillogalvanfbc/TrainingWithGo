package main

import (
	"fmt"
)

func main() {
	//i := 1

	// 	for i := 1; i <= 100; i++ {
	// 		fmt.Println(i)
	// 	}
	// i := 1
	// for i <= 100 {
	// 	fmt.Println(i)
	// 	i += 1
	// }

	var sentence = "This is a sentencer"

	for index, letter := range sentence {
		//		fmt.Println("Index:", index, "Letter:", letter)
		fmt.Println("Index:", index, "Letter:", string(letter))

	}
}
