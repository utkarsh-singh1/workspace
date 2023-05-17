package main

import "fmt"

var y string
var z int

func main(){
    
    fmt.Println("print this thing", y , "ended")
    fmt.Printf("%T\n", y)

    y = "This is uk, just uk"

    fmt.Println("Just say it man", y)
    fmt.Printf("%T\n", y)

    fmt.Println("print this also", z)
    fmt.Printf("%T\n", y)
}
