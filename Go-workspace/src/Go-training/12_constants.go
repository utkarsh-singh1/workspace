package main

import "fmt"


// to define a constant ientifier 
// u can write this in variable also to define multiple variable in one go
//find about constant declaration here =>> https://go.dev/ref/spec#Constant_declarations

const (
    a = 23
    b = 23.34
    c = "23 haha"
)

var (
    x = 32
    y = 32.43
    z = "32 nana"
)

func main(){
    fmt.Println(a,b,c)
    fmt.Printf("%T\t%T\t%T\n",a,b,c)
    fmt.Printf("%T\t%T\t%T\n",x,y,z)
}
