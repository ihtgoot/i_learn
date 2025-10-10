package main

import "fmt"

var x int

func main() {
	x = 23
	y := 24
	n := "fu"
	fmt.Printf("hi there ,%d, %T, i am %s\n", x+y, y, n)
	var z int
	a, b := outputThat(n, 797.584, x+y, &z)
	fmt.Printf("%s, %t, %d , %p\n", a, b, z, &z)
}

func outputThat(a string, b float64, c int, d *int) (string, bool) {
	fmt.Printf("%i %p\n", *d, d)
	*d = 78
	var z bool = (b == float64(c))
	fmt.Printf("%i %p\n", *d, d)
	var name string = a + ":fuck you"
	return name, z
}
 