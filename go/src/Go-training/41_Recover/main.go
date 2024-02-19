package main

import "fmt"

func main() {

	f()

	fmt.Println("There we recover function f")
	
}


func f() {


	defer func() {

		if r := recover() ; r != nil {
			fmt.Println("Recovered in f",r  )
		}
		
	}()

	fmt.Println("Calling from g")
	
	g(0)

	fmt.Println("Recovered from g")
}

func g(i int) {


	if i > 3 {
		fmt.Println("We are panicking")
		panic(fmt.Sprintf("%v",i))
	}

	defer fmt.Println("defer in g ->", i)

	fmt.Println("Printing in g ->",i)

	g(i+1)
	
}
