package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	c := make(chan int)

	
	// send to channel
	go foo(c)
	// receive from channel
	// wg.Add(1)
	bar(c)

	// wg.Wait()
	fmt.Println("Finished running the channels 😁")

}

func foo(c chan<- int) {
	c <- 42
}

func bar(c <-chan int) {
	//defer wg.Done()
	fmt.Println(<-c)
}
