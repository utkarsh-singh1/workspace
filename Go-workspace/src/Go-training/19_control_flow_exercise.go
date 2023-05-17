package main

import "fmt"

func main() {
  
   x := 1
   fmt.Printf("%v\n",x) 

    for i:=0 ; i<100; i++{
        if i % 2 == 0 {
            fmt.Println(i)
        }
    } 

}
