package main

import "fmt"

var x int
var y string
var z bool

func main() {
	// Print all zero value of x, y, z
	fmt.Println(x) // 0
	fmt.Println(y) // (empty string)
	fmt.Println(z) // false

	x = 42
	y = "Chayanon"
	z = true

	s := fmt.Sprintf("%v \t %v \t %v", x, y, z)

	fmt.Println(s) // false
}
