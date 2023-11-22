package main

import (
	"fmt"
	"go-training/02_Variable/variable"

	dog "gitlab.com/d3vus/go-pkgs/dog"
)

func main() {

	var x int = 32
	var y int = 43
	fmt.Println(variable.AddThis(x,y))
	fmt.Println(variable.SubtractThis(x,y))
	dog.Bark()
}
