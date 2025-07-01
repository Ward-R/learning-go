package main

import (
	"fmt"
)

// hoe #55
func main() {
	type engine struct {
		electric bool
	}
	type vehicle struct {
		engine
		make   string
		model  string
		doors  int
		colour string
	}
	c1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:   "toyota",
		model:  "camry",
		doors:  4,
		colour: "red",
	}
	fmt.Println(c1)
	fmt.Println(c1.colour)

	c2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:   "subaru",
		model:  "forester",
		doors:  4,
		colour: "black",
	}
	fmt.Println(c2)
	fmt.Println(c2.engine.electric)
}
