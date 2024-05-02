package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", set )
	http.HandleFunc("/read", read )
	http.ListenAndServe(":8080", nil)

}

func set(w http.ResponseWriter, req *http.Request) {


	http.SetCookie(w , &http.Cookie{
		Name: "hello",
		Value: "Everynya",
	})

	fmt.Fprintln(w ,"Look for cookie in your browser at -> /inspect/application/dev-tools" )

}

func read(w http.ResponseWriter, req *http.Request) {

	c,err := req.Cookie("hello")
	if err != nil {

		http.Error(w , "Cookie not fpund", 404)

	}

	fmt.Fprintln(w , "Your cookie", c )

}
