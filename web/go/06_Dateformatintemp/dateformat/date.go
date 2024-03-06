package dateformat

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tmp *template.Template


func monthdateyearformat(t time.Time) string {
	return t.Format("03-04-2005")
}

var fm = template.FuncMap{
	"dfmdy" : monthdateyearformat,
}


func init() {

	tmp = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func Dateformat() {

	err := tmp.ExecuteTemplate(os.Stdout,"index,gohtml",time.Now())
	if err != nil {
		log.Fatalln("Error =",err)
	}
}
