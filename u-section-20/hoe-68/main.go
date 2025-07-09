package main

import (
	"fmt"
)

// hoe #59
func main() {
	func() {
		fmt.Println("I'm an anonymous function")
	}()

}
