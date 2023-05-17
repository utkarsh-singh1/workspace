package main

import "fmt"

func main() {

	if true {
		fmt.Println("001")
	}

	if false {
		fmt.Println("002")
	}

	if !true {
		fmt.Println("003")
	}

	if !false {
		fmt.Println("004")
	}

	if 2 == 2 {
		fmt.Println("005")
	}

	if 2 != 2 {
		fmt.Println("006")
	}

	if !(2 == 2) {
		fmt.Println("007")
	}

	if !(2 != 2) {
		fmt.Println("008")
	}

	// In general in go we don't have to specify semi-colon(;) to seperate 2 statements , compiler does for us but if we want 2 statements on one line we could add one ;.
	fmt.Println("haleluja")
	fmt.Println("Hoh")

	// here scope of variable x is upto this (if) block
	if x := 23; x < 78 {
		fmt.Println("Yes got ya")
	}

	// try uncomment next line
	// fmt.Println(x)

	// Now if - else

	z := 33

	if z == 31 {
		fmt.Println("here it is not present")
	} else if z == 32 {
		fmt.Println("But here it is")
	} else {
		fmt.Println("here it is not present")
	}

}
