
/*
_   _ _        _                 _____ 
| \ | (_)_ __  (_) __ _          |___  |
|  \| | | '_ \ | |/ _` |  _____     / / 
| |\  | | | | || | (_| | |_____|   / /  
|_| \_|_|_| |_|/ |\__,_|          /_/   
             |__/                    
*/

package main

import "fmt"

type person struct {
    name string 
}


func main() {
   
    // Problem 1

    var x int = 5
    
    fmt.Println(&x)

    // Problem 2

    p1 := person{
        name : "Vishal MegaMart",  
    }
    
    fmt.Println(p1.name)
    changeme(&p1)
    fmt.Println(p1.name)

}

func changeme(p *person){
    p.name = "Utkarsh Singh" // Similiar to (*p).name ="Utkarsh Singh" -- >>becuase before changing any value to a pointer we have to derefrence it first.
}
