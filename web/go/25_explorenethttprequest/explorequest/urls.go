package explorequest

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tmp2 *template.Template

func init() {

	tmp2 = template.Must(template.ParseFiles("index2.gohtml"))
}

type person2 int

func (p person2) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {
		log.Fatalln("Error =",err)
	}


	data := struct {
		Method     string
		Submissions url.Values
		URL        *url.URL
	}{
		req.Method,
		req.Form,
		req.URL,
	}

	tmp2.ExecuteTemplate(w, "index2.gohtml", data)
}

func PrintingHttpURLS() {

	var p1 person2

	http.ListenAndServe(":8080", p1)

}
