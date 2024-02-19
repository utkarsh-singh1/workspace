package main

import (
	"fmt"
	"os"
)

func main() {

	os.Setenv("FOO", "1")
	fmt.Println("Get env variable FOO:",os.Getenv("FOO"))
	
}
