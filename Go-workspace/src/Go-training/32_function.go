
/*
_____                 _   _                 
|  ___|   _ _ __   ___| |_(_) ___  _ __  ___ 
| |_ | | | | '_ \ / __| __| |/ _ \| '_ \/ __|
|  _|| |_| | | | | (__| |_| | (_) | | | \__ \
|_|   \__,_|_| |_|\___|\__|_|\___/|_| |_|___/
 
*/

package main

import "fmt"

func main() {
    fmt.Println("Hello World")
    foo()
    bar("Utkarsh")

    S1 := woo("Ms. Monjolika")
    fmt.Println(S1)

    x,y := mouse("Miss", "Moneypenny")
    fmt.Println(x,y)
}


// func (r receiver) identifier(parameters) (return(S)) {...}   //  When you define a func and pass a value it is called parameter , when you call func a func and pass some value it is called argument.

func foo() {
    fmt.Println("Hello foo")
}


// Everything in Go is PASS BY VALUE
// Pass an argument
func bar(s string) {
    fmt.Println("Hello,",s)
}

// Single return
func woo(st string) string{
    return fmt.Sprint("Hello, ",st)
}


// Multiple returns
func mouse(fn string, ln string) (string, bool) {
    a  := fmt.Sprint(fn," ", ln, `,Says "Hello"`)
    b := true

    return a,b
}

