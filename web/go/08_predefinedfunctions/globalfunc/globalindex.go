package globalfunc

import (
	"log"
	"os"
	"text/template"
)

var tmp1 *template.Template

func init() {
	tmp1 = template.Must(template.ParseFiles("index2.gohtml"))
}

func IndexinTemplates(){

	xi := []string{"one","two","three","four"}

	err := tmp1.Execute(os.Stdout,xi)
	if err != nil {
		log.Fatalln("Error =",err)
	}
	
}
