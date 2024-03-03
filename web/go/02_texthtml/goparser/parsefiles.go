package goparser

import (
	"log"
	"os"
	"text/template"
)

func Parsefile() {


	tmp, err := template.ParseFiles("one.gohtml")
	if err != nil {
		log.Fatalln("Error =",err)
	}

	//err = tmp.Execute(os.Stdout, nil)

	// if err != nil {
	// 	log.Fatalln("Error =",err)
	// }

	tmp, err = tmp.ParseFiles("two.gohtml")

	err = tmp.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("Error =",err)
	}

	
	err = tmp.ExecuteTemplate(os.Stdout,"two.gohtml", nil)
	if err != nil {
		log.Fatalln("Error =", err)
	}
}
