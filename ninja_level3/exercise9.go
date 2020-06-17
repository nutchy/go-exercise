package main

import "fmt"

// Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string
// with the IDENTIFIER "favSport"

func main() {
	favSport := "tennis"
	switch favSport {
	case "swimming":
		fmt.Println("I Like to swimming")
	case "tennis":
		fmt.Println("I Like to tennis")
		fallthrough
	default:
		fmt.Println("This line should be print (by using fallthrough)")
	}
}
