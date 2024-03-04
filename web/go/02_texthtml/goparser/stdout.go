package goparser

import (
	"log"
	"os"
	"text/template"
)

func Stdout() {

	tmp, err := template.ParseFiles("index.gohtml")

	if err != nil {
		log.Fatalln("Error ", err)
	}

	err = tmp.Execute(os.Stdout, "hello")

	if err != nil {
		log.Fatalln("Error", err)
	}
	
	
}
