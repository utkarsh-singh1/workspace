package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseGlob("templates/*"))

}


func main() {

	http.HandleFunc("/", index )
	http.HandleFunc("/about", about )
	http.HandleFunc("/contact", contact )
	http.HandleFunc("/apply", apply )
	http.ListenAndServe(":8080", nil)

}


func about(w http.ResponseWriter, req *http.Request) {

	err := tpl.ExecuteTemplate(w, "about.gohtml", nil )
	if err != nil {

		log.Fatalln("Error =",err)

	}
}

func index(w http.ResponseWriter, req *http.Request) {

	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("Error =",err)
	}
}

func contact(w http.ResponseWriter, req *http.Request) {

	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	if err != nil {
		log.Fatalln("Error =",err)
	}
}

func apply(w http.ResponseWriter, req *http.Request) {

	if  req.Method == http.MethodPost {
		err := tpl.ExecuteTemplate(w , "applyProcess.gohtml", nil)
		if err != nil {
			log.Fatalln("Error =",err)
			
		}

		return
	}
	
	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	if err != nil {
		log.Fatalln("Error =",err)
	}
}


