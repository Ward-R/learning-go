package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%v\n", i)
	}

	for i := 0; i < 101; i++ {
		var x int = rand.IntN(10)
		y := rand.IntN(11)
		fmt.Printf("The value of x is %v y is %v \n", x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Println("The value of x and y is less than 4")
		case x > 6 && y > 6:
			fmt.Println("The value of x and y is greater than 6")
		case x >= 4 && x <= 6:
			fmt.Println("The value of x is less than or equal to 4 and greater than or equal to 6")
		case y != 5:
			fmt.Println("The value of y is not 5")
		default:
			fmt.Println("None of the cases were met")
		}
	}

}
