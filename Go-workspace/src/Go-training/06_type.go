package main

import "fmt"

var a int = 42 


// just created an type called as hotdog which have underline type int
type hotdog int


// created a variable called b which have underline type hotdog
var b hotdog

func main(){
       
      b = 43

      fmt.Println(a)
      fmt.Printf("%T\n", a)
        
      // but remember a & b both are diffrent type of variable  so a != b

      // but there is a method where we can convert one type into another more about conversions here - https://go.dev/doc/effective_go#conversions
      a = int(b)

      fmt.Println(a)
      fmt.Printf("%T\n", a)

}
