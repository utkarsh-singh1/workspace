package main

import "fmt"

// We create values of certain type which stored inside variables which have identifiers.

var x int = 56

type person struct {
    first string
    last string
}  


type foor int
var y foor 

func main() {
    
    p1:= person{
        first : "Utkasrh",
        last: "Singh",
    }

    fmt.Println(p1)
}
