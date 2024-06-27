package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	USERNAME string
	First    string
	Last     string
}

var tpl *template.Template

var dbUser = map[string]user{}
var dbSession = map[string]string{}

func init() {

	tpl = template.Must(template.ParseGlob("template/*"))

}

func main() {

	http.HandleFunc("/", foo)

	http.HandleFunc("/bar", bar)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("session")
	if err != nil {

		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}

	http.SetCookie(w, c)

	// check if user exist or not
	var u user
	if un, ok := dbSession[c.Value]; ok {

		u = dbUser[un]

	}

	// if user does not exist create one

	if req.Method == http.MethodPost {

		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		u = user{un, f, l}
		dbSession[c.Value] = un
		dbUser[un] = u

	}

}

func bar(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	un, ok := dbSession[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUser[un]

}
