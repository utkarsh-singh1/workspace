package main

import (
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

// getUser checks if User exst in dB or not
func getUser(w http.ResponseWriter, req *http.Request) user {

	c, err := req.Cookie("session")
	if err != nil {

		sID := uuid.NewV4()
		c = &http.Cookie{

			Name: "session",
			Value: sID.String(),
		}

	}

	// set Timer for cookie

	c.MaxAge = sessionLength

	http.SetCookie(w , c)


	// Check if user exist or not
	var u user

	if s, ok := dbSession[c.Value] ; ok {

		s.lastActivity = time.Now()

		dbSession[c.Value] = s
		
		u = dbUser[s.Username]

	}

	return u
}

func alreadyLoggedIn(req *http.Request) bool {

	c, err := req.Cookie("session")

	if err != nil {

		return false
	}

	
	// if User Exist or not
	un  := dbSession[c.Value]

	_, ok := dbUser[un]

	return ok
}
