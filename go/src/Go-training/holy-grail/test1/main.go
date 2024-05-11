package main

import "fmt"


func main() {


	// In go there are 2 diiferent ways to define types
	// 1. Simple type
	// 2. Aggregate or Composite type

	// 1. Simple type

	var x int = 23

	var y = 56.8

	z := "hola"

	var w bool

	fmt.Println(x,y,z,w)


	// Declaring a array
	var v [2]int

	v = [2]int{1,2}

	fmt.Println(v)
	
	
	// 2. Composite/Aggregate Type

	// We write composite data type values by ->

	// type{values acc. to type}

	// Ex.

	xv := []int{1,2,3,4}

	xm := map[int]string{

		1 : "Hello",
		2 : "World",
	}

	xs := struct{
		say , when string
	}{
		"hello", "world",
	}


	fmt.Print(xv,"\n",xm,"\n",xs,"\n" )

	type person struct {

		first , last string
		age int

	}

	type secretagent struct {

		person
		ltk bool

	}
}
