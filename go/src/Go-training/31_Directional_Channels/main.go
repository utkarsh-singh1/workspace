package main

import "fmt"

func main() {

	// Directional channels - A channel that can be bidirectional ( send and recieve ) or uni-directional ( send or recive) genereally denoted by make(chan <- int) send, make(<- chan int) recive during creation of channels

	// Ex -

	c := make(chan int)  // Bidirectional channel
	
	cs := make(chan <- int ) // This channel can only take values , can be send values to the channel

	cr := make(<- chan int)  // Recive only channel


	fmt.Println("------------")

	fmt.Printf("%T\n", c)

	fmt.Printf("%T\n", cr)

	fmt.Printf("%T\n", cs)


	// Assigning values from others to here

	cr = c
	cs = c

	// But from specfic to general doesnot

	// c = cr
	// c = cs

	// s := (<-chan int)(c)
	// s1 := (chan-> int)(c)
	
}
