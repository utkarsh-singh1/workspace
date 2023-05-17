/*
 _   _ _        _            ____  
| \ | (_)_ __  (_) __ _     | ___| 
|  \| | | '_ \ | |/ _` |____|___ \ 
| |\  | | | | || | (_| |_____|__) |
|_| \_|_|_| |_|/ |\__,_|    |____/ 
             |__/                  

*/

package main

import "fmt"

// problem 1 and 2

type person struct {
    first_name string;
    last_name string;
    flavour []string ;
}

// problem 3

type vehicle struct {
    doors int;
    color string
}

type truck struct {
    vehicle;
    fourWheel bool
}

type sedan struct {
    vehicle;
    luxury bool;
}

func main() {

    // problem 1
    p1 := person{
        first_name: "Utkarsh",
        last_name: "Singh",
        flavour: []string{
            "choco",
            "vanilla",
            "butter-scotch",
        },
    }

    fmt.Println(p1)


    p2 := person{
        first_name: "Hanuman",
        last_name: "Pawan-putra",
        flavour: []string{
            "strawberry",
            "capuccino",
            "coke",
        },
    }

    for _ , v := range p1.flavour {
        fmt.Println("This is my favourite flavour",v )
    }



    // problem 2

    m := map[string]person{

        p1.last_name : p1,
        p2.last_name : p2,
    }

    for w,v := range m {
        fmt.Println(w,v)
    }

    // problem 3

    truck1 := truck {
        vehicle : vehicle{
            doors : 2,
            color : "black",
        },
        fourWheel : false,
    }

    fmt.Println(truck1)

    sedan1 := sedan {
        vehicle : vehicle{
            doors: 4,
            color: "velvet red",
        },
        luxury : true,
    }
    
    fmt.Println(sedan1)


    fmt.Println(truck1.color,truck1.doors,truck1.fourWheel)

    fmt.Println(sedan1.doors,sedan1.color,sedan1.luxury)


    // problem 5

    ano_me := struct  {

        colors []string
        watches map[string]string
        first_name string
    }{
        
        first_name: "Utkarsh",
        
        colors: []string{ "red","black","white","pink"},
        
        watches: map[string]string{
            "casio" : "g-shock",
            "titan" : "htse",
            "sonata" : "super-fibre",
        },
    }

    
    ano1 := struct {
        firstName string
        lastName string
    }{
        firstName : "Uk",
        lastName: "Singh",
    }

    fmt.Println(ano1)

    fmt.Println(ano_me.watches["casio"])

    for i1 ,v1 := range ano_me.watches{
        fmt.Println(i1,v1)
    } 


}
