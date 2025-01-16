package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	//cpu := runtime.NumCPU()
	goroutines := runtime.NumGoroutine()

	fmt.Println("Gonna Launched 100 Goroutines")
	
	var counter int

	fmt.Printf("Current counter count is %d\n",counter)

	const gs = 100

	wg.Add(gs)

	for i := 0; i < gs; i++ {

		go func() {

			runtime.Gosched()
			counter = counter + 1
		}()
		fmt.Println("Current goroutine count is",goroutines, "and counter count is",counter)
		wg.Done()
	}

	wg.Wait()

	fmt.Printf("Current counter count is %d\n",counter)
	
	fmt.Println("Launched 100 Goroutines")

}
