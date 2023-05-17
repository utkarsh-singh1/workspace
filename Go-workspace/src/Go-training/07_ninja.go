 
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

// Problem 5
type hitrun int
var p hitrun
var q int


// Problem 4
type goo string

// Problem 2
var (a int ; 
b string ; 
c bool) 

// Problem 3
var (d int = 42; 
e string = "James Bond" ; 
f bool = true) 

func main(){
          
     

     // Problem 1
     x := 42
     y := "James Bond"
     z := true

     fmt.Println(x,y,z)

     fmt.Println(x)
     fmt.Println(y)
     fmt.Println(z)

     // Problem 2

     fmt.Println(a)
     fmt.Println(b)
     fmt.Println(c)

     // Problem 3

     s := fmt.Sprintf("%v\t%v\t%v\n",d,e,f)

     fmt.Println(s)

     // Problem 4

     type foo int

     var i foo

     fmt.Println(i)

     i = 42

     fmt.Println(float32(i))

     fmt.Printf("%T\n",float32(i))

     hello()

     // Problem 5
     fmt.Println(p)
     fmt.Println(q)
     
     q = int(p)

     fmt.Println(q)

     fmt.Printf("%T\n",q)
      
}

func hello() {
    var j goo

    j = "90"

    fmt.Println(j)
}
