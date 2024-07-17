package main

import (
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Username, First, Last, Role string
	Password []byte
}

type session struct {

	Username string
	lastActivity time.Time 
}



var dbSession map[string]session

var dbUser map[string]user

const sessionLength int = 30

var dbSessionCleaned time.Time

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

func bar(w http.ResponseWriter, req *http.Request) {

	u := getUser(w , req )

	if !(alreadyLoggedIn(req)) {

		http.Redirect(w , req , "/", http.StatusSeeOther)	
	}

	if u.Role != "007" {

		http.Error(w, "Can not enter , first become 007 buddy ", http.StatusForbidden)
	}
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
		c.MaxAge = sessionLength
		http.SetCookie(w , c)

		// Attach session to a Username

		dbSession[c.Value] = session{un,time.Now()}

		// Encrypt Password

		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost )
		if err != nil {

			http.Error(w , "Internal Server Error", http.StatusInternalServerError)
		}

		// Add userdetail to username

		dbUser[un] = user{un,f,l,r,bs}

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
		c.MaxAge = sessionLength
		http.SetCookie(w , c)
		dbSession[c.Value] = session{un, time.Now()}
		http.Redirect(w , req , "/", http.StatusSeeOther)
		return
	}

}


func logOut ( w http.ResponseWriter, req *http.Request) {

	// Go back to your home ,to your families 
	if !alreadyLoggedIn(req) {

		http.Redirect(w , req , "/", http.StatusSeeOther)
	}

	// Delete the session from the dbSession
	c, _ := req.Cookie("session")
	delete(dbSession, c.Value )

	// delete the cookies

	c = &http.Cookie{
		Name: "session",
		Value: "",
		MaxAge: -1,
	}

	http.SetCookie(w, c)

	// Go back to Home Page

	http.Redirect(w, req, "/login", http.StatusSeeOther)

}


