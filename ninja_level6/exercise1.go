package main

import "fmt"

// Create a func with the identifier foo that returns an int
// Create a func with the identifier bar that returns an int and string
// Call both funcs
// print out thier results

func foo() int {
	return 10
}

func bar() (int, string) {
	return 20, "Nutchy"
}

func main() {
	n := foo()
	x, s := bar()

	fmt.Println(n, x, s)
}
