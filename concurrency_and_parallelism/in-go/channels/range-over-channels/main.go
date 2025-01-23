package main

import (
	"fmt"
	"os"
)


func main(){

	c := make(chan int)

	if x , ok := <-c ; !ok {
		fmt.Println("Bye Bye 👋")
		os.Exit(1)
	}else{
		fmt.Println(x)
	}
	
	
	go func(c chan <- int) {
		for i := 0 ; i< 10; i++{
			c <- i
		}
		close(c)
	}(c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("All values are printed from channel 🙂")
}
