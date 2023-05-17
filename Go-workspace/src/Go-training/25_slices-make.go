package main 

import "fmt"

func main() {
    
    // make([]int, len, cap) -- (len/length, cap/capacity is underline array) 
    // Ex. x := make([]int, len, cap)

    a := make([]int, 10, 12)
    
    fmt.Println(len(a))
    fmt.Println(cap(a))

    fmt.Println(a)


    a[0] = 34
    a[9] = 78
    
    fmt.Println(a[0], a[9])

    fmt.Println(a)
    fmt.Println(len(a))
    fmt.Println(cap(a))

    a = append(a, 789)

    fmt.Println(a)

    fmt.Println(a)
    fmt.Println(len(a))
    fmt.Println(cap(a))

    a = append(a , 89, 76)

    fmt.Println(a)
    fmt.Println(len(a))
    fmt.Println(cap(a))


}
