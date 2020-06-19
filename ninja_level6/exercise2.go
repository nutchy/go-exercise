package main

import "fmt"

// Create a func foo with the identifier foo that
// - takes in a variadic parameter of type int
// - pass in a value of type []int into your func (unfurl the int (...))
// - returns the sum of all values of type int passed in

func foo(args ...int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}

// Create a func foo with the identifier foo that
// - takes in a parameter of type []int
// - returns the sum of all values of type int passed in
func bar(args []int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}

func main() {
	ii := []int{1, 2, 3, 4, 5}
	total := foo(ii...)
	total2 := bar(ii)
	fmt.Printf("%d\n", total)
	fmt.Printf("%d\n", total2)
}
