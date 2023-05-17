package main

import "fmt"

// no value assigned so compiler assigned false
var x bool

func main(){
    fmt.Println(x)
    x = true

    fmt.Println(x)


    a := 42
    b := 7
    fmt.Println(a == b, a<=b, a>=b)
}
