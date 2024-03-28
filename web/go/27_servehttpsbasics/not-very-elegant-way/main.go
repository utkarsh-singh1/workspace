package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func ( h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	switch req.URL.Path {

	case "/dog" :
		fmt.Fprintln(w , "Hello woof-woof")
	case "/cat" :
		fmt.Fprintln(w , "Hello nya nya")
	}

}

func main() {

	var d hotdog

	http.ListenAndServe(":8080",d) 
	
}
