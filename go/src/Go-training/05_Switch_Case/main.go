package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// switch-case
	switch {
	case false:
		fmt.Println("HelloWorld") ;
	case 5==5:
		fmt.Println("Yes i'm true")
	default:
		fmt.Println("A default we got")
	}

	// swicth - expretion - case

	x := 45

	switch x {
	case 1:
		fmt.Println("It is here")
	case 45:
		fmt.Println("Not here")
	}

	//or

	switch {
	case x>42:
		fmt.Println("X is greater than 42")
	case x == 45:
		fmt.Println("X is 45")
	case x<42:
		fmt.Println("x is greater than 42")
	default:
		fmt.Println("i don't know it may or may not true")
	}

	// Now like switch cases there is another control flow c/a select

	ch1 := make(chan int)
	ch2 := make(chan int)

	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))


	go func() {
		time.Sleep(d1*time.Millisecond)

		ch1 <- 41
	}()

	go func() {
		time.Sleep(d2*time.Millisecond)
		ch2 <- 42
	}()

	select {

	case v1 := <-ch1 :
		fmt.Println("value from channel 1",v1)
	

	case v2 := <-ch2 :
		fmt.Println("value from channel 2",v2)
	}
	
}


