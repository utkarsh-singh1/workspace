package main

import "fmt"

func main() {

	Print(4,"you")
}

func Print(x int,y string) {			
						
	fmt.Printf("%s is %#v years old \n",y,x)
						
}						
