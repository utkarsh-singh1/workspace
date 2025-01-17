package main

import (
	"fmt"
	"runtime"
	"sync"
)



func main() {

	// Declaring Variables
	var wg sync.WaitGroup
	
	const gr = 100

	var counter int
	
	var mu sync.Mutex

	// Add wait group
	
	wg.Add(gr)

	// Intializing Increament Counter

	fmt.Println("Initial goroutines is",runtime.NumGoroutine(),"And Initial counter is",counter)
	
	for i := 0 ; i < gr ; i++{

		go func() {
			mu.Lock()
			runtime.Gosched()
			counter = counter + 1
			fmt.Println("Current Count of counter is",counter)
			mu.Unlock()
			wg.Done()
		}()

		fmt.Println("Current Number of Goroutine is",runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Printf("Current counter count is %d\n",counter)
	
	fmt.Println("Launched 100 Goroutines")

}
