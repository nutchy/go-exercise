package main

import "fmt"

// Print out the remainder (modulus) which is not found for each number between 10 and 100
// when it is divided by 4

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("When %d is devided by 4, the remainder (modulus) is %d\n", i, i%4)
	}
}
