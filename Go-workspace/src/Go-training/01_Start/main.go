package main

import (
	"fmt"
	hello "go-training/01_Start/hello_world"
)

func main() {
	fmt.Println("Hello_World")
	x := hello.Hello("haha")
	fmt.Println(x)
	hello.Bolo()
	fmt.Println(hello.Kaho)
}
