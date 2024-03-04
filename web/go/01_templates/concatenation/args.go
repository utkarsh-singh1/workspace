package concatenation

import (
	"io"
	"log"
	"os"
	"strings"
)


func ArgsHtml() {

	x , y := os.Args[0], os.Args[1]

	str :=  `

<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" >
<meta name="viewport" content="width=device-width intial-scale=1.0">
</head>
<body>
<h1> `+ x + y +` </h1>
</body>
</html>

`

	f ,err := os.Create("index1.html")

	if err != nil {
		log.Fatalln("There is an err", err)
	}

	io.Copy(f, strings.NewReader(str))
}
