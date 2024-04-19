// In previous example FileServer will serve all the file in current root directory, now if i dont want to have all files serving execpt assets files we are gonna do it by this way

package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", dog)

	http.Handle("/assets/", http.StripPrefix("/assets",  http.FileServer(http.Dir("./assets"))))
	
	http.ListenAndServe(":8080", nil)

}


func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w ,`

<h1> Hello World </h1>
<img src="/assets/local.png"/> `)

}



