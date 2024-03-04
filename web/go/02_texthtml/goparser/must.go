package goparser

import (
	"os"
	"text/template"
)


var tmp *template.Template

func init() {

	// In case of parseFile and parseglob it return a *template and error but Must() returns only *template and do the error checking for us.

	// func Must(t *Template, err error) *Template {
	// if err != nil {
	// 	panic(err)
	// }
	// return t
	//}

	tmp = template.Must(template.ParseGlob("template/*"))
	

}


func ParseMust(){

	tmp.Execute(os.Stdout, nil)
	
}
