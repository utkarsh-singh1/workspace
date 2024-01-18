package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {


	fmt.Println("Initial No. of CPUs", runtime.NumCPU())

	fmt.Println("Intial no. of Goroutines", runtime.NumGoroutine())
	
	var count int64 


	const gs = 100

	var wg sync.WaitGroup

	wg.Add(gs)

	
	for i := 0 ; i < gs ; i ++ {


		go func () {
			atomic.AddInt64(&count,1)

			runtime.Gosched()
			
			fmt.Println("Counter\t: ",atomic.LoadInt64(&count))

			wg.Done()
		}()

		fmt.Println("In work Goroutines", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Finished No. of Goroutines", runtime.NumGoroutine())

	fmt.Println(count)
	
}

