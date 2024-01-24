package main

import (
	"fmt"
	"go-training/35_Comma_ok_idiom/withtypes"
)

func main() {

	even := make(chan int)

	odd := make(chan int)

	quitint := make(chan int)

	quitbool := make(chan bool)
	


	// send


	go withtypes.Sendbool(even,odd,quitbool)

	//recieve

	withtypes.Recievebool(even,odd,quitbool)


	// send

	go withtypes.Sendint(even, odd, quitint)

	// recieve

	withtypes.Recieveint(even,odd, quitint)



	// Littile test

	c := make(chan string)


	go func() {

		c <- "hello"
		close(c)
	} ()


	v, ok := <-c
	
	fmt.Println(v,ok)

	v, ok = <-c

	fmt.Println(v,ok)
	
	fmt.Println("We done here")
	
}
