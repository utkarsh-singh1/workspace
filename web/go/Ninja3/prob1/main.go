// problem - 1

// ListenAndServe on port ":8080" using the default ServeMux.

// Use HandleFunc to add the following routes to the default ServeMux:

// "/" "/dog/" "/me/

// Add a func for each of the routes.

// Have the "/me/" route print out your name.

package main

import (
	"fmt"
	"net/http"
)


func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
		fmt.Fprintln(w , "Welcom to my Havely")
	})

	http.HandleFunc("/dog/", func(w http.ResponseWriter, req *http.Request){
		fmt.Fprintln(w , "You have come to see dog, Surprise 😛")
	})

	http.HandleFunc("/me/", func(w http.ResponseWriter, req *http.Request){
		fmt.Fprintln(w , "Hi This is Utkarsh 😎")
	})
	
	http.ListenAndServe(":8080", nil)
}
