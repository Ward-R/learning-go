package main

import (
	"fmt"
)

func main() {
	// hands on exercise #42
	// x := [5]int{}
	// for index, num := range x {
	// 	fmt.Printf("position %v is %v of %T\n", index, num, num)
	// }

	// hoe #43

	// x := [10]int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// for index, num := range x {
	// 	fmt.Printf("position %v is %v of %T\n", index, num, num)
	// }

	// // hoe #44
	// x := [10]int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// for index, num := range x {
	// 	fmt.Printf("position %v is %v of %T\n", index, num, num)
	// }
	// a := x[:5]
	// fmt.Println(a)

	// b := x[5:]
	// fmt.Println(b)

	// c := x[2:7]
	// fmt.Println(c)

	// d := x[1:6]
	// fmt.Println(d)

	// hoe #45
	// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// x = append(x, 52)
	// fmt.Println(x)
	// x = append(x, 53, 54, 55)
	// fmt.Println(x)
	// y := []int{56, 57, 58, 59, 60}
	// x = append(x, y...)
	// fmt.Println(x)

	// hoe #46
	// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// x = append(x[:3], x[6:]...)
	// fmt.Println(x)

	// hoe #47
	// states := make([]string, 0, 50)
	// states = append(states,
	// 	` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	// fmt.Println(states)
	// fmt.Println(len(states))
	// fmt.Println(cap(states))

	// hoe #48
	x := []string{"james", "bond", "shaken, not stirred"}
	y := []string{"ms", "moneypenny", "I'm 008"}
	z := [][]string{x, y}
	for i, v := range z {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
	}
}
