package main

import (
	"encoding/json"
	"fmt"
)

type person []struct {
	Name   string `json:"Name"`
	Age    int    `json:"Age"`
	Loveit string `json:"Loveit"`
}

func main() {

	var s string

	s = `[
    {
        "Name":"utkarsh",
        "Age":18,
        "Loveit":"yes"
    },
    {
        "Name":"kyle",
        "Age":23,
        "Loveit":"no"
    }
    ]`

	bs := []byte(s)

	ps := person{
		{"jason", 34, "no"},
	}

	err := json.Unmarshal(bs, &ps)
	if err != nil {

		fmt.Println("Error =", err)

	}

	fmt.Println("Unmarshalled JSON ->", ps)
}
