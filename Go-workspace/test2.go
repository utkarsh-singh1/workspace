package main 

import "fmt"

func main() {

    arr := [5]int{769082435,210437958,673982045,375809214,380564127}

    var (
        sum int;
        min int;
        max int;
    )
    
    min = arr[0]
    
    for _,v := range arr {
        sum += v
    }
    
    for _,v1 := range arr {
        if v1 > max {
            max = v1
        }
    }
    
    for _,v2 := range arr {
        if v2 < min {
            min = v2
        }
    }

    fmt.Println(sum)
    fmt.Println(min)
    fmt.Println(max)


    fmt.Println(sum-max,sum-min)
}
