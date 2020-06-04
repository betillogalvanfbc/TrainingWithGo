package main

import (
	"fmt"
)

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("We panicked but everyone's fine")
		fmt.Println(r)
	}
}

func doThings() {
	defer recoverFromPanic()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 2 {
			panic("PANIC!")
		}
	}
}

func main() {
	doThings()
}

// package main

// import (
// 	"fmt"
// )

// func doThings() {
// 	defer fmt.Println("First Line")
// 	defer fmt.Println("Second Line")
// 	fmt.Println("Things And Stuff should happen first")
// }

// func main() {
// 	doThings()
// }

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// )

// func isGreaterThanThen(num int) error {
// 	if num < 10 {
// 		return errors.New("something bad happened")

// 	}
// 	return nil
// }

// func openFile() error {
// 	f, err := os.Open("file.txt")
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	return nil
// }

// func main() {
// 	num := 9
// 	// err := isGreaterThanThen(num)
// 	if err := isGreaterThanThen(num); err != nil {
// 		fmt.Println(fmt.Errorf("%d is NOT GREATER THAN TEN", num))
// 		// panic(err)
// 		// fmt.Println("Hey do somethig else")
// 		//log.Fatalln(err)
// 	}
// 	//fmt.Println(err)

// 	someOtherErr := openFile()

// 	if err := openFile(); someOtherErr != nil {
// 		fmt.Println(fmt.Errorf("%v", err))
// 	}
// }
