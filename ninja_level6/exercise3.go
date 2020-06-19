package main

import "fmt"
// Use the "defer" keyword to show that a deferred func runs after the func containing it exits

func main() {
	defer foo()
	fmt.Println("Main")
}

func foo() {
	fmt.Println("Hello from foo")
}
