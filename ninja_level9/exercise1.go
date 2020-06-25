package main

import (
	"fmt"
	"runtime"
	"sync"
)

// in addition to the main goroutine, launch two additional goroutines
// each additional goroutine should print something out
// use waitgroups to make sure each goroutine finishes before your program exists

func main() {
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of Go Routines:", runtime.NumGoroutine())

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

	fmt.Println("Number of CPUs before run go routine:", runtime.NumCPU())
	fmt.Println("Number of Go Routines  before run go routine:", runtime.NumGoroutine())
	// Waiting until all wait group done (2 goroutine)
	wg.Wait()
	fmt.Println("Number of CPUs after run go routine:", runtime.NumCPU())
	fmt.Println("Number of Go Routines  after run go routine:", runtime.NumGoroutine())
}
