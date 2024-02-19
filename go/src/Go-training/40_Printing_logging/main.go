package main

import (
	"fmt"
	"log"
	"os"
)


func main() {


	// Check for the existing file

	_,err := os.Open("Hello.txt")

	if err != nil {

		//fmt.Println("We have an error Mylord ->",err)
		log.Println("We have an error ->",err)
	}


	// Save an error to an file
	
	f, err := os.Create("log.txt")

	if err != nil {


		fmt.Println("We have an error on our hands ->",err)
		// log.Println("This is an error",err)
		
	}

	defer f.Close()

	log.SetOutput(f)

	if err != nil {
		fmt.Println("There is an error",err)
	}


 
	f2, err := os.Open("hello.txt")

	if err != nil {
		log.Println("An err happened ->",err)
	}

	defer f2.Close()

	fmt.Println("Check out log.txt for detailed info")

	// Fatalln() is equivalent to Println() followed by os.Exit(1)

	// os.Exit(1) runs for exit from the program , non-zero argument means failure of program, zero argument means success

	// defer also does not run if  Fatalln() run.

	_,err = os.Open("file.txt")

	if err != nil {

		log.Fatalln("There is an error occured->",err)
		
	}

	defer func(){
		fmt.Println("It does seems there is an defer")
	}()


	// log.Panic() is equivalent to Println() followed by panic()

	_,err = os.Open("thereisSomething.txt")

	if err != nil {
		log.Panic("There is some error->",err)
	}
}
