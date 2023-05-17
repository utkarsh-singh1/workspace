package main

import "fmt"

// package level scope
//var x int

func main() {
    var x int
    // CLOSURE - to enclose a certain variable scope inside limited area.
    fmt.Println(x)
    x++
    
    {
        y := 43
        // scope of y is here inside this code block
        y++
        fmt.Println(y)
    }
    
    // you cannot run anything related to y here example - 
    // fmt.Println(y) // <<-- uncomment this line and run it.
    fmt.Println(x)
    x--
    fmt.Println(x)

    a := increamentor()

    fmt.Println(a())
    fmt.Println(a())
    fmt.Println(a())
    fmt.Println(a())
    fmt.Println(a())
    
    b := increamentor()

    fmt.Println(b())
    fmt.Println(b())
    fmt.Println(b())
    fmt.Println(b())
    
   {
        fmt.Println(increamentor()())
        fmt.Println(increamentor()())
        fmt.Println(increamentor()())
        fmt.Println(increamentor()())
        fmt.Println(increamentor()())
        fmt.Println(increamentor()())
    }
}

func increamentor() func() int{
     var z int
    return func() int{
        z++ ; return z 
    }
}
