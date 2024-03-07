package globalfunc

import (
	"log"
	"os"
	"text/template"
)

var tmp3 *template.Template

func init() {
	tmp3 = template.Must(template.ParseFiles("index3.gohtml"))
}

func ComparisoninTemplate() {

	err := tmp3.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("Error =",err)
	}
	
}
