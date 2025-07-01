package main

import (
	"fmt"
)

// hoe #49
func main() {
	m := map[string][]string{
		"bond_james":       {"shaken, not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": {"intelligence", "literature", "computer science"},
		"no_dr":            {"cats", "ice cream", "sunsets"},
	}

	// hoe #50
	m["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	// hoe #51
	delete(m, "no_dr")

	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}

}
