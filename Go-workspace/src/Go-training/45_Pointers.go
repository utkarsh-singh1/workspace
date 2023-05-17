/*
 ____       _       _                
|  _ \ ___ (_)_ __ | |_ ___ _ __ ___ 
| |_) / _ \| | '_ \| __/ _ \ '__/ __|
|  __/ (_) | | | | | ||  __/ |  \__ \
|_|   \___/|_|_| |_|\__\___|_|  |___/
                                     
*/


package main

import "fmt"

func main() {


    // Pointers - when you point to a specific location in memory where some value is stored.
    // Remember - everything in go is pass by value

    var a int

    a = 42

    // To get the value we saved inside 
    fmt.Println(a)

    // To get the address of the location / memory 
    fmt.Println(&a)  // a is stored here at that address

    fmt.Printf("%T\n",a)   // o/p => int 

    fmt.Printf("%T\n",&a)  // o/p --> *int --( that means pointer to an int) , &a shows address of a certain location in memory where a int is stored , and its type is *int means pointing to a certain location that holds value of type int.
                            

    // Now int and *int are 2 different types.

    // like -->> var b int = a // o/p -->> No problem
    // but -->> var b int = &a // o/p -->> can't dp that 2 different types

    // but 
    var b *int = &a // this is good no problem . --- also can be written as -->> b := &a.

    fmt.Println(b)

    // Now to get the value from any given address 

    fmt.Println(*b)  // This will print the value stored at this location , also can be called as derefrencing a address to get the value

    // Above can also be written as
    fmt.Println(*&a)

    // Since *b is similiar to a -->> because a also represents a value it holds thats why *b also represents value it holds.
    // So like 'a' we can reassign a value also by using *b.
    
    *b  = 43

    fmt.Println(a)


    // Also * in *b and *int are different -->> in *b it is acting like operator to get the value from a address which has a type *int(here * is part of type symbolising pointing at acertain location) 


}
                      
