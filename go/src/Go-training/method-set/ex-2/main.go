package main

import "fmt"

type man struct {
	first string
	last  string
}

type user struct {
	man
	age int
}

func (m man) hello() {

	fmt.Println("Hello sir")
}

func (u user) hello() {

	fmt.Println("Hello user")
}

func main() {

}
