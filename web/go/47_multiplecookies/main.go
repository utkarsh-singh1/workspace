package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.ListenAndServe(":8080", nil)

}

func set ( w http.ResponseWriter, req *http.Request) {


	http.SetCookie(w , &http.Cookie{
		Name: "1st",
		Value: "First cookie",
	})

	w.Header().Set("Content-Type" , "text/html ; charset=utf-8")
	fmt.Fprintln(w , "set first cookie")

}


func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("1st")
	if err ==  nil {

		fmt.Fprintln(w ,"Here is cookie #1",c )
	}else {
		log.Println("Error =",err)
	}

	c2, err := req.Cookie("2nd")
	if err !=  nil {
		log.Println("Error =",err)
	}else {

		fmt.Fprintln(w ,"Here is cookie #2",c2 )
	}

	c3, err := req.Cookie("3rd")
	if err !=  nil {
		log.Println("Error =",err)
	}else {

		fmt.Fprintln(w ,"Here is cookie #3",c3 )
	}
}

func abundance(w http.ResponseWriter, req *http.Request) {

	http.SetCookie(w , &http.Cookie{
		Name: "2nd",
		Value: "Second cookie",
	})

	http.SetCookie(w , &http.Cookie{
		Name: "3rd",
		Value: "Third cookie",
	})

	w.Header().Set("Content-Type" , "text/html ; charset=utf-8")
	fmt.Fprintln(w , "set 2 cookies")
}
