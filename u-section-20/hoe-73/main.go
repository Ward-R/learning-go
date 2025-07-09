package main

// a wrapper function provides an additional layer of abstraction or functionality around an existing function. often used in timeing, logging, error handling, etc.

import (
	"fmt"
	"time"
)

// Wrapper function for adding timing information
func TimedFunction(fn func()) {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	fmt.Println("Elapsed time:", elapsed)
}

// Function to be wrapped
func MyFunction() {
	time.Sleep(2 * time.Second) // simulate work
	fmt.Println("MyFunction completed")
}

func main() {
	// Call the wrapped function
	TimedFunction(MyFunction)
}
