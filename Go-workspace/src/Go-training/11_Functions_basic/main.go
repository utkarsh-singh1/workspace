package main

import (
	"fmt"
	fu "go-training/11_Functions_basic/function"
)

func main() {

	fmt.Println(fu.Variadic_max(1,2,4,5,6,12,3,456,67,321,543,90,54,9,11))

	fmt.Println(fu.Variadic_min(1,2,4,5,6,12,3,456,67,321,543,90,-1,8,0,-67,54,9,109))

	xi := []int{2,5,6,3,5,8,90,43,45,21,5,7,9,5}

	fmt.Println(fu.Unfurling_sum(xi...))

	s := fu.Unfurling_sum()

	fmt.Printf("%T\n",s)

	fmt.Println(fu.Unfurling_max(xi))

	// Defer

	fmt.Println("foo1")

	fmt.Println("foo2")

	fu.Foo()

	defer fu.Bar()

	fmt.Println("bar1")

	fmt.Println("bar2")


	// Methods

	x := fu.Utkarsh {

		Age: 32,
		FullName: "utkarsh singh",
		Money: 34000,
	}
	fmt.Println(x.Age)

	fmt.Printf("%T\n",x)
	fu.Utkarsh.Details(x)
	
}

// func (r reciever) some_identifier (some_parameters some_type) (some_return) {code}


