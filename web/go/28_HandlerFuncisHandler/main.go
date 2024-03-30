package main

import (
	"fmt"
	"net/http"
)


func d(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(w, "Do you want a dog?")
	
}

func c(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(w ,"Why do you want a cat?")
}

func main() {

	http.Handle("/dog", http.HandlerFunc(d))

	http.Handle("/cat", http.HandlerFunc(c))


	http.ListenAndServe(":8080", nil)
}
