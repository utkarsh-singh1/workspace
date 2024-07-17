package main

import "fmt"

const (

	// Add 100 zeroes in  front of 1
	Big = 1 << 100


	// Remove 99 zeroes from Big

	Small = Big >> 99
)

func needInt(x int) int {

	return x*10 + 1

}

func needConst(x float64) float64 {

	return x * 0.1
	
}

func main() {

	fmt.Printf("Type of Big is %T \t \n", Big)

}
