package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Username, First, Last, Password string
}

var dbSession map[string]string

var dbUser map[string]user

func main() {

	http.HandleFunc("/", index)

	http.HandleFunc("/signup", signUp)

	http.HandleFunc("/bar", bar)

	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {

	getUser(w , req ) 

}

func signUp(w http.ResponseWriter, req *http.Request) {


	if alreadyLoggedIn(req) {

		http.Redirect(w , req , "/", http.StatusSeeOther)
	}

	// if User does not exist create one
	if req.Method == http.MethodPost {

		un := req.FormValue("Username")
		f := req.FormValue("First")
		l := req.FormValue("Last")
		p := req.FormValue("Password")

		// if username is already exist or not

		if _, ok := dbUser[un] ; ok {

			http.Error(w , "Username already exist", http.StatusForbidden)
			return

		}

		// Create a session for new signUp

		c := &http.Cookie{Name: "session", Value: uuid.NewV4().String()}
		http.SetCookie(w , c)

		// Attach session to a Username

		dbSession[c.Value] = un

		// Encrypt Password

		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost )
		if err != nil {

			http.Error(w , "Internal Server Error", http.StatusInternalServerError)
		}

		// Add userdetail to username

		dbUser[un] = user{un,f,l,string(bs)}

		// Redirect Back to Homepage after signup

		http.Redirect(w , req, "/", http.StatusSeeOther)
		return
		
	}
	
}

func bar(w http.ResponseWriter, req *http.Request) {

	u := getUser(w , req )

	if !(alreadyLoggedIn(req)) {

		http.Redirect(w , req , "/", http.StatusSeeOther)	
	}

	
	
}
