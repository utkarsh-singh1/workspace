package main

import "fmt"


func main() {

	// there are two types of channels bidirectional channels { make (chan int) }  and single directionals channels, Single directional channels have 2 types attached to it send values to the channels
	// make(chan <- int) and recieve value from the channels make(chan -> int).
	c := make(chan <- int)
	
	fmt.Printf("%T\n",c)

	var ch = make(<- chan int )

	fmt.Printf("%T\n",ch)

}
