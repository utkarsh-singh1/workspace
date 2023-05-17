// Apart from this course to learn about go easily follow this tutorial - https://gobyexample.com/


package main

import "fmt"

var j string = "hello world"

func main() {
	fmt.Println("Helllo World")

	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i+3)
		}
	}
	var i int

	fmt.Scanln(i)

	if i >= 1 {
		Hello()
	}

	bar()

	zoo()
}

func Hello() {
	fmt.Println("I , I Like it!")
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("Here to rock")
}

func zoo() {
	fmt.Println("line in zoo and bear doing boo")
}

func row() {
	fmt.Println("row again again")
}

/// (2) loop/ iterative
// (3) conditional
