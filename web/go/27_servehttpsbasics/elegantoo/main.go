package main

import (
	"fmt"
	"net/http"
)



func d(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(w , "hello woof woof")
}



func c(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(w, "hello nya nya")
}


func main() {



	http.HandleFunc("/dog", d)

	http.HandleFunc("/cat/", c)
	
	http.ListenAndServe(":8080",nil)

}
