package main

import (
	"log"
	"os"
)

// type person struct {
// 	first string
// }

// func (p person) writeout(w io.Writer) error {

// 	_,err := w.Write([]byte(p.first))
// 	return err
// }

func main() {

	f,err := os.Create("random.txt")

	if err != nil {
		log.Fatalf("error %s\n", err)
	}

	defer f.Close()
	
	s := []byte("Hello gophers!")
	
	_,err = f.Write(s)

	if err != nil {
		log.Fatalf("error %s", err)
	} 
	
	
}
