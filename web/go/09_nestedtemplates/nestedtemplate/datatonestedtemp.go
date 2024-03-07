package nestedtemplate

import (
	"log"
	"os"
	"text/template"
)

var tmp2 *template.Template

func init() {
	tmp2 = template.Must(template.ParseGlob("template2/*.gohtml"))
}

func DatatoNestedTemplates() {

	err := tmp2.ExecuteTemplate(os.Stdout,"index2.gohtml",42 )
	if err != nil {
		log.Fatalln("Error =",err)
	}
	
}
