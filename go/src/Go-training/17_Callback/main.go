package main

import "fmt"

func main() {

	// passing a function as a argument to another function , process is k/a callback


	y := doMath(43,21,add)

	fmt.Println(y)

	x := doMath(43,21,subtract)

	fmt.Println(x)
	
	fmt.Printf("%T\t%T\t%T\n",add,subtract,doMath)

}

func doMath(x int, y int, f func(int,int) int) int {
	return f(x,y)
}


func add (x int, y int) int {
	return x+y
}

func subtract(x int, y int) int {

	return x-y

}
