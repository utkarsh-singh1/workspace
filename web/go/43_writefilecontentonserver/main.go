package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

func main() {

	http.HandleFunc("/favicon.ico", http.NotFound)
	http.HandleFunc("/", handleit)
	http.ListenAndServe(":8080",nil)

}

func handleit(w http.ResponseWriter, req *http.Request) {

	var s string

	tpl := template.Must(template.ParseGlob("templates/*"))

	if req.Method == http.MethodPost {

		f, h, err := req.FormFile("q")
		if err != nil {
			log.Fatalln("Error =",err)
		}

		defer f.Close()

		fmt.Println("\nfile:",f,"\nheader:",h,"\nError:",err)

		//read
		bs, err := io.ReadAll(f)
		if err != nil {

			log.Fatalln("Error =",err)

		}

		s = string(bs)

		// write to file

		dst, err := os.Create(filepath.Join("./user/", h.Filename))
		if err != nil {

			log.Fatalln("Error =",err)
		}
		
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {

			log.Fatalln("Error =",err)
		}		

	}

	w.Header().Set("Content-Type", "text/html ; charset=utf-8")
	tpl.ExecuteTemplate(w, "index.gohtml", s)

}
