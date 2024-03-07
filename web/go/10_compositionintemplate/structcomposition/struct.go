package structcomposition

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

var ( x int ; y string ; z bool)  

func init() {

	tmp = template.Must(template.ParseFiles("index1.gohtml"))
}

type man struct {
	First , Second string
}

func Structs() {

	m1 := man{ "Utkarsh", "Singh"}

	err := tmp.Execute(os.Stdout, m1)
	if err != nil {
		log.Fatalln("Error =",err)
	}
	
}
