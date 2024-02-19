package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeout(w io.Writer) {

	w.Write([]byte(p.first))
}

func main() {

	p := person{
		first: "Utkarsh",
	}

	f, err := os.Create("random.txt")

	if err != nil {
		log.Fatalf("error %s\n", err)
	}

	defer f.Close()

	p.writeout(f)

	var b bytes.Buffer

	p.writeout(&b)

	fmt.Println(b.String())

	o := os.Environ()

	fmt.Println(o)
}
