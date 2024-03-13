package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li ,err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Panic("Error =",err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println("Error =",err)
		}

		// Write to connection

		io.WriteString(conn, "\nHello from the tcp server\n")

		fmt.Fprintln(conn, "How is your day?")

		fmt.Fprintf(conn, "%v", "well i hope!\n")

		conn.Close()
	}
}
