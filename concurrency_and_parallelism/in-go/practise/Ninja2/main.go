package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Printf("This is the first goroutine number %d has been launched\n", runtime.NumGoroutine())

	// ok there can be use of for loop, but I want to enable for each iteration.
	wg.Add(2)
	go func() {
		fmt.Printf("This is goruntine number %d has been launched\n", runtime.NumGoroutine())
		wg.Done()
	}()
	go func() {

		fmt.Printf("This is goroutine number %d has been launched\n", runtime.NumGoroutine())
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Mission Completed 🫡")

}
