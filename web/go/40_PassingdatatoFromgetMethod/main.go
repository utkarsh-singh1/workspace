package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", pickit)
	http.ListenAndServe(":8080", nil )

}

func pickit(w http.ResponseWriter, req *http.Request) {

	vs := req.FormValue("q")


	// If Method is Post , the data or value will through body section of http Request
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w ,`
<form method="GET">
<input type="text" name="q" />
<input type=submit />
</form>
<br />
`+vs )

}
