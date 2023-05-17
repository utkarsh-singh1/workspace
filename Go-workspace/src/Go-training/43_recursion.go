package main

import "fmt"

func main() {
    fmt.Println(4*3*2*1) 

    // find the factorial using recursion
    a := factorial(4)

    fmt.Println(a)

    // find the factorial using loop

    var n int
    
    fmt.Print("Give me a value :< ")
    fmt.Scanln(&n)

    fact := 1

    for i := n ; i > 0 ; i -- {
        fact *= i
    }
    
    fmt.Println(fact)

    y := loopFact(4)

    fmt.Println(y)
}

func factorial(n int) int {
    
    if n == 0 {
        return 1
    }

    return n * factorial(n-1)
}

func loopFact(n int) int {
    total := 1
    
    for ; n > 0 ; n -- {
        total *= n
    } 

    return total
}
