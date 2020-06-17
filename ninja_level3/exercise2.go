package main

import "fmt"

// Print A-Z
// Each char, print 3 times
func main() {
	for i := 65; i <= 90; i++ {
		println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U\n", i)
		}
	}
}
