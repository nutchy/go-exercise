package main

import (
	"fmt"
	"sync"
)

// in addition to the main goroutine, launch two additional goroutines
// each additional goroutine should print something out
// use waitgroups to make sure each goroutine finishes before your program exists

func main() {

	var wg sync.WaitGroup

	// Add number of Go routine to wait group
	wg.Add(2)

	go func() {
		fmt.Println("Routine 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("Routine 2")
		wg.Done()
	}()

	// Waiting until all wait group done (2 goroutine)
	wg.Wait()
}
