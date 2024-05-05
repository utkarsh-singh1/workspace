package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", set)

	http.HandleFunc("/favicon.ico", http.NotFound)

	http.ListenAndServe(":8080", nil)

}


func set ( w http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("cookie")
	if err == http.ErrNoCookie {

		cookie = &http.Cookie{
			Name:       "cookie",
			Value:      "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {

		log.Println("error =",err)	
	}

	count++

	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w , cookie)

	io.WriteString(w , cookie.Value )
	
}
