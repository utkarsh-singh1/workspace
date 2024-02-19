package main

import "fmt"

func main() {


	f := increament()

	fmt.Println(f())
	fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())
	
}

func increament() func() int {

	x := 0
	return func() int {
		x++
		return x
	}
}
