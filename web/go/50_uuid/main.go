package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/favicon.ico", http.NotFound )
	http.ListenAndServe(":8080", nil)
	

}


// Create a cookie for a particular session
func foo(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		c = &http.Cookie{
			Name: "session",
			Value: id.String() ,
		}

		http.SetCookie(w, c)

	}

	fmt.Println(c)
}
