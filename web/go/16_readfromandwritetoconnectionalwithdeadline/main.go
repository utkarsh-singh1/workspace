package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	li , err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Panic("Error =", err)
	}

	defer li.Close()

	for {

		conn , err := li.Accept()
		if err != nil {
			log.Println("Error =",err)
		}
		
		go handle(conn)
	}
	
	
}

func handle(conn net.Conn) {

	// SetDeadline set a deadline for the connection, if deadline passed connection automatically got disconnected.
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("Error =",err)
	}

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {

		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "This is my connection %s\n",ln)
	}

	defer conn.Close()

	fmt.Println("Code got here")
}
