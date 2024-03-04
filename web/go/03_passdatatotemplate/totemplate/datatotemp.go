package totemplate

import (
	"log"
	"os"
	"text/template"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseFiles("index.gohtml")) 
}

func DataToTemplate(){

	err := tmp.ExecuteTemplate(os.Stdout, "index.gohtml", 42)
	if err != nil {
		log.Fatalln("Error =",err)
	}

}
