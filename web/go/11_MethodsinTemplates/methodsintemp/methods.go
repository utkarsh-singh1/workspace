package methodsintemp

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseFiles("index.gohtml"))
}

type person struct {
	First string
	Age int
}

func (p person) ReturnNum() int {
	return 7
}

func (p person) AgeDoubler() int {
	return p.Age * 2
}

func (p person) Doubler(x int) int{
	return x*2
}


func MethodinTemp() {

	p1 := person{
		"Utkarsh",
		30,
	}

	err := tmp.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln("Error =",err)
	}
}
