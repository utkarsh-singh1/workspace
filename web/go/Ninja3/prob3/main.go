// Prob3
// Take the previous program and change it so that:
// func main uses http.Handle instead of http.HandleFunc
// Contstraint: Do not change anything outside of func main

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

	

	http.Handle("/", http.HandlerFunc(Index))
	
	http.ListenAndServe(":8080", nil )
	
}

