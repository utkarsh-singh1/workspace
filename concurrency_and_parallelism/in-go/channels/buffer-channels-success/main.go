package main

import "fmt"

func main() {

	// A certain method to create a channel and let it hold the value and get back at the same time from the same goroutine that created is buffer-channels, but using to much buffer-channels is not a good idea as it might cost some high memory bandwidth.

	// as make syntax is make(t T, size ...int), it can take size of type int of any magnitude - 1,2,3...infinite.
	// Below code says this channel can hold 1 value in its buffer.
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)

}
