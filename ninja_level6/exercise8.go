package main

import "fmt"

// create a func which returns a func
// assign the returned func to a variable
// call the returned func

func main() {
	fn := foo()
	fn()
}

func foo() func() {
	return func() {
		fmt.Println("Hello")
	}
}
