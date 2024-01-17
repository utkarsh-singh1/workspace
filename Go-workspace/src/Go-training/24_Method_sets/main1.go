package main

import (
	"fmt"
	"math"
)

type circle struct{

	radius float64
	
}

func (c *circle) area() float64 {

	return math.Pi*c.radius*c.radius
}

type shape interface {

	area() float64
}

func info(s shape) {

	fmt.Println(s.area())
}


func main() {

	c1 := circle{
		radius: 4,
	}

	c2 := &circle{
		radius: 5,
	}

	fmt.Println(c1.area())
	
	// info(c1)

	info(c2)
}

