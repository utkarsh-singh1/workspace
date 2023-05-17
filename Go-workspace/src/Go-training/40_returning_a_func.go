package main

import "fmt"

func main() {

    s1 := foo()

    fmt.Println(s1)

    x := bar(6)

    /* 
    x gets the value -->> func() int{
                        return 451
                        }
    thats why x is of type func() int
    */
    i := x()

    // Now here i gets the return value 451 thats why i is of type int.
    fmt.Println(i)

    fmt.Println(x) // or fmt.Println(x)  // Prints a hex value - probably address of the func bar()

    // or 
    fmt.Println(bar(3)())

    fmt.Printf("%T\n",x)
    fmt.Printf("%T\n",i)

    fmt.Printf("%T\n",bar)
}

func foo() string{
    s := "hello world"

    return s
}

// func() int is also become a type after it is been returned.
func bar(x int) func() int{
    return func() int{
        return 43
    }
}
