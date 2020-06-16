package main

import "fmt"

// Create type and untype constants, Print the values of the contants.
const (
	x        = 1        // untype
	y string = "Nutchy" // type
)

// Using the following the operators, write expression and assign their values to variable
// ==, <= , >=, !=, <, >
func main() {
	a := 1 == 1
	b := 1 <= 1
	c := 1 >= 1
	d := 1 != 1
	e := 1 < 1
	f := 1 > 1
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(x)
	fmt.Println(y)
}
