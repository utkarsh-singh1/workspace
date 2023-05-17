package main

import "fmt"

// see the pattern -->> (keyword identifier type)

type person struct {
    first string ; last string
}

type secretAgent struct {
    person  ; ltk bool
}

// func (r reciever) identifier(parameter(s)) (return(s))

func (s secretAgent) playFootball() {
    fmt.Println(s.first, s.last,"is playing football and has",s.ltk)
}

func (p person) playFootball() {
    fmt.Println("Football is playing by", p.first,p.last)
}

// Interface helps to define behaviour and allow us to use polymorphism
// more here -->> https://go.dev/ref/spec#Interface_types
// Interface says -- if you have this method , then you are also my type , if this interface does not have any method, then this become "empty interface" and thats what it is all about , every other type is also of type empty interface.
type human interface {
   playFootball() 
}

func bar(h human) {

    // assertion
    switch h.(type) {

    case person:
        fmt.Println("It is of type",h.(person).first)
    case secretAgent:
        fmt.Println("It is of type",h.(secretAgent).first)
    }
    
    fmt.Println("Haha you punny",h)
}

func main() {
    p1 := person{
        first : "Utkarsh",
        last : "Singh",
    }

    sa1 := secretAgent {
        person: person{
            first : "Utkarsh",
            last: "singh",
        },
        ltk: false,
    }

    sa2 := secretAgent{
        person: person{
            first: "Somu",
            last: "shekh",
        },
        ltk: true,
    }
     
    // Now value can be more than one type
    p1.playFootball()

    fmt.Printf("%T\t%T\t%T\n",p1,sa1,sa2)
    

    // this also known as polymorphism - as bar a function can take multiple different types as args - sa1,sa2,p1(All of these are of type {human}. Also it can come handy if we want to do more with bar() by assigning different types. This all become possible due to interface. Any value of type secretAgent and person is also of type human - due to polymorphism.
    bar(sa1)
    bar(sa2)
    bar(p1)
}
