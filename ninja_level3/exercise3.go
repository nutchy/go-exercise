package main

import "fmt"

// Create a for loop using this syntax (for condition {})
// Have it print out the years you been alive

func main() {
	y := 1997
	for y <= 2020 {
		fmt.Println(y)
		y++
	}
}
