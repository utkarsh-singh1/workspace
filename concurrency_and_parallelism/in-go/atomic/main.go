package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	// Declare Variables

	var wg sync.WaitGroup

	const gs = 100

	var counter int64 
	
	fmt.Println("Current number of CPUS is", runtime.NumCPU(), "and Current number of Goroutines is",runtime.NumGoroutine())

	wg.Add(gs)

	// Increase the counter variable
	for i := 0 ; i < gs ; i++{


		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Printf("The current value of counter is %d\n",atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Current number of launched goroutines are",runtime.NumGoroutine())

	}

	wg.Wait()

	fmt.Println("Current count of counter is", counter)
	fmt.Println("Current number of goroutines is",runtime.NumGoroutine())
}
