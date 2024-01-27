package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	// Almost never ever ignore errors

	n , err := fmt.Println("This is an error")

	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(n)


	// Important - file operations 

	// f , err := os.Create("hello.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// r := strings.NewReader("Wassup buddy")

	// io.Copy(f, r)

	// Open a file

	f1, err := os.Open("hello.txt")

	if err != nil {
		log.Fatalln(err)
	}

	defer f1.Close()


	bs , err := io.ReadAll(f1)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
	
}


