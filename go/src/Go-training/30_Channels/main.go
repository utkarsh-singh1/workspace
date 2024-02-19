package main

import (
	"fmt"
	"runtime"
)


func main() {


	c :=  make(chan int)

	// c <- 42 // Channel blocks - the sender and the reciever transaction happens at the same time. So channel can have a value send to it but it can not recive at the same time becasue c <- 42 and <-c is a sequential event not a parallel or cocurrent one, 

	fmt.Println("Intial no. of gorutines",runtime.NumGoroutine())

	go func(){

		c <- 42
		
	}()


	// So it was necessary to have a 2 different goroutines where reciever and sender transaction happen at the same time , as from block `func main()` there was 1 goroutine was launched and another goroutine was launched when code run go func(){}(), 
	
	fmt.Println("Final no. of goroutines",runtime.NumGoroutine())
	
	fmt.Println(<-c)


	// Another way of receiving values from channels are buffer channel


	// syntax -> c := make(chan type_name, buffer_size)
	
	c1 := make(chan int, 3) // Buffer channels is a type of channel that can take certain amount of values and store it without worring when it should be taken off.

	c1 <- 32
	c1 <- 45
	c1 <- 56

	fmt.Println(<-c1)
	fmt.Println(<-c1)
	fmt.Println(<-c1)

	
}
