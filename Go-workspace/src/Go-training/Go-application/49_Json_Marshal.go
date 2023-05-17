/*
     _ ____   ___  _   _
    | / ___| / _ \| \ | |
 _  | \___ \| | | |  \| |
| |_| |___) | |_| | |\  |
 \___/|____/ \___/|_| \_|

*/

// Json store data-structure that can be used to communicate between application.

// https://pkg.go.dev/encoding/json@go1.20.3

// Learn about it here -->> Marshal -->> https://pkg.go.dev/encoding/json@go1.20.3#Marshal

// How to use it --->>> func Marshal(v any) ([]byte, error)

// Marshal returns the JSON encoding of v.

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {

	p1 := person{
		First: "Utkarsh",
		Last:  "Singh",
		Age:   25,
	}

	p2 := person{
		First: "Miss", 
        Last: "Moneypenny",
		Age: 67,
	}


    people := []person{
        p1,
        p2,
    }

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println("err : ", err)
	}

	js := string(bs)

	fmt.Println(js)

    Single := "me"

    bs1 , err1 := json.Marshal(Single)

    if err1 != nil {
        fmt.Println("new err : ", err1)
    }

    fmt.Println(string(bs1))

	fmt.Println(p1)

}
