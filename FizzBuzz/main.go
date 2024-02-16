package main

import "fmt"

func FizzBuzz(i int) {
	if i > 100 {
		return
	}

	if i%3 == 0 {
		fmt.Println("Fizz")
	} else if i%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(i)
	}

	FizzBuzz(i + 1)
}

func main() {
	// for i := 1; i <= 100; i++ {
	// 	if i%3 == 0 {
	// 		fmt.Println("Fizz")
	// 	} else if i%5 == 0 {
	// 		fmt.Println("Buzz")
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }

	FizzBuzz(1)
}
