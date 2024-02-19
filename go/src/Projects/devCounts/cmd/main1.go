package main

import (
	"fmt"
	"os"
	"strings"
)


func main() {

	fmt.Println(os.Getenv("FOO"))

	for _,v := range os.Environ(){

		pair := strings.SplitN(v,"=",2)
		fmt.Println(pair[0])
		
	}
}
