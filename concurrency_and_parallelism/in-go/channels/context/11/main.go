package main

import "fmt"

func main() {

	var x int
	fmt.Println("Enter a value -")
	fmt.Scan(&x)

	switch true {

	case 100 <= x && x <= 500:
		fmt.Println("You are eligible for a coupon of worth", 20)
	case 501 <= x && x <= 1000:
		fmt.Println("You are eligible for a coupon of worth", 50)

	}
}
