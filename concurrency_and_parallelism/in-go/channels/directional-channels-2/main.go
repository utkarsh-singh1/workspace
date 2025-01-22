package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	
	c := make(chan int)
	cr := make(chan <- int)
	cs := make(<- chan int)

	go func(){
		c<-42
	}()

	fmt.Println(<-c)

	// vice versa does not work
	cr = c
	cs = c

	wg.Add(1)

	go func() {
		fmt.Println("Value of cs is", cs)
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Finished the setup")
}
