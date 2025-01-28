package main

import (
	"fmt"
	"sync"
)

func main() {

	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	// send
	go send(even, odd)

	// recieve
	go recieve(even, odd, fanIn)

	for val := range fanIn {

		fmt.Println(val)
	}

}

func send(even, odd chan<- int) {

	for i := 0 ; i < 10 ; i++ {

		if i % 2 == 0 {
			even <- i
		}else{
			odd <- i
		}

	}

	close(even)
	close(odd)
}

func recieve(even, odd <-chan int, fanin chan<- int) {

	var wg sync.WaitGroup

	wg.Add(1)
	
	go func(){
		for v := range even {

			fanin <- v
		}
		wg.Done()
	}()

	go func(){

		for v := range odd {

			fanin <- v

		}
		wg.Done()

	}()

	wg.Wait()
	close(fanin)
}

