package main

import "fmt"

func main() {

	// In Go functions are treated as first-class citizens , means 
	// in a given programming language design, a first-class citizen[a] is an entity which supports all the operations generally available to other entities. These operations typically include being passed as an argument, returned from a function, and assigned to a variable.

	// Expressions -> It is a combination of variables, operators and function calls that evaluates to a single value.



	x := foo()

	fmt.Println(x)

	y := bar()

	fmt.Println(y())

	z := cat()

	z()

	fmt.Printf("%T\t%T\t%T\n",x,y,z)

	fmt.Printf("%T\t%T\t%T\n", foo, bar, y)
}

func foo() int{
	return 42
}

func bar() func() int {

	return func() int {
		return 32
	}
}

func cat() func() {
	return func() {
		fmt.Println("I cat!")
	}
}
