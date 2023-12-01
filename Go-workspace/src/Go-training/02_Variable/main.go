package main

import (
	"fmt"
	"go-training/02_Variable/variable"

	dog "gitlab.com/d3vus/go-pkgs/dog"
)

func main() {

	var x int = 32
	var y int = 43
	fmt.Println(variable.AddThis(x, y))
	fmt.Println(variable.SubtractThis(x, y))
	dog.Bark()

	// _ is blank identifier that could assign value to itself but needed not to be used.
	a, b, _ := 2, 3, 1 == 2

	fmt.Println(a, b)

	// Zero Values

	var p, q, r int

	fmt.Println(p, q, r)

	// Declaration without type

	var l, m = 55, 56

	// Multiple declaration at same time
	var (
		i string
		j float32
		k bool
	)

	i = "hello"
	j = 45
	k = true

	fmt.Println(i, j, k, l, m)

	fmt.Printf("The binary value of a b x y l is %b %b %b %b %b\n", a, b, x, y, l)

	fmt.Printf("The hexadecimal value of a b x y l is %x %x %x %x %x", a, b, x, y, l)
}
