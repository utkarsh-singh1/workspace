// Serve the files in the "starting-files" folder
// To get your images to serve, use:
// 	func StripPrefix(prefix string, h Handler) Handler
// 	func FileServer(root FileSystem) Handler
// Constraint: you are not allowed to change the route being used for images in the template file

package main

import (
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", pics)
	http.Handle("/resources/",http.StripPrefix("/resources", http.FileServer(http.Dir("./public")))) 
	http.ListenAndServe(":8080",nil )

}

func pics( w http.ResponseWriter, req *http.Request) {

	tpl := template.Must(template.ParseGlob("templates/index.gohtml"))

	tpl.Execute(w , nil)
}
