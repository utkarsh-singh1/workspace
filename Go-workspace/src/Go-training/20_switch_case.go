package main

import "fmt"

func main(){
    

    // find more about switch cases =>> https://go.dev/ref/spec#Switch_statements
    // Also this - https://go.dev/doc/effective_go#switch
   
    // One way of doing it
    switch {
    case false:
        fmt.Println("hello world")
    case (2==3):
        fmt.Println("one more hello world")
    case (3==3):
        fmt.Println("one on one hello world")
        fallthrough // as there is no default fallthrough in go. Means if this is true, fall down
    case (4==4):
        fmt.Println("two more hello world")
        fallthrough
    case (4==6):
        fmt.Println("three more hello world")
        fallthrough
    case (5==4):
        fmt.Println("four more hello world")
        fallthrough
    case (5==4):
        fmt.Println("five more hello world")
        fallthrough
    default:
        fmt.Println("this is default")  // this case will run when every other case is false
    }

    // Another way of doing it

    n := "uk"

    switch n {
       
    case "Moneypenny":
        fmt.Println("Money ....")
    case "uk":
        fmt.Println("uk this is uk")
    case "q":
        fmt.Println("why in hindi is q")
    default:
        fmt.Println("no body else here is telling truth")
    }

    // Another way of doing it

    m := 1

    switch m{
    
    case 2,3,1:
        fmt.Println("Here it evaluates with multiple cases also")
    case 4:
        fmt.Println("Nope nothing here")
    default:
        fmt.Println("Alas you came here")
    }

}
