package function

import "fmt"

type Utkarsh struct {

	FullName string
	Age int
	Money int
}

func (u Utkarsh) Details () {
	fmt.Println(u.FullName,u.Age,u.Money)
}
