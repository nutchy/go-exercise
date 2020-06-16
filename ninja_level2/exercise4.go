package main

import "fmt"

// Write a program that
// - Assigns an int to a variable
// - prints that int in decimal, binary, and hex
// - shift the bits of that int over 1 position to the left, and assign that to variable
// - prints that variable in decimal, binary, and hex

func main() {
	x := 42
	fmt.Printf("Decimal: %d \n", x)
	fmt.Printf("Binary: %b \n", x)
	fmt.Printf("Hex: %#x \n", x)

	y := x << 1 // shift left 1 position (add 0 into the end of the x in binary)
	fmt.Printf("Decimal: %d \n", y)
	fmt.Printf("Binary: %b \n", y)
	fmt.Printf("Hex: %#x \n", y)
}
