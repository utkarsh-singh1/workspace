package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) user {

	c, err := req.Cookie("session")
	if err != nil {

		sID := uuid.NewV4()
		c = &http.Cookie{

			Name: "session",
			Value: sID.String(),
		}

	}

	http.SetCookie(w , c)

	var u user

	if un, ok := dbSession[c.Value] ; ok {

		u = dbUser[un]

	}

	return u
}

func alreadyLoggedIn(req *http.Request) bool {

	c, err := req.Cookie("session")

	if err != nil {

		return false
	}

	

	un  := dbSession[c.Value]

	_, ok := dbUser[un]

	return ok
}
