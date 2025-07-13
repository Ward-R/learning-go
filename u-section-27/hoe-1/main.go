package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	// Needed to add wait groups otherwise the goroutines execute? without printing.
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hello go routine one in anonymous function")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello go routine two in anonymous function")
		wg.Done()
	}()

	fmt.Println("mid cpu", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("exiting main")
	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())
}
