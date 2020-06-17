package main

import "fmt"

// Create a for loop using this syntax (for {})
func main() {
	y := 1997
	for {
		if y > 2020 {
			break
		}
		fmt.Println(y)
		y++
	}
}
