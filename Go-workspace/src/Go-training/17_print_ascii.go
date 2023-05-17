package main

import "fmt"

func main(){
    
    // to learn more about go format specifier - https://pkg.go.dev/fmt

    s := "Hello World"
    
    a := 100

    fmt.Println(s)
    
    for i:=33 ; i<=122 ; i++{
        fmt.Printf("%v\t%#x\t%#U\n",i,i,i)
    }

    fmt.Printf("%#U\n", a)

    fmt.Println(s)

}
