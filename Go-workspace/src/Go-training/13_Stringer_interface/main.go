package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
	
}

func (b book) String () string {
	return fmt.Sprintln("The title of this book is",b.title)
}

type count int

func (c count) String () string {
	return fmt.Sprintln("The number is ",strconv.Itoa(int(c)))
}


// Since Stringer is interface with String() method underhood , Any value of any type (underline type string,int,struct whatsever...) have method String, will also of type Stringer due to polymorphism property comes with interface

func logInfo(s fmt.Stringer) {

	log.Println("The log from file 13 is:-", s.String())
	
}

func main() {

	b1 := book{
		title: "Arthashastra",
	}

	var c count = 42


	fmt.Println(b1)
	fmt.Println(c)

	// Use log to get log of the current code

	log.Println(b1)
	log.Println(c)

	// If i wanted to add a additional text in front of output from log.Println(b1) and log.Println(c) i will use wrapper function

	// A wrapper function is a fuction that is created to add a additonal layer of abstraction or functionality over an existing method or function. It acts as an intermediatery between a caller and wrapped function, and provide additional or modified inputs, output or behaviour without even changing original function.


	// Here logInfo is a wrapper function as it wrap over a existing method String()
	logInfo(b1)
	logInfo(c)

	fmt.Printf("%T\n",b1.String())
}
