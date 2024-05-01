package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", defaulttype)

	http.Handle("/favicon.ico",  http.NotFoundHandler() )
	
	http.ListenAndServe(":8080", nil)

}

func defaulttype(w http.ResponseWriter, req *http.Request) {

	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln("Error =",err)
	}

	bs := make([]byte, req.ContentLength)

	req.Body.Read(bs)

	s := string(bs)

	tpl.ExecuteTemplate(w ,"index.gohtml", s)
}
