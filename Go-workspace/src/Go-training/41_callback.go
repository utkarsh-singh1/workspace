package main

import "fmt"

func main() {
    // function can be returned, can be assigned to variable and Now
    // callback - passing function as an argument

    ii := []int{2,3,4,5,6,7,8,9,}

    s := sum(ii...)

    fmt.Println(s)
    
    y := []int{3,4,6,7,8,2,0,12,34,65,78} 

    fmt.Printf("%T\n",sum)
    t := evenSum(sum ,y...)

    o := oddSum(sum ,y...)

    fmt.Println("Sum of all enteries in slice y is",sum(y...))

    fmt.Println("Sum of even numbers in slice y",t)

    fmt.Println("Sum of enteries in slice y is",o)

}

func sum(xi ...int) int{
    
    // type of xi

    fmt.Printf("%T\n",xi)

    n := 0

    for _,v := range xi{
        n += v
    }

    return n
}

func evenSum(f func(xi ...int) int, vis ...int ) int{
    yi := []int{}

    for _,v := range vis{

        if v %2 == 0 {
            yi = append(yi, v)
        }
    }  

    return f(yi...)
}

func oddSum(f func(xi ...int) int, vio ...int) int {

    oi := []int{}

    for _,v := range vio {
        if v %2 != 0 {
            oi = append(oi,v)
        }
    }

    return f(oi...)

}
