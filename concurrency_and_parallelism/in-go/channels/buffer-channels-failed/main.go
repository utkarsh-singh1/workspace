package main

import "fmt"

func main() {

	// as make syntax is make(t T, size ...int), it can take size of type int of any magnitude - 1,2,3...infinite.
	// Below code says this channel can hold 1 value in its buffer. What if we add two values to the channel. 
	c := make(chan int, 1)

	c <- 42
	c <- 43

	// it will fail
	// But to make it success we have to increase buffer size
	// c := make(chan int, 2)

	fmt.Println(<-c)

}
