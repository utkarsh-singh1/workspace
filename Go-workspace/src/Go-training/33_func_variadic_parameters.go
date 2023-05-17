package main

import "fmt"

func main() {
    fmt.Println("Hello World")

    x := sum(2,3,0,4,5,5,6,8,7,9)
    
    fmt.Println("Total of all values is ",x)
    
}

// func (r reciever) identifier(parameter(s) type) (return(s)) { code }

// Thats how to define variadic parameter (...int) -> unlimited number of ints
func sum(x ...int) int {
    fmt.Println(x)
    fmt.Printf("%T\n",x)

    sum := 0

    for ind,val := range x{
        sum += val
        fmt.Printf("For index i %v we are adding %v to the total and now value is %v\n",ind,val,sum)
        
    }

    return sum
}
