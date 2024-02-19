package main

import (
	"fmt"
	struct_test "go-training/10_Structs/struct"
)

type person struct {
	first string
	last string
	age int
}

type secretAgent struct {
	admi person
	ltk bool
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


	// Embedded struct
	m1 := struct_test.Man{
		First_name: "Monkey d",
		Last_name: "luffy",
		CodeName: 007,
		
	}


	fmt.Printf("%T\t%#v\n",m1,m1)

	var sa1 secretAgent

	sa1 = secretAgent {
		admi: person{
			first: "Utkarsh",
			last: "Singh",
			age: 21,
		},

		ltk: true, 
	}

	Sa1 := struct_test.Secretagent{
		Man : struct_test.Man{
			First_name: "Utkarsh",
			Last_name: "Singh",
			CodeName: 007,
		},

		Ltk: true,
	}

	fmt.Printf("%T\n",Sa1.Man)
	fmt.Printf("%T\n",sa1.admi)



	// Anonymous struct
	var a1 struct{
		emp_id int
		yoj int
		salary float64
	}

	fmt.Printf("%T\n",a1)

	// assign a value to a variable or identifier id1 = type{value}
	a1 = struct{
		emp_id int
		yoj int
		salary float64
	}{
		emp_id: 2345,
		yoj: 2014,
		salary: 540089.902,
	}

}
