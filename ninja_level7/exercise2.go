package main

import "fmt"

// Create a person struct
// Create a func called "changeMe" with a *person as a paremeters
// - change the value stored at the *person address
// -- to deference a struct, use (*value).field
// --- p1.first
// --- (*p1).first
// -- As an exception, if the type of x ia a named pointer type
//    and  (*x).f is a valid selector expression denoting a field (but not a method)
//    x.f is shorthand for (*x).f
// Create a value of type person -> print out the value
// in func main call "changeMe"

type person struct {
	first string
}

func main() {
	p1 := person{
		first: "Nutchy",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	p.first = "Chayanon"

	// this also works
	// (*p).first = "Thongpila"

}
