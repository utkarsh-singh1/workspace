package explorequest

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tmp3 *template.Template

func init() {

	tmp3 = template.Must(template.ParseFiles("index3.gohtml"))
}

type person3 int

func (p person3) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {
		log.Fatalln("Error =",err)
	}


	data := struct {
		Method     string
		Submissions url.Values
		URL        *url.URL
		Header http.Header
	}{
		req.Method,
		req.Form,
		req.URL,
		req.Header,
	}

	tmp2.ExecuteTemplate(w, "index3.gohtml", data)
}

func PrintingHttpHeaders() {

	var p3 person3

	http.ListenAndServe(":8080", p3)

}
