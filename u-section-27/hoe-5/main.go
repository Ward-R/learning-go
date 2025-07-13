package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// fmt.Println("CPUs:", runtime.NumCPU())
	// fmt.Println("Goroutines:", runtime.NumGoroutine())
	var wg sync.WaitGroup
	var incrementer int64

	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
		// fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
	// fmt.Println("Goroutines:", runtime.NumGoroutine())
	// fmt.Println("count:", incrementer)
}

// run with ❯ go run -race ./u-section-27/hoe-3/main.go  if you want to check at compile for race condition.
