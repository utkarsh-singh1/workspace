package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)


func main() {


	http.HandleFunc("/", foo)

	http.HandleFunc("/bar", bar)
	
	http.ListenAndServe(":8080", nil )
	
}

func foo(w http.ResponseWriter, req *http.Request) {

	c , err := req.Cookie("cookie")
	if err != nil {

		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:"session" ,
			Value: sID.String() ,
		}
		
	}

	http.SetCookie(w , c )
	
}

func bar(w http.ResponseWriter, req *http.Request){

}
