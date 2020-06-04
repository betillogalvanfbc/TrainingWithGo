package main

import (
	"fmt"
)

func main() {
	// var myArray [5]int
	// var mySlice []int = make([]int, 5, 10)

	// myArray[0] = 1
	// mySlice[0] = 1

	// myArray[0] = 1
	// mySlice[0] = 1

	// fmt.Println(myArray)
	// fmt.Println(mySlice)
	// fmt.Println(len(mySlice))
	// fmt.Println(cap(mySlice))

	fruitArray := [5]string{"banana", "pear", "apple", "watermelon"}

	splicedFruit := fruitArray[1:3]
	//fmt.Println(splicedFruit)
	//fmt.Println(len(splicedFruit))
	//fmt.Println(cap(splicedFruit))

	fruitToAdd := append(splicedFruit, "cantelaoupe", "cherries", "lemon")

	fmt.Println(splicedFruit, fruitToAdd)
	fmt.Println(len(splicedFruit), cap(splicedFruit))
	fmt.Println(len(fruitToAdd), cap(fruitToAdd))
}
