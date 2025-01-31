package main

import "fmt"

func main() {

	// Using function to assign it self to other variable
	x := multBy4

	fmt.Printf("%T",x)

	// Using function to assign its value to other variable

	y := multBy4(5)

	fmt.Println(y)
}

func multBy4(n int) (int){

	return n*4

}
