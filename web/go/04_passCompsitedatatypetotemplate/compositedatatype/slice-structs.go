package compositedatatype

import (
	"html/template"
	"log"
	"os"
)

var tmp4 *template.Template

func init() {
	
	tmp4 = template.Must(template.ParseFiles("index5.gohtml"))
}


type man struct {
	First string
	Last string
}

func StructsinSlicetoTemp() {

	p1 := man{
		First: "Utkarsh",
		Last: "Singh",
	}

	p2 := man{
		First: "Hari",
		Last: "Om",
	}

	p3 := man{
		First: "Hello",
		Last: "Brother",
	}

	person := []man{p1,p2,p3}

	err := tmp4.Execute(os.Stdout, person)
	if err != nil {
		log.Fatalln("Error =",err)
	}
	
	

}
