package main

import (
	"fmt"
)

type person struct {
	first string
	age   int
}

// hoe #61
func main() {
	me := person{
		first: "Ryan",
		age:   450,
	}
	fmt.Println(me.speak())

}

func (p person) speak() string {
	say := fmt.Sprintf("My name is %v and I'm %v years old", p.first, p.age)
	return say
}
