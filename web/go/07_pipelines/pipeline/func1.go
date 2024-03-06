package pipeline

import (
	"log"
	"math"
	"os"
	"text/template"
)


var tmp *template.Template

func double(i int) int {
	return i*2
}

func square(i int) int {
	return int(math.Pow(float64(i), 2))
}

func squareroot(i int) float64 {

	return math.Sqrt(float64(i))
}

var fm = template.FuncMap{
	"fd" : double,
	"fsq" : square,
	"fsqrt" : squareroot,
	
}

func init() {

	tmp = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
	
}

func FuncPipe() {

	err := tmp.ExecuteTemplate(os.Stdout, "index.gohtml", 3)
	if err != nil {
		log.Fatalln("Error =",err)
	}
	
	
}
