package main

import (
	"fmt"
)

// this way is running a channel with an anon function
// func main() {
// 	c := make(chan int)

// 	go func() {
// 		c <- 42
// 	}()
// 	fmt.Println(<-c)
// }

// this way is running a channel with a buffer (1)
func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
