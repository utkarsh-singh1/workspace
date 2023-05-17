package main

import "fmt"

func main(){

    // Go works with UTF-8 encoding system
    
    // to know more about string type ==>> https://go.dev/ref/spec#String_types

    // Now according to upper link - string can also be called as slice of bytes means ==>>

    s := "hello world "
    

    bs := []byte(s) // byte also known as uint8 refer 09_numeric.go file - also point here is there is type conversion 

    fmt.Println(bs)
    /* o/p of this is =>> [104 101 108 108 111 32 119 111 114 108 100 32], as we know this is slice of byte and these numbers are ascii code of letters =>> "hello world" space included so ascii code of h is 104 (source - https://www.cs.cmu.edu/~pattis/15-1XX/common/handouts/ascii.html) means we are using coding schemes to represent our characters. All these numbe    rs are in decimal form. 
    */
    fmt.Printf("%T\n", bs)

    // this =>>``<<= is also known as string literal, we can use it to use qoutes (double/single) as other character 
    q := `"hello 
    you know me very" well`

    
    fmt.Println(s)
    
    fmt.Println(q)
    
    fmt.Printf("%T\n",s)
    

    // As in above we notice that the ascii code represent chars in decimal way but in UTF-8 format chars are in hexa-decimal format to see for each character print it with %#U
    for i:=0 ; i<len(s) ; i++{
        fmt.Printf("%#U\t", s[i])
    }

    /* output of this part is =>> [U+0068 'h'      U+0065 'e'      U+006C 'l'      U+006C 'l'      U+006F 'o'      U+0020 ' '      U+0077 'w'     U+006F 'o'       U+0072 'r'      
    U+006C    'l'      U+0064 'd'      U+0020 ' ' ] - all these are in hexa decimal form and each code point is called "rune" (int32) */

    fmt.Println("")

    // Concatination in go-lang
    fmt.Println("Hello" + "World")


    for i,v := range s{

        // output in decimal form
        fmt.Println(i,v)

        // for output in hexa decimal form
        fmt.Printf("for each %d we have %#X\n", i,v)


    }

    // well well well this was quite a ride today in string - we have gone deep into it to understand how string works in golang but if wanna lnow more check this out =>>
    // "https://go.dev/blog/strings"
}
