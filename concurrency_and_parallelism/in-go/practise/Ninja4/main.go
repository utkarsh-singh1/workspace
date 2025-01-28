package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	// Declaration

	var counter int64
	const gs = 10

	var mu sync.Mutex
	wg.Add(gs)

	for i := 0 ; i < gs ; i++ {

		go func(){
			mu.Lock()
			counter = counter + 1
			mu.Unlock()
			wg.Done()
		}()

		fmt.Println("Current Number of Goroutines are", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Final counter of counter is", counter)
	
	fmt.Println("\nFinished 😌 running the process")
	
}
