package compositedatatype

import (
	"html/template"
	"log"
	"os"
)

var tmp5 *template.Template

func init() {
	
	tmp5 = template.Must(template.ParseFiles("index6.gohtml"))
}


type car struct {
	Brand string
	Power int
	Colour string
}

type truck struct {
	Weight int
	Model string
	Capacity int
}

type item struct {
	Cars []car
	Trucks []truck
}

func StructsinSliceinStrcuttoTemp() {

	a := car{
		Brand: "bmw",
		Power: 200,
		Colour: "velvet",
	}

	b := car{
		Brand: "tata",
		Power: 150,
		Colour: "white",
	}

	c := car{
		Brand: "atto",
		Power: 120,
		Colour: "black",
	}

	x := truck{
		Weight: 1000,
		Model: "Loader",
		Capacity: 15000,
	}

	y := truck{
		Weight: 1200,
		Model: "Power",
		Capacity: 19000,
	}

	z := truck{
		Weight: 1400,
		Model: "Maxer",
		Capacity: 21000,
	}


	cars := []car{a,b,c}
	

	trucks := []truck{x,y,z}

	i := item{
		cars,
		trucks,
	}

	err := tmp5.Execute(os.Stdout, i)
	if err != nil {
		log.Fatalln("Error =",err)
	}
	
	

}
