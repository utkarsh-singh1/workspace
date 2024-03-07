package nestedtemplate

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("template1/*.gohtml"))
}

func NestedTemplates() {

	err := tmp.ExecuteTemplate(os.Stdout,"index1.gohtml",nil )
	if err != nil {
		log.Fatalln("Error =",err)
	}
	
}
