package main

import (
	"fmt"
	"math"
)

// create a type square
// create a type circle
// attach method to each that calculates area and return it
// create a type shape witch define an interface as anything which has the area method
// create a func info which takes type shape and then print the area
// create a value of type square
// create a value of type circle
// use func info to print the area of circle
// use func info to print the area of square

type square struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (s square) area() float64 {
	return s.height * s.width
}

func info(s shape) {
	fmt.Printf("%.2f\n", s.area())
}

func main() {
	c := circle{
		radius: 12.345,
	}

	s := square{
		width:  15,
		height: 15,
	}

	info(c)
	info(s)
}
