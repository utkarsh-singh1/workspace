package main

import (
	"io"
	"net/http"
	"os"
)


func main() {

	http.HandleFunc("/", dog)
	http.HandleFunc("/local.png", dogPic)

	http.ListenAndServe(":8080", nil)

	
}


// func dog(w http.ResponseWriter, req *http.Request) {

// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")


// 	// As we does not serve the image file yet thats why there is error 404 on the response
// 	io.WriteString(w , `
// <!-- image does not serve -->
// <img src="/local.png" >
// `)

// }

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w , `
<img src="/local.png" />
`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {

	f , err := os.Open("local.png")
	if err != nil {
		http.Error(w , "File not found", 404)
		return
	}

	defer f.Close()

	io.Copy(w, f)
}

