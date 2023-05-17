package main

import "fmt"

const  (
    _ = iota
    kb = 1 << (iota * 10) 
    mb = 1 << (iota * 10)
    gb = 1 << (iota * 10)
)


func main(){
    
    x := 4

    fmt.Printf("%d\t\t%b\n",x,x)
    

    // Bit shifting
    y := x << 1
    
    fmt.Printf("%d\t\t%b\n",y,y)

    z := x >> 1

    fmt.Printf("%d\t\t%b\n",z,z)


    // kb mb gb
    // https://go.dev/ref/spec#Iota
    /*
    kb := 1024
    mb := 1024 * kb
    gb := 1024 * mb
    */

    // Learn about bit hacking with go - https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827


    fmt.Printf("%d\t\t\t%b\n", kb,kb)
    fmt.Printf("%d\t\t\t%b\n", mb,mb)
    fmt.Printf("%d\t\t%b\n", gb,gb)
}
