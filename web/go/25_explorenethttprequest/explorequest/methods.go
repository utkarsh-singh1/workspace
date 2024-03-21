package explorequest

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tmp1 *template.Template

func init() {

	tmp1 = template.Must(template.ParseFiles("index1.gohtml"))
	
}


type person1 int

func (p person1) ServeHTTP( w http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {

		log.Fatalln("Error =",err)
	}

	data := struct{
		Method string
		Submissions url.Values
	}{

		req.Method,
		req.Form,
	}
	tmp1.ExecuteTemplate(w, "index1.gohtml", data)
	
}


func PrintingHttpMethods() {

	var p1 person1
	http.ListenAndServe(":8080", p1)
	
}


