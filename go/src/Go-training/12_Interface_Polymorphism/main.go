package main

import "fmt"

type foo int


type person struct {

	first string
	codename int
}

type secretAgent struct {

	person
	ltk bool
	
}

func (p person) said () {
	fmt.Println("Hi ! my name is ",p.first)
}

func (s secretAgent) said () {
	fmt.Println("Hello! My name is ",s.first)
} 


type human interface {

	said()
}


func speak( h human) {
	h.said()
}

func main() {

	p1 := person{
		first: "James Bond",
		codename: 007,
	}

	sa1 := secretAgent{
		person : person{
			first: "Jenny",
			codename: 000,
		},
		ltk: false,
	}

	fmt.Println(p1.codename)

	p1.said()

	sa1.said()

	sa1.person.said()

	fmt.Printf("%T \t %T\n",p1,sa1)

	sa1.said()

	
	speak(p1)
	speak(sa1)
	
}
