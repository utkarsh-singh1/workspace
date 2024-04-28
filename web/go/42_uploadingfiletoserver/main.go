// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// )

// func main() {

// 	http.HandleFunc("/", readfile)
// 	http.Handle("/fevicon.ico", http.NotFoundHandler())
// 	http.ListenAndServe(":8080", nil)

// }

// func readfile(w http.ResponseWriter, req *http.Request) {

// 	var s string
// 	fmt.Println(req.Method)
// 	if req.Method == http.MethodPost {

// 		f, h, err := req.FormFile("q")
// 		if err != nil {
// 			log.Fatalln("Error =",err)
// 			return
// 		}

// 		fmt.Print("\nFile:",f,"\nHeader:",h,"\nError:",err)

// 		defer f.Close()

// 		bs , err := io.ReadAll(f)
// 		if err != nil {
// 			log.Fatalln("Error =",err)
// 			return
// 		}

// 		s = string(bs)
// 	}

// 	w.Header().Set("Content-Type", "text/html; charset=utf-8 ")
// 	io.WriteString(w ,`
// <form enctype="multipart/form-data metod="POST">
// <input type="file" name="q" />
// <input type="submit" />
// </form>
// <br />
// ` +s )

// }

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {

		// open
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// for your information
		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		// read
		bs, err := io.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)
}
