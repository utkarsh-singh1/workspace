package main

import "fmt"

func main(){
	// There is no while / do-while keyword
	// how to write loop in golang
	/*
	for init ; condition ; post {
		fmt.Println("Hello World")
	}
	*/
	for i := 0 ; i <=100; i++{
		fmt.Println(i,"and Hello World")
	}

	// nesting loop
	for i := 0 ; i <=10; i++{
		for j := 0 ; j < 3; j++{
			fmt.Printf("sum of %d and %d is %d\n", i,j,i+j)
		}
		
	}

    // for acts like while
    x := 1                // init

    for x < 10{           // condition
        fmt.Println(x)    
        x++               // post
    }

    fmt.Println("done")

    // for acts as infinite loop
    var a int = 1

    for {
        if a > 9{
            break
        }
        fmt.Println(a)
        a++
    }

    fmt.Println("done again")


    // break and continue

    b := 1

    for {
        if b > 100{
            break
        } 
        
        if b%2 != 0{
            b++
            continue
        } 

        fmt.Println(b)
        b++
    }

}
