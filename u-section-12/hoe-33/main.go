package main

import (
	"fmt"
)

func main() {
	// hands on exercise #33
	// for i := 1; i <= 100; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println("Skipping even number: ", i)
	// 		continue
	// 	}
	// 	fmt.Println("Printing odd number:", i)
	// }

	// hands on exercise #34
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("outside loop iteration: ", i)
	// 	for j := 0; j < 5; j++ {
	// 		fmt.Println("\t inside loop iteration: ", j)
	// 	}
	// }

	// hands on exercise #35
	xi := []int{42, 43, 44, 45, 46, 47}

	for _, num := range xi {
		fmt.Println("range of xi numbers", num)
	}
}
