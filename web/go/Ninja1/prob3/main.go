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

type breakfast struct {
	Time string
	Item []string
}


type lunch struct {
	Time string
	Item []string
}

type dinner struct {
	Time string
	Item []string
}

type menu struct {
	Morning breakfast
	Noon lunch
	Night lunch
}

func main() {

	foods := menu{
		Morning: breakfast{Time: "9:00 AM", Item: []string{"bread", "Jam"}},
		Noon:    lunch{Time: "1:00 PM", Item: []string{"bread", "curry", "rice"}},
		Night:   lunch{Time: "7:00 PM", Item: []string{"bread", "curry"},},
	}

	err := tpl.Execute(os.Stdout, foods)
	if err != nil {
		log.Fatalln("Error =",err)
	}
	
}
