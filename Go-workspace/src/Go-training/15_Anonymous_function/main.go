package main

import "fmt"

func main() {

	// Normal function syntax

	// func (r reciever) identifier(s type) (return(s)) { code }


	// Anonymous function here

	// func () (return(s)){ code }()

	// Ex.

	func () {
		fmt.Println("hello world")
	}()

	x := func(s string) {
		fmt.Println("Hello My name is ",s)
	}

	x("Utkarsh")
	
	
}
