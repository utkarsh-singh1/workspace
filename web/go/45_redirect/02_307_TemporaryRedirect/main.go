package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", foo )

	http.HandleFunc("/barred", barred )

	http.HandleFunc("/bar", bar)
	
	http.ListenAndServe(":8080", nil)
	
}


func foo(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Request of foo is",req.Method)
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "Hello there from foo")

}

func barred(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Request of barred is ", req.Method)
	
	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln("Error =",err)
	}

	tpl.Execute(w , nil)

}

func bar(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Request of bar is", req.Method)

	// Instead of using below code ->

	// w.Header().Set("Location", "/")
	// w.WriteHeader(http.StatusSeeOther)

	// We can use

	http.Redirect(w , req, "/", http.StatusTemporaryRedirect)

}
