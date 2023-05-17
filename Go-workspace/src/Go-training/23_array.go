package main

import "fmt"

func main() {
    // https://go.dev/doc/effective_go#arrays
    // https://go.dev/ref/spec#Array_types

    // zero value array 
    var x [5]int
    
    var y [8]int
    // as for any identifier type is what value assign to it like ==>> var x = 4 or var x int.
    // in array size is part of its type means variable x and y are both of different type.
    
    // assigning value in array
    x[3] = 42
    y[7] = 78
    fmt.Println(x, len(x), y)
    
    var z [5]int 
}
