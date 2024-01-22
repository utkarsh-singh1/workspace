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

	wg.Add(30)

	var mu sync.Mutex


	for i := 0 ; i < gs ; i ++ {


		go func () {

			mu.Lock()
			v := count

			runtime.Gosched()
			
			v ++

			count = v

			mu.Unlock()
			
			wg.Done()
		}()

		fmt.Println("In work Goroutines", runtime.NumGoroutine(),"The value of i is",i)
	}

	wg.Wait()

	fmt.Println("Final No. of Goroutines", runtime.NumGoroutine())

	fmt.Println(count)
	
}
