package main

import "fmt"


// Now variable z is defined here scope of "z" is package level
var z = 5

// can define multiple variables of same type in a single line.
var p,q int


// unassigned value variables in go assign themseleves value 0
// These value also called as Zero value variables.
/*false for booleans, 0 for numeric types, "" for strings, and 
nil for pointers, functions, interfaces, slices, channels, and maps.*/
// You can find about it more here - https://go.dev/ref/spec#The_zero_value

var d int

func main(){
   
    // Most used or orginal way of assigning values to variable

    var a int = 3
    
    var b = "fleet"

    // short declaration variable can be used as assigning values to variables . :)
    // short declaration variable can not be used outside function body. 
    
    x := 1
    
    // also scope of x is here beacuse before x := 1 it does not have any personality
    fmt.Println(x)
    
    x = 34
    
    fmt.Println(x)
    
    y := 2+ 3
    
    fmt.Println(y)

    fmt.Println(b)

    fmt.Println(a)

    p = 34
    q = 90

    // or p,q = 34,90

    fmt.Println(p,q)

    // swap values like python :)
    p,q = q,p

    r,s := 56,998

    fmt.Println(r,s)

    p := 90

    p , q :=  34, 78

    fmt.Println(p,q)
}
