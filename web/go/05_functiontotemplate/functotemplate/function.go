package functotemplate

import (
	"os"
	"strings"
	"text/template"
)

type person struct {
	first string
	last string
	age int
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
		first: "Hari",
		last: "Om",
		age: 1000,
	}

	b := person{

		first: "Utkarsh",
		last: "Singh",
		age: 24,
		
	}

	c := person{

		first: "Hola",
		last: "Bhola",
		age: 23,
	}

	sg := []person{a,b,c}


	tpl.ExecuteTemplate(os.Stdout,"index.gohtml",sg)
	
}
