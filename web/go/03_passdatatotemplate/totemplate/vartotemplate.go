package totemplate

import (
	"os"
	"text/template"
)

var tmp1 *template.Template

func init(){

	tmp1 = template.Must(template.ParseFiles("index1.gohtml"))
	
}

func VartoTemp() {

	tmp1.Execute(os.Stdout, `There is an essence of life that everobody should know "Enjoy it while you can it not gonna last forever"`)
	
}
