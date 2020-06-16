package main

// Using iota, create 4 constants for the last 4 years, print the constant values
const (
	a = 2017 + iota // iota = 0
	b = 2017 + iota // iota = 1
	c = 2017 + iota // iota = 2
	d = 2017 + iota // iota = 3
)

// Note: To keep readable this code for all programmers
// avoid to use the iota, because it make programmer to confuse how's does iota it works

func main() {
	println(a, b, c, d)
}
