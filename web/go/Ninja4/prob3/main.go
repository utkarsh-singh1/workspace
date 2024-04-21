// Serve the files in the "starting-files" folder
// Use "http.FileServer"

package main

import (
	"log"
	"net/http"
	"text/template"
)



func main() {

	fs := http.FileServer(http.Dir("public"))
	
	http.Handle("/pics/", fs)
	http.HandleFunc("/", dog)
	
	http.ListenAndServe(":8080", nil)

}

func dog(w http.ResponseWriter, req *http.Request) {

	tpl,err := template.ParseGlob("templates/index.gohtml")
	if err != nil {

		log.Fatalln("Error =",err)
	}

	tpl.Execute(w, nil)	
}


