/*
     _ ____   ___  _   _ 
    | / ___| / _ \| \ | |
 _  | \___ \| | | |  \| |
| |_| |___) | |_| | |\  |
 \___/|____/ \___/|_| \_|
                         
*/


// Unmarshal -->> https://pkg.go.dev/encoding/json@go1.20.3#Unmarshal

// func Unmarshal(data []byte, v any) error

// Remember -->> Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v. If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.

// V is pointer of any type. 


package main

import (
      "fmt"
		"encoding/json"
   )

type  person struct {
   First string `json:"First"`
   Last  string `json:"Last"`
   Age   int    `json:"Age"`
}

func main() {

   s := `[{"First":"Utkarsh","Last":"Singh","Age":25},{"First":"Miss","Last":"Moneypenny","Age":67}]`

   bs := []byte(s)

   fmt.Printf("%T\n",s)

   fmt.Printf("%T\n",bs)

   people := []person{}

   err := json.Unmarshal(bs, &people)

   if err != nil {
      fmt.Println("An error is here :-",err)
   }

   for i ,v := range people {
      fmt.Println("for person",i,"the information is",v.First,v.Last,"Age -->> ",v.Age)
   }

}