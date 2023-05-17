package main

import "fmt"

func main() {
    // Assigning anonymous function to a variable.

    f := func(){
        fmt.Println("Hello my dear friends")
    }
    
    // calling the function 
    f()

    g := func(x int){
        fmt.Println("Haha very *",x,"funny")
    }

    g(2)

    k := func(y bool) (int, int) {
        if y{
            return 0,1
        }

        return 1,0
    }

    k(true)

    fmt.Println(k(true))


}
