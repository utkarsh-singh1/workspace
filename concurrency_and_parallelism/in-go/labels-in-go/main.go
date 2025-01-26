package main

import (
	"fmt"
	//"go/format"
)

func main() {

forloop:
	for i := 0 ; i < 10; i++{

		fmt.Println("Hello World")

	}

	fmt.Println("First print this")
	goto forloop
	fmt.Println("Then print that")
}
