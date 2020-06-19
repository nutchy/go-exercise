package main

import "fmt"

// Create a user define strut with
// - the identifier "person"
// - the field: first, last, age
// attach a method to type person with the identifier "speak"
// a method should have the person say thier name of age
// create a value of type person
// call the method from the value of type person

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}

func main() {
	p := person{
		first: "Chayanon",
		last:  "Thongpila",
		age:   23,
	}
	p.speak()
}
