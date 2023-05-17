/*
 _   _ _        _              __
| \ | (_)_ __  (_) __ _       / /_
|  \| | | '_ \ | |/ _` |_____| '_ \
| |\  | | | | || | (_| |_____| (_) |
|_| \_|_|_| |_|/ |\__,_|      \___/
             |__/
*/

package main

import (
	"fmt"
	"math"
)

// problem 4

type person struct {
    first string
    last string
    age int
}

func (p person) speak() {
    fmt.Println("Hi my name is", p.first, p.last ,"and my age is" ,p.age)
}


// Problem 5

type square struct {
    side float64
}

type circle struct {
    radius float64
}

func (s square) area() float64 {
    return s.side*s.side
}

func (c circle) area() float64 {

    return math.Pi*c.radius*c.radius
}

type shape interface {
    area() float64
}

func info ( s shape) {
    fmt.Println(s.area())
}


//-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------//

func main() {

    



    // Problem 1
    x := foo()

    y , z := bar()

    fmt.Println(x)
    fmt.Println(y)
    fmt.Println(z)
    
    fmt.Print("\n")


   


    // Problem 2

    a := []int{1,2,3,4,5}

    fmt.Println(sum_var(a...))

    fmt.Println(sum_arr(a))

    fmt.Print("\n")


    




    // Problem 3

    b := []int{2,5,6,3,4,9}


    // with defer
    
    fmt.Println("This is with defer")
    show_defer(b)
        

    fmt.Print("\n")
    // without defer
    
    fmt.Println("This is without defer")
    show_no_defer(b)

    fmt.Println("\n")

    
    



    // Problem 4
    

    p1 := person {
        first: "Utkarsh",
        last: "Sheikh",
        age: 21,
    }

    p1.speak()
    
    fmt.Println("\n")
    
    


    
    // problem 5


    s1 := square {
        side : 4,
    }

    s2 := square {
        side : 6,
    }

    c1 := circle {
        radius: 3,
    }

    c2 := circle {
        radius : 6.8,
    }

    fmt.Println(s1.area())

    fmt.Println(c1.area())

    info(s2)

    info(c2)

    fmt.Println("\n")

    // Problem 6

    q := func() int{
        return 45
    }

    fmt.Println(q())


    


    // Problem 7

    j := func() (int,string) {
        return 89, "HAHA"
    }

    fmt.Println(j())

    var f func() (int,string) 

    f = j

    f()

    var r int = 7

    fmt.Println(r)

    // functions are first class citizens

    // We can use identifiers to assign them function as , functions are also type, look at the code below so interesting.

    var g func() = func(){
        fmt.Println("Hello World")
    }
    g()

    



    // Problem 8

    
    e := func() func() int {
        return func() int{
            return 78
        }
    }()


    fmt.Printf("%T\n",e)

    fmt.Println("\n")


    

    // Problem 9

    {
        x := []int{1,8,9,10,6,34,2,3}

        y := sum_all(sum_arr, x)

        fmt.Println(y)
    }


    fmt.Println("\n")

    // Problem 10

    
    {
        g := func( x []int) int{

            if len(x) == 0 {
                return 0
            }

            if len(x) == 1 {
                return 1
            }

            return x[0] + x[len(x)-1]
        }

        c := func( f func(x []int) int, y []int) int{

            return g(y)
        }

        d := c(g, []int{1,2,4,56,7,89})
        
        fmt.Println(d)
    }

}






// Problem 1
func foo() int {
    return 67
}

func bar() (string, int) {
    return "xtring", 90
}


// Problem 2

// pass variadic parameter
func sum_var(x ...int) int {
    sum := 0
    
    for _ , v := range x{

        sum += v
    }

    return sum
}


// pass a parameter of type int-slice
func sum_arr(x []int) int {

    sum := 0

    for _ , v := range x{

        sum += v
    }

    return sum
}

func show_defer(x []int) {

    defer fmt.Println("sorry i came second",sum_arr(x))

    fmt.Println("I will be running first")
} 

func show_no_defer(x []int) {

    fmt.Println("sorry i came second",sum_arr(x))

    fmt.Println("I will be running first")
} 




// Problem 9

func sum_all(f func(a []int) int, x []int) int {
    return f(x) 
}

