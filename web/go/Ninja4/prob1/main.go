// ListenAndServe on port 8080 of localhost
// For the default route "/" Have a func called "foo" which writes to the response "foo ran"

// For the route "/dog/" Have a func called "dog" which parses a template called "dog.gohtml" and writes to the response "

// This is from dog
// " and also shows a picture of a dog when the template is executed.
// Use "http.ServeFile" to serve the file "dog.jpeg"

package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {

	http.Handle("/", http.HandlerFunc(foo) )
	http.HandleFunc("/dog/", dog )
	http.HandleFunc("/dog/local.png", func(w http.ResponseWriter, req *http.Request) {
		f, err := os.Open("local.png")
		if err != nil {

			log.Fatalln("Error =",err)
		}
		
		http.ServeFile(w, req , f.Name())
		
	} )
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w , "foo")
	
}

func dog(w http.ResponseWriter, req *http.Request) {

	tpl, err := template.ParseFiles("dog.gohtml")

	if err != nil {

		log.Fatalln("Error =", err)

	}
	
	tpl.Execute(w , nil)
}
