package main

import "fmt"

var y int = 42

func main(){
    
    // for general printing from standard output

    fmt.Println(y)

    // To know the type of y
    fmt.Printf("%T\n",y)

    // To know the int value of y in base 2(binary)
    fmt.Printf("%b\n", y)

    // For output in 0x(hex) 
    fmt.Printf("%#x\n", y)
    // For output in 0X(hex)
    fmt.Printf("%#X\n", y)

    // For ouput in base 16 
    // For output in 
    fmt.Printf("%x\n", y)
     
    y = 2711
    
    // use multiple format specifier at once
    fmt.Printf("%x\t%b\t%#x\n", y,y,y)


    // to know about it more find out here ==>> https://pkg.go.dev/fmt
    // also to know more about escape characters ==>> https://go.dev/ref/spec#Rune_literals


    // for printing to a string which you can assign to variable

    s := fmt.Sprintf("%x\t%b\t%#x\n", y,y,y)
    fmt.Println("print these value in string", s)

    // for printing to a file or a web  server's response(Fprint, Fprintf, Fprintln)

    // also %v to know the value of the identifier
    fmt.Printf("%v\t%v\n", y,s)

}
