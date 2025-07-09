package main

import (
	"fmt"
)

func funCounter() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// hoe #69
func main() {

	counter := funCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

}
