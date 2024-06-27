package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)


type user struct {

	USERNAME, FirstName, LastName, Password string
}

var dbSession map[string]string

var dbUser map[string]user

func main() {


	http.HandleFunc("/",foo)

	http.HandleFunc("/signup", signUp)

	http.HandleFunc("/bar", bar)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	
	http.ListenAndServe(":8080", nil)


}


func foo( w http.ResponseWriter, req *http.Request) {


	u := getUser(w , req )
	

}


func signUp( w http.ResponseWriter, req *http.Request) {


	// If user already exists
	if alreadyLoggedIn(req) {

		http.Redirect(w , req , "/", http.StatusSeeOther )

	}


	// If user does not exist

	if req.Method == http.MethodPost {

		un := req.FormValue("Username")
		f := req.FormValue("Firstname")
		s := req.FormValue("Lastname")
		p := req.FormValue("Password")
		

	}

	
	
}

func bar( w http.ResponseWriter, req *http.Request) {


	u := getUser(w , req )

	if !alreadyLoggedIn(req) {

		http.Redirect(w , req, "/", http.StatusSeeOther)

		return

	}
	
	
}

