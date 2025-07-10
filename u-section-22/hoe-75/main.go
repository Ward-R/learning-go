package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("%v is of %T, deref value is at that location %v\n", a, a, *a)
	fmt.Printf("%v is of %T, deref value is at that location %v\n", b, b, *b)
	fmt.Printf("%v is of %T, deref value is at that location %v\n", c, c, *c)
	fmt.Printf("%v is of %T, deref value is at that location %v\n", d, d, *d)
}
