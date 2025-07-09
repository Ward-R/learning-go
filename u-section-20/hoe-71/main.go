package main

import (
	"fmt"
)

func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

// hoe #71
func main() {
	result1 := applyOperation(5, 3, add)
	result2 := applyOperation(5, 3, subtract)

	fmt.Println(result1)
	fmt.Println(result2)
}
