package main

import "fmt"

// Create a program that shows the "if statement" in action
func main() {
	x := "bababa"
	if x == "Nutchy" {
		fmt.Println(x)
	} else if x == "bababa" {
		fmt.Println("Hello:", x)
	}
}
