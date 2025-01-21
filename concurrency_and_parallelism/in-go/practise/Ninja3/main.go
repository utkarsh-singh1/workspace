package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	// Declaration

	counter := 0
	const gs = 10

	wg.Add(gs)

	for i := 0 ; i < gs ; i++ {

		go func(){
			v := counter
			runtime.Gosched()
			v = v + 1
			counter = v

			wg.Done()
		}()

		fmt.Println("Current Number of Goroutines are", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Final counter of counter is", counter)
	
	fmt.Println("\nFinished 😌 running the process")
	
}
