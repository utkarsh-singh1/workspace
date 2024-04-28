package main

import (
	"net/http"
	"text/template"
)

type person struct {

	FirstName string
	LastName string
	Subscribed bool
}

func main() {

	http.HandleFunc("/", servefile )
	http.ListenAndServe(":8080", nil)

}


func servefile(w http.ResponseWriter, req *http.Request) {


	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("subscribe") == "on"
	
	tpl := template.Must(template.ParseGlob("templates/*"))

	tpl.ExecuteTemplate(w , "index.gohtml", person{f,l,s}  )

}
