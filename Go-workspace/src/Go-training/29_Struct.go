package main

import "fmt"

// 1. Thats how to create a type using struct -- (by composing values of diffrent type)

type person struct {
    first string
    last string
    age int
}

// 2. embedded struct ==>> here embedding type person in another type secretAgent
// person is the inner type and secretAgent outer type


type secretAgent struct {
    person
    ltk bool
}


// 3. Anonymus Struct



func main() {
       
    
    // Find more about structs ==>> https://go.dev/ref/spec#Struct_types
    // struct ==>> It is a data structure that compose values of other types.

    sa1 := secretAgent{
        person : person{
            first: "Utkarsh",
            last: "Singh",
            age: 24,
        },
        ltk : true,
    }
    
    p1 := person{
        first: "James",
        last: "Bond",
    }

    p2 := person{
        first: "Miss",
        last: "Moneypenny",
    }
    fmt.Println(sa1)
    fmt.Println(p1)
    fmt.Println(p2)
   
    fmt.Println(sa1.first, sa1.last, sa1.ltk, sa1.age) //-->> Here first from type person(inner) promoted to upper(secretAgent) type. 


    // both are similiar ^^ (sa1.first = sa1.person.first)
    fmt.Println(sa1.person.first)
    /*
    Above syntax also can be used when
    type secretAgent struct{
        person
        ltk bool
        first
    }

    as person has also value first and secretAgent also
    */

    fmt.Println(p1.first)
    fmt.Println(p2.last)

    // 3. Anonymus struct -- because it does not have any identifier as compared to others like person, secretAgent, syntax is same except u define it like that -- why do you need it 
    // because to avoid code pollution. 
    
    p3 := struct{
        first string
        last string
        age int
    }{
        first: "James",
        last: "Bond",
        age: 25,
    }
 
    fmt.Println(p3)  
}
