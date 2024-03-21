package explorequest

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tmp4 *template.Template

func init() {

	tmp4 = template.Must(template.ParseFiles("index4.gohtml"))
}

type person4 int

func (p person4) ServeHTTP(w http.ResponseWriter, req *http.Request) {

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

	tmp2.ExecuteTemplate(w, "index4.gohtml", data)
}

func PrintingHttpContentLength() {

	var p4 person4

	http.ListenAndServe(":8080", p4)

}
