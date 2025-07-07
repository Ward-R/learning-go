package main

import (
	"fmt"
)

// hoe #58
func main() {
	fmt.Println(foo(42))
	fmt.Println(bar(43, "fubar"))
}

func foo(a int) int {
	return a
}

func bar(b int, c string) (int, string) {
	return b, c
}
