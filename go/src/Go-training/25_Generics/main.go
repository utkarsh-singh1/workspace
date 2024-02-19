package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)


type myalias int


func addI(a,b int) int {
	return a+b
}

func addF(a,b float64) float64 {
	return a+b
}


type myNumbers interface {
	~int | ~float64
}


func addT[T  int | float64 ] (a,b T)  T {
	return a+b
}


func addTI[T myNumbers] (a,b T) T {
	return a+b
}


type myNumContraints interface {
	constraints.Integer | constraints.Float
}



func addTC[T myNumContraints] (a,b T) T{
	return a+b
}


func main() {

	fmt.Println(addI(1,2))

	fmt.Println(addF(1.2,2.1))

	fmt.Println(addT(1,2.1))

	fmt.Println(addT(1,2))

	fmt.Println(addT(1.3,4.6))

	fmt.Println(addTI(2,3))


	var x myalias = 32

	fmt.Println(addTI(x,3))
}
