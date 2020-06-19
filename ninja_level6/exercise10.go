package main

import "fmt"

// Closure is when we have "enclosed" the scoped of a variable in some code block.
// For this hands-on exercise, create a func which "encloses" the scope of a variable
func main() {
	fn := foo(10)
	fmt.Println(fn(2)) // 20
	fmt.Println(fn(3)) // 30
	fmt.Println(fn(4)) // 40

	incr := incrementer()
	fmt.Println(incr()) // 1
	fmt.Println(incr()) // 2
	fmt.Println(incr()) // 3
	fmt.Println(incr()) // 4
}

func incrementer() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func foo(multiplier float64) func(num float64) float64 {
	return func(num float64) float64 {
		return num * multiplier
	}
}
