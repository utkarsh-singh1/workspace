package main

import "fmt"

func main() {

    // When to use pointers

    /* 

    1. When you have a huge chunk of data and you don't want to pass it around, you create a memory location and store the data there and you have address of that location , now you can access your data by only using address of that memory location where that data is stored.

    2. You want to change the data at a location, a certain chaneg in value at that location , get the address , derefrence it and assign new value to it.

    */

    x := 0

    fmt.Println("Before foo x is",x)

    foo(x)

    fmt.Println("After foo x is", x)


    fmt.Printf("\n")


    fmt.Println("Before bar x is",x)
    fmt.Println("Before bar",&x)

    bar(&x)

    fmt.Println("After bar x is",x)
    fmt.Println("After bar",&x)
}


/* o/p -->> 
0
32
0

-- but why -->> because x is assigned to y so change gonna be happening to y not x
*/
func foo(y int) {
    fmt.Println(y)

    y = 32

    fmt.Println(y)
}

func bar(a *int) {
    fmt.Println("Before change", a)
    fmt.Println("before change", *a)

    *a = 43

    fmt.Println("After change", a)
    fmt.Println("After change", *a)
}
