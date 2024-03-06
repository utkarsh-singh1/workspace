package globalfunc

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

func init() {

	tmp = template.Must(template.ParseFiles("index1.gohtml"))
}

type person struct {
	First string
	Last string
	Age string
}

func GlobalAnd() {

	a := person{
		"Utkarsh",
		"Singh",
		"24",
	}

	b := person{
		"Kishore",
		"Kumar",
		"Dead",
	}

	c := person{

		"",
		"Don'tknow",
		"Ageless",
	}

	xi := []person{a,b,c}
	
	err := tmp.Execute(os.Stdout, xi)
	if err != nil {
		log.Fatalln("Error =", err)
	}
	
}
