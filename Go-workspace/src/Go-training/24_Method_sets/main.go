package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println(d.first)
}

func (d *dog) run() {

	d.first = "Rover"
	fmt.Println(d.first)
}

type youngin interface {
	walk()
	run()
}

func hyper(y youngin) {
	y.run()
}

func main() {

	d1 := dog{"henry"}

	d2 := &dog{"janet"}

	d1.walk()
	d1.run()

	// hyper(d1.first)

	d2.walk()
	d2.run()

	hyper(d2)
}
