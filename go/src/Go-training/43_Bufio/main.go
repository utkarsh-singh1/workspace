package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f , err := os.Open("file.txt")

	check(err)

	b := []byte{}

	fmt.Println(f.Read(b))
	
	defer f.Close()


	scanner := bufio.NewScanner(f)

	fmt.Println(scanner)
	
	
	
	newReader := bufio.NewReader(f)
	
	data := make([]byte,100)

	_, err = newReader.Read(data)

	check(err)

	fmt.Println(string(data))

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		
	}
}

// func openFile(path string) (*os.File) {

// 	f, err := os.Open(path)


// 	f.WriteString("My Name is Utkarsh, Just testing function")
	
// 	if os.IsNotExist(err) {
// 		f , _ = os.Create(path)
// 		f.WriteString("My Name is Utkarsh, Just testing function")
// 	}
	
	
// 	return f
	
// }
