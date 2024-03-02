package concatenation

import (
	"io"
	"log"
	"os"
	"strings"
)


func CreateHtml(){

	name := "hello world"

	str := `

<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" >
<meta name="viewport" content="width=device-width intial-scale=1.0">
</head>
<body>
<h1> `+ name +` </h1>
</body>
</html>

`

	f, err := os.Create("index.html")

	if err != nil {
		log.Fatalln("Error ->",err)
	}

	defer f.Close()

	io.Copy(f, strings.NewReader(str))
}
