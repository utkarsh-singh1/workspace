package goparser

import (
	"log"
	"os"
	"text/template"
)

func FileOut(s string) {

	tmp, err := template.ParseFiles(s)

	if err != nil {
		log.Fatalln("Error", err)
	}

	nf, err := os.Create("index.html")

	err = tmp.Execute(nf, nil)

	if err != nil {
		log.Fatalln("Err", err)
	}
	
}
