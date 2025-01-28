package main

import "fmt"

type person struct {
	FirstN string
	LastN  string
}

func (p *person) speak() string {

	return p.FirstN

}

type human interface {
	speak() string
}

func saySomething(h human) string {

	return fmt.Sprintf("Please say something Mr. %s", h.speak())

}

func main() {

	p1 := person{
		"Hello", "World",
	}

	p2 := &person{
		"Bolo",
		"World",
	}

	fmt.Println(p1)
	fmt.Println(saySomething(p2))
}
