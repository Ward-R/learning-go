package main

import (
	"fmt"
)

type person struct {
	first string
}

func changeWithValue(p person, s string) {
	p.first = s
	fmt.Printf("change by value name is : %s\n", p.first)
}

func changeWithPointer(p *person, s string) {
	p.first = s
	fmt.Printf("change by pointer name is : %s\n", p.first)
}

// hoe #77
func main() {
	x := person{first: "Dave"}
	fmt.Printf("Original name: %v\n", x.first)
	changeWithValue(x, "Ryan")
	fmt.Printf("Original name still unchanged?: %v\n", x.first)

	changeWithPointer(&x, "Kanta")
	fmt.Printf("Original name still unchanged?: %v\n", x.first)

}
