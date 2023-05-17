package main

import "fmt"

func main() {
    
    // defer also known as - put off (an action or event) to a later time; postpone.
    // More here  =>> https://go.dev/ref/spec#Defer_statements
    // from the docs - 
/* A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is panicking.
*/

     
    x := ifagain(1)
    defer fmt.Println(x)
    bar()
    defer foo()

}

func foo() {
    fmt.Println("foo")
}

func bar() {
    fmt.Println("bar")
}

func ifagain(i int) string {
    if i == i+1 {
        fmt.Println("Haha you got that damn right")
    }   

    return "Aww, go do it again!"
}
