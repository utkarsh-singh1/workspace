package functotemplate

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type person struct {
	First string
	Last string
	Age int
}

var tpl *template.Template


func FirstThree(str string) string {

	s := strings.TrimSpace(str)

	s = s[:3]

	return s
	
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": FirstThree,
}



func init() {

	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
	
}

func FuncinTemp() {

	a := person{
		First: "Hari",
		Last: "Om",
		Age: 1000,
	}

	b := person{

		First: "Utkarsh",
		Last: "Singh",
		Age: 24,
		
	}

	c := person{

		First: "Hola",
		Last: "Bhola",
		Age: 23,
	}

	sg := []person{a,b,c}


	err := tpl.ExecuteTemplate(os.Stdout,"index.gohtml",sg)
	if err != nil {
		log.Fatalln("Error =",err)
	}
}
