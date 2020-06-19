package main

import "fmt"

// A "callback" is when we pass a func into a func as an argument

func main() {
	foo(bar)
}

func bar() {
	fmt.Println("Hello from bar")
}

func foo(fn func()) {
	fn()
}
