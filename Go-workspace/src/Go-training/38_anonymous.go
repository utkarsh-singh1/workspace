package main

import "fmt"

func main() {
   
    // Anonymous -->> without any identification.

    foo()

    // Creating a function and calling a function at the same time.

    // an Anonymous function without any parameter
    func() {
        fmt.Println("Anonymous")
    }()

    // an Anonymous function with an parameter
    func(s string) {
        fmt.Println("Run like you are",s)
    }("Lion")
    
    // an Anonymous function return an value
    func() string{
        return "hello hi, bye"
    }()


}

func foo () {
    fmt.Println("Foo hello")
}
