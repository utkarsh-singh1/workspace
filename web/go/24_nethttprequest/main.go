package main

import (
	"html/template"
	"log"
	"net/http"
)


type person int

var tmp *template.Template

func init() {

	tmp = template.Must(template.ParseFiles("index.gohtml"))
}

func (p person) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {
		log.Fatalln("Error =",err)
	}

	tmp.ExecuteTemplate(w, "index.gohtml", req.Form)
}


func main() {

	var man person

	http.ListenAndServe(":8080",man)
	
}
