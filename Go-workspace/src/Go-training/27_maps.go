package main

import "fmt"

func main() {
   

    // MAPS

    // to learn more about -- maps ==>> https://go.dev/doc/effective_go#maps
    // m := type{value} -- composite literal
    // m := map[type of key]type of value{}
    m := map[string]int{
    
        "James" : 32,
        "Miss Moneypenny" : 27,


    } //-- map[string]int is type

    
    fmt.Println(m)

    fmt.Println(m["James"])

    fmt.Println(m["Bernabus"])

    // Now look here "Bernabus" key does not exist in the map but it prints value 0 which is not resonable but there is a method where we can check that.
    // declare v, ok -- ok takes stores the value return by m[key] which is true/false means the keys exist or not and v stores key value.

    
    v , ok := m["Bernabus"]

    fmt.Println(v)
    fmt.Println(ok)

    // Next method to check

    if v, ok := m["Bernabus"] ; ok {
        fmt.Println("If ok prints", v)
    }

    if v, ok := m["James"] ; ok {
        fmt.Println("If ok prints", v)
    }

    if v, ok := m["Miss Moneypenny"] ; ok {
        fmt.Println("If ok prints", v)
    }


    // add elements in maps

    m["utkarsh"] = 25

    for a,b := range m {
        fmt.Println(a,b)
    }

    
    m["maa"] = 52

    fmt.Println(m)
    
    for a,b := range m {
        fmt.Println(a,b)
    }


    // using range and for

    /*for k,x := range m {
        fmt.Println(k,x)
    }*/

    // delete in maps
    // syntax ==>> delete(<map>,"key")
    
    delete(m, "utkarsh")

    fmt.Println(m)
    

    // What if key does not exist --- then nothing.
    delete(m, "Ian")

    fmt.Println(m)

    // to delete with check if key exist or not...

    if c, ok := m["James"] ; ok {
        fmt.Println("value:", c)
        delete(m, "James")
    }

    fmt.Println(m)
}
