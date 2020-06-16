package main

// Create your own type

import "fmt"

type nutchy int

var x nutchy
var y int

func main() {
	x = 42
	y = int(x)
	fmt.Printf("%v\n", x) // Print valur
	fmt.Printf("%T\n", x) // Print Type
	fmt.Printf("%v\n", y) // Print value
	fmt.Printf("%T\n", y) // Print Type
}
