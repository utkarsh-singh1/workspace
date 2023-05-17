 
/*
 _   _ _        _       
| \ | (_)_ __  (_) __ _ 
|  \| | | '_ \ | |/ _` |
| |\  | | | | || | (_| |
|_| \_|_|_| |_|/ |\__,_|
             |__/       

*/

package main

import "fmt"

// Problem 1
var i int = 68

func main(){
    
    // Problem 1
    fmt.Printf("%d\t%b\t%#x\n",i,i,i)

    // Problem 2
    x := (1==1)
    y := (2<=4)
    z := (3>=5)
    t := (6!=8)
    w := (9>0)
    q := (0<0)

    fmt.Println(x,y,z,t,w,q)

    // Problem 3

    const Y int = 6
    const Z = "Uk"

    fmt.Println(Y,Z)

    // problem 4

    var a int = 1

    fmt.Printf("%d\n",a)

    b := a << 1

    fmt.Printf("%d\t%b\t%#x\n",b,b,b)

   // Problem 5

   var c string = `"Hello from here" as they say from there`

   fmt.Println(c)

   // problem 6

   const (
       g = 2023 + iota
       f = 2023 + iota
       e = 2023 + iota
       d = 2023 + iota
   )
    fmt.Println(d,e,f,g)

}
