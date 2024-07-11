package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Username, First, Last, Role string
	Password []byte
}

var dbSession map[string]string

var dbUser map[string]user

func main() {

	http.HandleFunc("/", index)

	http.HandleFunc("/signup", signUp)

	http.HandleFunc("/login", logIn)

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
		r := req.FormValue("Role")

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

		dbUser[un] = user{un,f,l,bs,r}

		// Redirect Back to Homepage after signup

		http.Redirect(w , req, "/", http.StatusSeeOther)
		return
		
	}
	
}


func logIn(w http.ResponseWriter, req *http.Request) {

	// Check if User is already login
	if alreadyLoggedIn(req) {

		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// if User is not already logged In

	if req.Method == http.MethodPost {

		un := req.FormValue("Usernama")
		p := req.FormValue("Password")

		u,ok := dbUser[un]

		// If user exist or not
		if !ok {
			http.Error(w , "User Does not Exist", http.StatusForbidden)
			return
		}

		// Check if password Matches or not

		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {

			http.Error(w, "Username and Password Don't Match", http.StatusForbidden)
			return
		}


		// Create a session

		sID := uuid.NewV4()

		c := &http.Cookie{

			Name: "session",
			Value: sID.String(),
		}

		http.SetCookie(w , c)
		dbSession[c.Value] = un
		http.Redirect(w , req , "/", http.StatusSeeOther)
		return
	}

}

func bar(w http.ResponseWriter, req *http.Request) {

	u := getUser(w , req )

	if !(alreadyLoggedIn(req)) {

		http.Redirect(w , req , "/", http.StatusSeeOther)	
	}


	if u.Role != "007" {

		http.Error(w, "You are not allowed to enter because you are not 007 buddy ",http.StatusForbidden)
	}
	
	
}

