package main

import (
	"fmt"
	"log"
	"os"
)

func main() {


	xb , err := readFile("readText.txt")

	if err != nil {
		log.Fatalf("From main there was an error - %s", err)

	}

	fmt.Println(xb)
	fmt.Println(string(xb))
}

func readFile(fileName string) ([]byte,error) {

	xb , err := os.ReadFile(fileName)

	if err != nil {
		return nil, fmt.Errorf("There was an error found in readFile: %s",err)
	}

	return xb, nil
}
