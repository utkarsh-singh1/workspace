package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", dog)

	http.HandleFunc("/local.png", servePic)
	
	http.ListenAndServe(":8080", nil)

}


func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w ,`

<h1> Hello World </h1>
<img src="/local.png"/> `)

}

func servePic(w http.ResponseWriter, req *http.Request) {


	f, err := os.Open("local.png")
	if err != nil {
		http.Error(w , "File not found", 404)
		return
	}

	defer f.Close()

	http.ServeFile(w , req, f.Name())
}


