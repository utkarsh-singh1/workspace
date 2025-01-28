package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){

	c := fanIn(boringTalks("John"),boringTalks("Melissa"))

	for i := 0 ; i < 10 ; i++{

		fmt.Println(<-c)
	}

	fmt.Println("Stop it both of you")
}

func boringTalks(msg string) <-chan string{

	c := make(chan string)
	go func() {

		for i := 0 ; ; i++{
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3))*time.Millisecond)
		}

	}()

	return c

}

func fanIn(x, y <- chan string) <-chan string {

	c := make(chan string)

	go func(){
		for {
			c <- <-x 
		}
	}()

	go func() {
		for {
			c <- <-y
		}
	}()

	return c


}
