package main

import "fmt"

func main() {

	// for init ; condition ; post {}

	x := 1

	for x*15 < 1000 {
		fmt.Printf("Loop will run once again because value of x is %d\n", x)
		x++
	}

	// range in for

	y := [5]int{1, 2, 3, 4, 5}

	z := [5]int{}

	fmt.Printf("\n")

	for key, value := range y {
		z[key] = value + 3
	}

	fmt.Println("index No.", "value")

	for i, j := range z {

		fmt.Println(" ", i, "     |  ", j)
	}

	// Modulus operator

	for i := 1; i < 20; i++ {

		if i%2 != 0 {

			continue
		}

		fmt.Println(i)

	}
}
