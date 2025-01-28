package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go sendTo(c1)

	go sendfromTo(c1, c2)

	for v := range c2 {

		fmt.Println("Value from c2", v)

	}

	fmt.Println("All values counted")
}

func sendTo(c1 chan int) {
	for i := 0; i < 100; i++ {
		c1 <- i
	}
	close(c1)
}

func sendfromTo(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			c2 <- doingSomeWork(v2)
			wg.Done()
		}(v)

	}
	wg.Wait()
	close(c2)
}

func doingSomeWork(n int) int {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	return n + rand.Intn(500)
}
