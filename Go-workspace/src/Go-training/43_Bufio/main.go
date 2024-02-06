package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f := openFile("file.txt")

	reader := bufio.NewScanner(f)

	defer f.Close()
	
	fmt.Println(reader)

	

	
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		
	}
}

func openFile(path string) (*os.File) {

	f, err := os.Open(path)


	f.WriteString("My Name is Utkarsh, Just testing function")
	
	if os.IsNotExist(err) {
		f , _ = os.Create(path)
		f.WriteString("My Name is Utkarsh, Just testing function")
	}
	
	
	return f
	
}
