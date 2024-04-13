// Problem 2
// Take the previous program in the previous folder and change it so that:
// a template is parsed and served
// you pass data into the template

package main

import (
	"net/http"
	"text/template"
)



func Index(w http.ResponseWriter, req *http.Request) {

	var tpl *template.Template
	
	tpl = template.Must(template.ParseFiles("index.gohtml"))

	data := "Say my Name -> Utkarsh Singh 😎"
	
	tpl.Execute(w , data)
	 
}

func main() {

	

	http.HandleFunc("/", Index)
	
	http.ListenAndServe(":8080", nil )
	
}
