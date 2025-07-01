package main

import (
	"fmt"
)

// hoe #53
func main() {
	type person struct {
		first            string
		last             string
		favoriteIceCream []string
	}

	p1 := person{
		first:            "ryan",
		last:             "ward",
		favoriteIceCream: []string{"strawberry", "choco minto", "maple walnut"},
	}

	p2 := person{
		first:            "forest",
		last:             "gump",
		favoriteIceCream: []string{"vanilla", "bubblegum", "rocky road"},
	}

	fmt.Printf("first name: %s\t last name: %s\n", p1.first, p1.last)
	fmt.Println("favorite ice cream flavours:")
	for _, v := range p1.favoriteIceCream {
		println(v)
	}
	fmt.Println("************************************")
	fmt.Printf("first name: %s\t last name: %s\n", p2.first, p2.last)
	fmt.Println("favorite ice cream flavours:")
	for _, v := range p2.favoriteIceCream {
		println(v)
	}

	// hoe #54
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v)
		for _, v2 := range v.favoriteIceCream {
			fmt.Println(v.first, v.last, v2)
		}
	}
}
