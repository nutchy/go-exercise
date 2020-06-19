package main

import "fmt"

// Assign a func to a variable, then call it
func main() {
	fn := func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}
	fn()
	fmt.Println("Done")

}
