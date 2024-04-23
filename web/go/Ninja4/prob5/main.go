// Serve the files in the "starting-files" folder
// To get your images to serve, use:
// 	func StripPrefix(prefix string, h Handler) Handler
// 	func FileServer(root FileSystem) Handler
// Constraint: you are not allowed to change the route being used for images in the template file

package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", pics)

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))) )

	http.ListenAndServe(":8080", nil)
}

func pics(w http.ResponseWriter, req *http.Request) {

	tpl, err := template.ParseGlob("templates/index.gohtml")
	if err != nil {
		log.Fatalln("Error =",err)
	}

	tpl.Execute(w, nil)
}
