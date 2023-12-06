package main

import (
	"fmt"
	"go-training/04_Control_flow/controlflow"
	"math/rand"
)

func main() {


	x := 56
	y := 78

	z := controlflow.GreaterThen(x,y)
	fmt.Println(z)

	i := controlflow.LessThen(x,y)
	fmt.Println(i)

	// Logical Operator

	if x < 60 && y > 57 {
		fmt.Println("Then only run this code")
	}

	if x > 67 || y > 70 {
		fmt.Println("Only Then run this code")
	}

	if x != 43 && y != 78 {
		fmt.Println("Somehow i agree with you")
	}

	if !false {
		fmt.Println("How dare you say that ")
	}

	// Statements ; Statements idioms

	
	if x := 5 ; x<=5 {

		fmt.Println("bye")
	}else {
		fmt.Println("Hello")
	}

	if a := 2 * rand.Intn(4); a > 4 {
		fmt.Printf("%v is something in front of 4\n",a)
	}else {
		fmt.Printf("%v is nothing in front of 4\n",a)
	}

	// comma ok ; idioms

	if success , ok := 3,4; ok == 45 {

		fmt.Println(ok)
		
	}else {
		fmt.Println(success)
	}
	
}
