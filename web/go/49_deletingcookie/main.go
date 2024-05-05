package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", index)

	http.HandleFunc("/set", set)

	http.HandleFunc("/read", read)

	http.HandleFunc("/expire", expire)

	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(w, `<h1><a href="/set">set a cookie</a></h1>`)

}

func set(w http.ResponseWriter, req *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:  "hello",
		Value: "world",
	})

	fmt.Fprintln(w, `<h1><a href="/read">read the cookie</a></h1>`)
}

func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("hello")
	if err != nil {

		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return

	}

	fmt.Fprintln(w, `<h1>Cookie key is</h1>`, c, "<br />", `Wanna delete a cookie press this <a href="/expire">link  </a>`)

}

func expire(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("hello")
	if err != nil {

		http.Redirect(w, req, "/set", 303)
		return
	}

	c.MaxAge = -1

	http.SetCookie(w , c)
	http.Redirect(w , req, "/", 303)

}
