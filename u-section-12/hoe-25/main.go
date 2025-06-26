package main

import (
	"fmt"
	"math/rand/v2"
)

func init() {
	fmt.Println("This is where initialization for my program occurs")
}

func main() {
	var x int = rand.IntN(250)
	fmt.Printf("The value of x is %v\n", x)

	switch {
	case x > 0 && x <= 100:
		fmt.Println("The value of x is between 0 and 100")
	case x > 100 && x <= 200:
		fmt.Println("The value of x is between 100 and 200")
	case x > 200 && x <= 250:
		fmt.Println("The value of x is between 200 and 250")
	}
}
