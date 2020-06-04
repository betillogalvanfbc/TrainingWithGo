package main

import (
	"fmt"
)

func main() {
	// 	var scores [5]float64
	// 	fmt.Println(scores)
	//
	// var scrores [5]float64 = [5]float64{9,1.5,4.5,7,8}
	// scores := [5]float64 {9,1.5.,4.5,7.8}

	scores := [5]float64{1, 2, 3, 4, 5}
	for _, values := range scores {
		fmt.Println(values)
	}
}
