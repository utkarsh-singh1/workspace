package main

import (
	"fmt"
	"time"
)

func main() {
    
    var contributionNum int 
        
    var targetNum int

    var alreadydoneNum int
    
    fmt.Println("*************************************************************************")
    fmt.Println()
    fmt.Println("..........Welcome to devCount - count your daily contribution.............")
    fmt.Println()
    fmt.Println("**************************************************************************")
    fmt.Println()
    
    fmt.Print("What is your Target -> ")
    fmt.Scan(&targetNum)

    fmt.Print("How much is already done -> ")
    fmt.Scan(&alreadydoneNum)
    
    fmt.Print("How much you have done at ",time.Now(), " -> ") 
    fmt.Scan(&contributionNum)

    fmt.Println("You have", alreadydoneNum+contributionNum, "contribution tille this moment.")
    

    if targetNum - (alreadydoneNum + contributionNum) == 0 {
    
            fmt.Println("Hooray you completed your Target")

    }else{

        fmt.Println("You need another",targetNum - (alreadydoneNum + contributionNum),"No. of contributions" )
    }
}
