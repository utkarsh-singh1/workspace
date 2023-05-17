package main

import "fmt"

func main() {
    xi :=  []int{2,3,4,5,6,7,8}

    x := sum("James",xi...)  // Also instead of mention a lot of parameters we can also do that like 


    y := sum("Hello")  // To know more about varaiadic parameters ==>> https://go.dev/ref/spec#Passing_arguments_to_..._parameters

    z := sum("James", 4,5,6,7,8)

    fmt.Println(z)
    
    /*
    Now according to the upper linked document 

    --> If f is variadic with a final parameter p of type ...T, then within f the type of p is equivalent to type []T. If f is invoked with no actual ar-guments for p, the value passed to p is nil. Otherwise, the value passed is a new slice of type []T with a new underlying array whose successive elements are the actual arguments, which all must be assignable to T. The length and capacity of the slice is therefore the number of arguments bound to p and may differ for each call site.

Given the function and calls

-------------------------------------------------------------------------------------------------------------------------
|    func Greeting(prefix string, who ...string)                                                                         |
|    Greeting("nobody")                                                                                                  |
|    Greeting("hello:", "Joe", "Anna", "Eileen")                                                                         |
|   within Greeting, who will have the value nil in the first call, and []string{"Joe", "Anna", "Eileen"} in the second. |
-------------------------------------------------------------------------------------------------------------------------

If the final argument is assignable to a slice type []T and is followed by ..., it is passed unchanged as the value for a ...T parameter. In this case no new slice is created.

Given the slice s and call

----------------------------------------------------------------------------------------
|   s := []string{"James", "Jasmine"}                                                   |
|   Greeting("goodbye:", s...)                                                          |
|    within Greeting, who will have the same value as s with the same underlying array. |
----------------------------------------------------------------------------------------
 
    */

    /* Means
    if function is defined like that in below code

    func sum(x ...int,s string) int{
        for i,v := range x{
            fmt.Println(i,v)
        }
    }

If f is variadic with a final parameter p of type ...T, then within f the type of p is equivalent to type []T. If f is invoked with no actual arguments for p, the value passed to p is nil. Otherwise, the value passed is a new slice of type []T with a new underlying array whose successive elements are the actual arguments, which all must be assignable to T. The length and capacity of the slice is therefore the number of arguments bound to p and may differ for each call site.

Given the function and calls

func Greeting(prefix string, who ...string)
Greeting("nobody")
Greeting("hello:", "Joe", "Anna", "Eileen")
within Greeting, who will have the value nil in the first call, and []string{"Joe", "Anna", "Eileen"} in the second.

If the final argument is assignable to a slice type []T and is followed by ..., it is passed unchanged as the value for a ...T parameter. In this case no new slice is created.

Given the slice s and call

s := []string{"James", "Jasmine"}
Greeting("goodbye:", s...)
within Greeting, who will have the same value as s with the same underlying array.


    then it will throws error as variadic parameter can not be passed as first argument according to the doc.

    but this will go throgh

    func sum(s string, x ...int) {
        for i, v := range x {
            fmt.Println(i,v)
        }
    }

    */


    fmt.Println("Sum of all values passed here ", x)

    fmt.Println(y)
}


func sum(s string,x ...int) int {
    fmt.Println(x)
    fmt.Printf("%T\n",x)

    sum := 0

    fmt.Println(len(x))
    fmt.Println(cap(x))

    for ind,val := range x{
        sum += val
        fmt.Printf("For index i %v we are adding %v to the total and now value is %v\n",ind,val,sum)
        
    }

    return sum
}
