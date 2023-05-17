package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

const (
    d = iota 
    e
    f
)


const (
    r = iota +1
    q= iota +2
    i = iota + 3
)

const (
    j = iota
    k
    l
    m = iota
    n
    o
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

    fmt.Println(r,q,i)
    fmt.Println(j,k,l,m,n,o)
    fmt.Println(d,e,f)

}


// to find about iota more - https://go.dev/ref/spec#Iota
