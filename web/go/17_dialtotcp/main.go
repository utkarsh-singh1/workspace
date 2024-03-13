package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:5050")
	if err != nil {
		log.Panic("Error =",err)
	}

	defer conn.Close()

	bs , err := io.ReadAll(conn)
	if err != nil {
		log.Println("Error =",err)
	}

	fmt.Println(string(bs))
}
