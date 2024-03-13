package main

import (
	"fmt"
	"log"
	"net"
)


func main() {

	conn, err := net.Dial("tcp", "localhost:5050")
	if err != nil {
		log.Panic("Error =",err)
	}

	defer conn.Close()

	fmt.Fprintln(conn, "Hey i dialed you")
}
