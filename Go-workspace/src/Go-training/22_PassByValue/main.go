package main

import "fmt"

func main() {

	// When to use pointers

	/*

	   1. When you have a huge chunk of data and you don't want to pass it around, you create a memory location and store the data there and you have address of that location , now you can access your data by only using address of that memory location where that data is stored.

	   2. You want to change the data at a location, a certain chaneg in value at that location , get the address , derefrence it and assign new value to it.

	*/

	x := 0

	fmt.Println("Before foo x is", x)

	foo(x)

	fmt.Println("After foo x is", x)

	fmt.Printf("\n")

	fmt.Println("Before bar x is", x)
	fmt.Println("Before bar", &x)

	bar(&x)

	fmt.Println("After bar x is", x)
	fmt.Println("After barx is ", &x)

	xi := []int{1, 2, 3, 4}

	fmt.Println("Before sliceDelta xi is ", xi)

	sliceDelta(xi)

	fmt.Println("After sliceDelta xi is", xi)

	m := make(map[string]int)

	m["Utkarsh"] = 32

	fmt.Println("Before mapDelta m is ", m)

	mapDelta(m, "Utkarsh")

	fmt.Println("After mapDelta m is", m)

}

/*
	o/p -->>

0
32
0

-- but why -->> because x is assigned to y so change gonna be happening to y not x
*/
func foo(y int) {

	y = 32
}

func bar(a *int) {

	*a = 43

}

func sliceDelta(ii []int) {

	ii[0] = 99
}

func mapDelta(md map[string]int, s string) {
	md[s] = 21
}
