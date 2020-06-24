package main

import "fmt"

// Use the "defer" keyword to show that a deferred func runs after the func containing it exits

func foo() {
	fmt.Println("Hello from foo")
}

func bar() {
	defer func() {
		fmt.Println("Hello from inner bar")
	}()
	fmt.Println("Hello from bar")
}

func main() {
	// defer will execute bottom up (Stack)
	defer foo()
	defer bar()
	fmt.Println("Main")
}
