package main

import "fmt"

func main() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// Send values to the channels
	go send(even, odd, quit)

	// receive values from channels
	receive(even, odd, quit)
}

func send(e, o, q chan<- int) {

	for i := 0; i < 20; i++ {

		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}

	}
	//close(e)
	//close(o)
	close(q)

}

func receive(e, o, q <-chan int) {

	for {
		select {

		case v := <-e:
			fmt.Println("Current even value is",v)
		case v := <-o:
			fmt.Println("Current odd values is",v)
		case v, ok := <-q:
			fmt.Println("\nLets quit this operation by status command",v,ok)
			return
		}
	}

}
