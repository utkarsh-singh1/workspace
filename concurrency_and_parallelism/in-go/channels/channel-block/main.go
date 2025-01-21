package main

import "fmt"

func main() {

	// Channels are like pipeline between 2/more different goroutine, Remember running this main.go executable field will execute any command between between "func main() {}", it helps to synchronise between goroutines, but remember if there is only one goroutine and it launches a channel and send a value to the channel, it (channel) will block , why because it synchronises, send and recieve transition at the same time must happen, or it will throw an error like here. 
	
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
	
}
