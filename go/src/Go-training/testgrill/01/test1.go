package main

import "fmt"

func main() {

	var i, j, k int

	var ( a int ; b string ; c bool ) 

	var ( x int  = 5 ;y string = "abc" ; z bool = true )

	var (
		p = 23
		q = "hello world"
		r = 45.90
	)

	const d = 45

	fmt.Println(i,j,k)

	fmt.Println(a,b,c)

	fmt.Println(x,y,z)

	fmt.Println(p,q,r)

	fmt.Println(d)

	// Explicit type variable def
	var a1 int = 32

}
