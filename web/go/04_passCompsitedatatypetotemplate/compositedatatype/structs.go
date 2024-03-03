package compositedatatype

import (
	"html/template"
	"log"
	"os"
)

var tmp3 *template.Template

func init() {
	
	tmp3 = template.Must(template.ParseFiles("index4.gohtml"))
}


type person struct {
	First string
	Last string
}

func StrcutstoTemp() {

	p := person{
		First: "Utkarsh",
		Last: "Singh",
	}

	err := tmp3.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln("Error =",err)
	}
	
	

}
