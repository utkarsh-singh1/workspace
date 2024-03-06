package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

func monthdateyearformat(t time.Time) string {

	return t.Format("01-02-2006")
}

func main() {

	// dateformat.Dateformat()


	fm := template.FuncMap{

		"mdy" : monthdateyearformat,
		
	}

	tmp := template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))

	err := tmp.ExecuteTemplate(os.Stdout,"index.gohtml", time.Now())
	if err != nil {
		log.Fatalln("Error =",err)
	}
	
}
