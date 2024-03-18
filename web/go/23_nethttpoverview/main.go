package main

import (
	"fmt"
	"log"
	"net/http"
)

type person int

func (p person) ServeHTTP(w http.ResponseWriter, r *http.Request){

	body := `<!DOCTYPE HTML><html lang="en-us"><head><meta charset="utf-8" /><title></title></head>
<body>
Hello World
</body>
</html>
`
	fmt.Fprintln(w, body)
}

func main(){

	var p person
	
	err := http.ListenAndServe(":8080",p)
	if err != nil {
		log.Panic("Error =",err)
	}
	
}

