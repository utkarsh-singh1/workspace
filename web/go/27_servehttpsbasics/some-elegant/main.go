package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(w , "hello woof woof")
}

type hotcat int

func (c hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(w, "hello nya nya")
}


func main() {

	var d hotdog

	var c hotcat

	mux := http.NewServeMux()

	mux.Handle("/dog", d)

	mux.Handle("/cat/", c)
	
	http.ListenAndServe(":8080",mux)

}
