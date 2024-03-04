package compositedatatype

import (
	"html/template"
	"log"
	"os"
)

var tmp *template.Template


func init() {

	tmp = template.Must(template.ParseFiles("index2.gohtml"))
	
}

func SlicetoTemp() {

	sling := []string{"onee", "twooo", "threee", "fooour"}
	
	err := tmp.Execute(os.Stdout, sling)
	if err != nil {
		log.Fatalln("Error =", err)
	}
}
