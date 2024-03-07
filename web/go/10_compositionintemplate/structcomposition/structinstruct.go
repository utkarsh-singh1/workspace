package structcomposition

import (
	"log"
	"os"
	"text/template"
)

var tmp2 *template.Template

func init() {
	tmp2 = template.Must(template.ParseFiles("index2.gohtml"))
}

type person struct{
	First , Second string
}

type agent struct{
	person
	Ltk bool
}

func StructInHeritance() {

	p1 := person{"Utkarsh", "Singh"}

	a1 := agent{p1 , true}

	err := tmp2.Execute(os.Stdout, a1)
	if err != nil {
		log.Fatalln("Error =",err)
	}
}
