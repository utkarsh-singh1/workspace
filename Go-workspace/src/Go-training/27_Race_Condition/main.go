package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {


	fmt.Println("Initial No. of CPUs", runtime.NumCPU())

	fmt.Println("Intial no. of Goroutines", runtime.NumGoroutine())
	
	count := 0


	const gs = 100

	var wg sync.WaitGroup

	wg.Add(gs)


	for i := 0 ; i < gs ; i ++ {


		go func () {
			v := count

			runtime.Gosched()
			
			v ++

			count = v

			wg.Done()
		}()

		fmt.Println("Goroutines", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Goroutines", runtime.NumGoroutine())

	fmt.Println(count)
	
}
