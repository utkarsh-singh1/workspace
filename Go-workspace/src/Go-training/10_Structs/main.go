package main

import (
	"fmt"
	"go-training/10_Structs/struct"
)

type person struct {
	first string
	last string
	age int
}

func main() {


	// Assign a value of type person to p1
	p1 := person {
		first: "James",
		last: "Bond",
		age: 32,
	}

	p2 := person {
		first: "James",
		last: "Gatling",
		age: 23,
	}

	var p3 person

	p3 = person{
		"utkarsh",
		"singh",
		32,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	fmt.Printf("%T\t %#v\n",p1,p1)

	m1 := struct_test.Man{
		First_name: "Monkey d",
		Last_name: "luffy",
		CodeName: 007,
		
	}


	fmt.Printf("%T\t%#v\n",m1,m1)
	
}
