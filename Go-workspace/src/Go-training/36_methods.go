package main

import "fmt"

type person struct {
    first string
    last string
}

type secretAgent struct {
    person
    ltk bool
}

// func (r reciever) identifier(parameters) (return(S)) { code }
// Now r reciever -->> s secretAgent , means now our func is now attached to type secretAgent , Now our func has become method and now any value of secretAgent can access this speak() function.


func (s secretAgent) speak() {
    fmt.Println("I am",s.first,s.last)
}

func main() {
    sa1 := secretAgent{
        person : person{
            first: "Utkarsh",
            last: "Singh",
        },
        ltk: true,
    } 

    fmt.Println(sa1)

    sa1.speak()
}
