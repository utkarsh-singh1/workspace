package compositedatatype

import (
	"html/template"
	"log"
	"os"
)


var tmp1 *template.Template

func init() {

	tmp1 = template.Must(template.ParseFiles("index3.gohtml"))
}


func MaptoTemp() {

	mapped := map[string]string{

		"1" : "Oonee",
		"2" : "Twwooo",
	}
	
	err := tmp.Execute(os.Stdout, mapped)
	if err != nil {

		log.Fatalln("Error =",err)
		
	}
	
}
