package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// m := map[string]int{
	// 	"James":      42,
	// 	"Moneypenny": 32,
	// }

	// // hoe #36
	// // for k, v := range m {
	// // 	fmt.Println("range of m key, value", k, v)
	// // }

	// //HOE #37
	// m1 := m["James"]
	// fmt.Println(m1)

	// // because this doesnt exist in the map it will return 0
	// m1 = m["Q"]
	// fmt.Println("That is the age of Q", m1)

	// if v, ok := m["Q"]; !ok {
	// 	fmt.Println("There is no Q, and here is the zero value of an int", v)
	// }

	// if v, ok := m["James"]; ok {
	// 	fmt.Println("There James, and here is the value of mr bond's age", v)
	// }

	// HOE #38
	for i := 0; i < 100; i++ {
		if x := rand.IntN(5); x == 3 {
			fmt.Printf("iteration %v \t x is %v\n", i, x)
		}
	}
}
