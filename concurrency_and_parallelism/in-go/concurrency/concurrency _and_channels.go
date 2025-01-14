package main

import "fmt"

func multiplyby5(x int) int{

	return x*5
}

func main() {

	ch := make(chan int)
	
	go func() {

		ch <- multiplyby5(5)

	}()

	fmt.Println(<-ch)
	

}
