package main

import "fmt"

func main() {

	// switch-case
	switch {
	case false:
		fmt.Println("HelloWorld") ;
	case 5==5:
		fmt.Println("Yes i'm true")
	default:
		fmt.Println("A default we got")
	}

	// swicth - expretion - case

	x := 45

	switch x {
	case 1:
		fmt.Println("It is here")
	case 45:
		fmt.Println("Not here")
	}

	//or

	switch {
	case x>42:
		fmt.Println("X is greater than 42")
	case x == 45:
		fmt.Println("X is 45")
	case x<42:
		fmt.Println("x is greater than 42")
	default:
		fmt.Println("i don't know it may or may not true")
	}

	//
}


