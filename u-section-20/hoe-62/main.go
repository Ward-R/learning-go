package main

import (
	"fmt"
	"math"
)

// This excercise I learned about polymorphism and interfaces in go. In ruby
// this is kind of like duck typing, but it is not like that here due to types.

type Square struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (s Square) Area() float64 {
	return s.length * s.width
}

type Shape interface {
	Area() float64
}

func Info(s Shape) {
	fmt.Println(s.Area())
}

// hoe #62
func main() {
	var circle = Circle{radius: 10}
	var square = Square{length: 2, width: 4}

	Info(circle)
	Info(square)
}
