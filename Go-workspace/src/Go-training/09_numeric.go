package main

import (
	"fmt"
	"runtime"
)


var x int

var y float64

func main(){
    // You can find about it more here - https://go.dev/ref/spec#Numeric_types

    /*

    the value of int is generally architecture of the particular device u are using - like found amd64 architecture will represent int as int64

    uint8       the set of all unsigned  8-bit integers (0 to 255)
    uint16      the set of all unsigned 16-bit integers (0 to 65535)
    uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
    uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

    int8        the set of all signed  8-bit integers (-128 to 127)
    int16       the set of all signed 16-bit integers (-32768 to 32767)
    int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
    int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

    float32     the set of all IEEE-754 32-bit floating-point numbers
    float64     the set of all IEEE-754 64-bit floating-point numbers

    complex64   the set of all complex numbers with float32 real and imaginary parts
    complex128  the set of all complex numbers with float64 real and imaginary parts

    byte        alias for uint8
    rune        alias for int32


    */

    var a int8 = 89  // range between -128 to 127 -> 256 values (2^8) - can also try assigning var a int8 = -129 or 128

    fmt.Println(a)

    x = 9
    y = 8.987

    fmt.Println(x)
    fmt.Printf("%T\n",x)

    fmt.Println(22*3*3/7)

    fmt.Println(y)
    fmt.Printf("%T\n", y)

    // To know more about runtime package go here ==>> https://pkg.go.dev/runtime

    fmt.Println(runtime.GOOS)
    fmt.Println(runtime.GOARCH)

}

func hello(){
  fmt.Println("Hello")
}
