package main

import (
	"fmt"
	//"math"
)

// example

// type circle struct {
// 	radius float64
// }

// type shape interface {
// 	area() float64
// }

// func (c *circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func info(s shape) {
// 	fmt.Println("area", s.area())
// }

// func main() {
// 	c := circle{5}
// 	// info(c)
// 	info(&c)
// 	// fmt.Println(c.area())
// }

// my code

type person struct {
	first string
}

// This interface defines "BEHAVIOUR" of humans. They can speak.
type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "James",
	}

	saySomething(&p1)

	//saySomething(p1) will not work because:
	// Receiver			Values
	// (t T)					T and *T
	// (t *T)					*T
}
