package goparser

import (
	"log"
	"os"
	"text/template"
)


func ParseGlobes(){

	tmp, err := template.ParseGlob("template/*.gohtml")
	if err != nil {
		log.Fatalln("Error =",err)
	}

	err = tmp.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("Error =",err)
	}

	err = tmp.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln("Error =",err)
	}

	err = tmp.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln("Error =",err)
	}

	err = tmp.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln("Error =",err)
	}
}
