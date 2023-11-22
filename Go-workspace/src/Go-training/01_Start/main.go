package main

import (
	"fmt"
	welcome "go-training/01_Start/hello_world"
)

func main() {
	fmt.Println("Hello_World")
	x := welcome.Hello("haha")
	fmt.Println(x)
}
