package exercise_1

import "fmt"

// Write a program that prints a number in decimal, binary, and hex
func main() {
	x := 123
	fmt.Printf("Decimal: %d \n", x)
	fmt.Printf("Binary: %b \n", x)
	fmt.Printf("Hex: %#x \n", x)
}
