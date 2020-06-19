package main

import "fmt"

// Closure

func main() {
	fn := foo()
	fn()
}

func foo() func() {
	return func() {
		fmt.Println("Hello")
	}
}
