package main

import "fmt"

func main() {

	x := 45
	y:= 45.67

	fmt.Printf("%#v is of type %T\n", x,x)
	fmt.Printf("%#v is of type %T\n",y,y)

	var a float32 = 45.98

	fmt.Printf("%#v is of type %T\n",a,a)

	var z float64
	
	z = float64(a)

	fmt.Printf("Type of %v is %T\n",z,z)
}
