package main

import (
	"encoding/json"
	"fmt"
)

type user struct {

	Name string
	Age int
	Loveit string
}

func main() {

	u1 := user{

		"utkarsh", 18, "yes",
	}

	u2 := user{

		"kyle", 23, "no",
		
	}

	u := []user{u1,u2}
	
	bs, err := json.Marshal(u)
	if err != nil {

		fmt.Println("Error =",err)
	}

	fmt.Println(string(bs))
	
}
