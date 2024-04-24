package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", dog )

	http.Handle("/favicon.ico", http.NotFoundHandler() )

	http.ListenAndServe(":8080", nil)

}

func dog(w http.ResponseWriter, req *http.Request ) {

	fmt.Println(req.URL)

	fmt.Fprintln(w, "go look at your terminal")

}

func NotFoundHandler() http.Handler {

	return http.HandlerFunc(http.NotFound)

}
 
func NotFound(w http.ResponseWriter, req *http.Request) {
	http.Error(w , "404 page does not exist", 404)
}
 
