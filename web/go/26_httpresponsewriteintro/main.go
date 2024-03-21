package main

import (
	"fmt"
	"net/http"
)


type person int

func (p person) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Here", "I am do you know me")
	w.Header().Set("Content-Type","text/plain; charset=utf-8")

	/* Run above code and also try
	   w.Header().Set("Content-Type","text/html; charset=utf-8 ")
	*/
	fmt.Fprintln(w, "<h1> Haha you knwo me</h1>")
}

func main() {

	var p1 person
	http.ListenAndServe(":8080",p1)
	
}
