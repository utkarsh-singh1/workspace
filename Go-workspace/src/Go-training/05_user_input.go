package main

import "fmt"

func main() {
   
    // Take input from user , newline is seperation between different user inputs , taking input after space is not allowed
    // Ex. input -->> (go programming) -- we will only get go not programming. 
    var x , y int
    fmt.Println("Please provide some input")
    fmt.Scan(&x,&y)
    
    fmt.Println(x,y)

    // Take input from user , space is seperation between differnt user inputs, newline will end the user input prompt.
    // Ex. input -->> (go programming) -- will take it as 'go' and 'programming' are 2 different inputs.
    var a , b int
    fmt.Println("Give me some more values")
    fmt.Scanln( &a,  &b)
    
    fmt.Println(a,b)


    // Similiar to Scanln() and Scanf() take inputs but with the help of formatters.

    var (i string ; j int)
    
    fmt.Scanf("Plaese provide string value first and integer next - %s  ,%d ", &i , &j)

    fmt.Println(i,j)
}
