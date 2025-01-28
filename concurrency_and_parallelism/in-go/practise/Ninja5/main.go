package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {

	// Declaration

	var counter int64
	const gs = 10

	wg.Add(gs)

	for i := 0 ; i < gs ; i++ {

		go func(){
			atomic.AddInt64(&counter, 1)
			fmt.Printf("Current value of counter is %d\n",atomic.LoadInt64(&counter))
			wg.Done()
		}()

		fmt.Println("Current Number of Goroutines are", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("\nFinal counter of counter is", counter)
	
	fmt.Println("\nFinished 😌 running the process")
	
}
