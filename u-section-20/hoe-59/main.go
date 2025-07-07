package main

import (
	"fmt"
)

// hoe #59
func main() {
	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// call foo, pass intSlice and unfurl it.
	fmt.Println(foo(intSlice...))
	fmt.Println(bar(intSlice))
}

func foo(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func bar(numbers []int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}
