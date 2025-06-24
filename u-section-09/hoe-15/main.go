package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

func main() {
	fmt.Print("Hello World from terminal/nano\n")
	s1 := puppy.Bark()
	s2 := puppy.BigBarks()

	fmt.Println(s1)
	fmt.Println(s2)
}
