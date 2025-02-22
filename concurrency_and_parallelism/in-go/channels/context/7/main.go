package main

import "fmt"

func main() {

	fmt.Println(multiplyBy4(3))

}

func multiply(x,y int) int{

	return x*y
}

// func multiplyBy4(x int, y func(i,j int)int ) int{

// 	return y(x,4)
// }

func multiplyBy4(x int) int {

	return multiply(x,4)
}
