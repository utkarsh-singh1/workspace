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

func main() {
	// problem 1

	//var x rune = 45

	for i := 0; i <= 10000; i++ {
		fmt.Println(i)
	}

	// problem 2
    k := 65
	for j := 1; j <= 26; j++ {
		fmt.Println(j)
		
		fmt.Printf("\t%#U\n",k)
        fmt.Printf("\t%#U\n",k)
        fmt.Printf("\t%#U\n",k)
        k++
    }

    // or solution for problem 2

    for l := 65 ; l <=90; l++{
        fmt.Println(l)
        for m := 0 ; m <3 ; m++{
            fmt.Printf("\t%#U\n",l)
        }
    }

    // problem 3
    
    
    a := 1999
    b := 0

    for  a <= 2023 {
        b++
        a++
    }

    fmt.Println(b)
    
    // other sol for problem 3
    
    c := 1995

    for c <= 2023 {

        fmt.Println(c)
        c++
    }

    // problem 4

    d := 1992

    for {

        if d > 2023 {
            break
        }
        fmt.Println(d)
        d++

    }

    // problem 5

    for e := 10 ; e<=100 ; e++{
        
        fmt.Println( e % 4)
    }

    // problem 6

    n := "good"

    if n != "good" {
        fmt.Println("Don't be like that")
    }else if n == "good"{
        fmt.Println("You are the man")
    }else{
        fmt.Println("Haha got you")
    }
    
    // problem 7

    // add elseif and else in above problem 6 code


    // problem 8

    switch {
    case 23==23:
        fmt.Println("Yes here it is ")
        fallthrough
    case false:
        fmt.Println("No there is none")
    case true:
        fmt.Println("Haha ye hoo")

    }

    // problem 9
    
    favSport := "Hockey"

    switch favSport {

        case "Football":
            fmt.Println("Not my type")

        case "Cricket":
            fmt.Println("Okayish")
        case "Hockey":
            fmt.Println("Thats my bro")
    }

    // problem 10

    fmt.Println(true && true)
    fmt.Println(true && false)
    fmt.Println(true || true)
    fmt.Println(true || false)
    fmt.Println(!true)


    

}
