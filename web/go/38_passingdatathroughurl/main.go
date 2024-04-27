///State is a persistant awareness , who is communicating with the server

package main

import (
	"fmt"
	"net/http"
)

func main() {


	http.HandleFunc("/", handleit)
	
	http.ListenAndServe(":8080", nil )
	
}

func handleit( w http.ResponseWriter, req *http.Request) {

	v := req.FormValue("q")

	fmt.Fprintln(w , "The passed value is q = "+v," and a is "+ v )
	

}
