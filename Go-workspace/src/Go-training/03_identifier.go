package main

import "fmt"


var y = 42

// DECLARE THE VARIABLE with the identifier is of type string.
var z string = "Shaken , not stirred"

var a string = `Bhaiya said, "Jai shree ram"`

func main(){
  
    fmt.Println(y)
    
    // Use %T to know the type of the identifier
    fmt.Printf("%T\n", y)

    fmt.Println(z)
    fmt.Printf("%T\n", z)

    // since go is a statically typed language like js / python 
    // we cannot do that ==>> z = 12 as z is type string.

    fmt.Println(a)
} 
