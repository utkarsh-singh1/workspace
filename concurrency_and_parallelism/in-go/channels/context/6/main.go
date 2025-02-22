package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		n := 0
		for {
			n++
			time.Sleep(time.Millisecond * 200)
			fmt.Println(n)
		}
	}()

	fmt.Println("That's it now 😄")
}
