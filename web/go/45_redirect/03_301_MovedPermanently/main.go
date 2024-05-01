package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo )

	http.HandleFunc("/bar", bar)
	
	http.ListenAndServe(":8080", nil)
	
}


func foo(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Request of foo is",req.Method)
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "Hello there from foo")

}


func bar(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Request of bar is", req.Method)

	// w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// io.WriteString(w , "Do you think you can come here?" )
	// Instead of using below code ->

	// w.Header().Set("Location", "/")
	// w.WriteHeader(http.StatusSeeOther)

	// We can use

	http.Redirect(w , req, "/", http.StatusMovedPermanently)

}
