package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type hotel struct{
	Name, Address, City, Region string
	Zip int32
}

func main() {

	Hotels := []hotel{

		hotel{
			"vistara", "sect-24", "lucknow", "central", 2122344,
		},

		hotel{
			"tatabata", "sect-21", "lucknow", "northern", 231435,
		},

		hotel {
			"hellobo", "sect-23", "lucknow", "southern", 345678,
		},
	}

	err := tpl.Execute(os.Stdout, Hotels)
	if err != nil {
		log.Fatalln("Error =",err)
	}
}
