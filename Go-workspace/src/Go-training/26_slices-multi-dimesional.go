package main

import "fmt"

func main() {
    x := []int{1,2,3,4,5}

    fmt.Println(x)

    y := []int{6,7,8,9,10}

    fmt.Println(y)

    z := [][]int{x,y} // -- slice of slice of int

    fmt.Println(z)

    xi := []int{2,3,5,7,9}

    for i ,v := range xi{
        fmt.Println(i,v)
    }

    
}
