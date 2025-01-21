package main

import "fmt"

func main() {
	
	c := make(chan int)

	// A succesful channel as it takes input from one side and ouput from other side at the same time.
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

}
