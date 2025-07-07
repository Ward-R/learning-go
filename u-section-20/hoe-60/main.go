package main

import (
	"fmt"
)

// hoe #60
func main() {
	// defered functions are placed on the stack to be called when the surrounding function returns. They are executed Last in, First out. so the most recently deferred is executed first.
	defer fmt.Println(3)
	defer fmt.Println(2)
	fmt.Println(1)
}
