package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {

	mux := httprouter.New()

	mux.GET("/", index)

	http.ListenAndServe(":8080", mux)

}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	err := tpl.ExecuteTemplate(w,"index.gohtml",nil)
	if err != nil {
		log.Fatalln("Error =", nil)
	}
	
}


