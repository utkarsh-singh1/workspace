
package main

import "fmt"

func main(){

    // x := type{values} ==>> also known as composite literal

    x := []int{4,5,6,8,9,35,56} // slice of int
    fmt.Println(x)

    // Slice for range

    // method 1 for printing all values of slice x
    fmt.Println(len(x))
    fmt.Println(x[0])
    fmt.Println(x[1])
    fmt.Println(x[2])
    fmt.Println(x[3])
    fmt.Println(x[4])

    // method 2 for printing all values of slice x

    for i,v := range x{
        fmt.Printf("value of x[%v] is %v\n",i,v)
    }

    for j:=0 ; j < len(x) ; j++{
        fmt.Println(x[j])
    }

    // slice of slice

    a := []int{4,3,6,2,78}

    fmt.Println(a[4])
    fmt.Println(a)
    fmt.Println(a[:])
    fmt.Println(a[1:3])
    fmt.Println(a[:4])

    // Append on slice

    // https://go.dev/doc/effective_go#arrays
    // syntax for append - func append(slice []T, elements ...T) []T ==>> ....T means any number of values of type T , .... reperesents variadic perameter means any number of values.
    // Example -

    a = append(a, 67,78,90,98)

    b := []int{3,5,7,9,10,11}

    fmt.Println(a)
    fmt.Println(b)

    // also take a look at the use of variadic parameter

    a = append(a, b...)  // b... means all values inside b

    fmt.Println(a)

    // c := []string{"go", "python", "java", "js", "c", "cpp"}
    
    // Delete in slice

    // Now value of a is =  [4 3 6 2 78 67 78 90 98 3 5 7 9 10 11]
  
    // if want remove elements in slice like a[4] = 78 and a[5] = 67

    a = append(a[:4],a[6:]...)

    fmt.Println(a)
}


// a SLICE allows you to group toghether values of same type
