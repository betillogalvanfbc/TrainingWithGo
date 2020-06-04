package main

import "fmt"

//func average(num1, num2, num3 int) {
// func average(num1, num2, num3 float64) float64 {
// 	total := num1 + num2 + num3
// 	return total / 3

// }

func average(numbers ...float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers))
}

func main() {
	fmt.Println(average(15, 15, 15, 50, 30, 102, 203))
}
