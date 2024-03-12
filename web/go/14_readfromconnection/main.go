package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	li, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Panic("Errror = ",err)

	}

	defer li.Close()

	for {
		conn , err := li.Accept()
		if err != nil {
			log.Println("Error =", err)
		}

		go handle(conn)
	}


}

func handle(conn net.Conn) {

	scanner := bufio.NewScanner(conn)

	for scanner.Scan(){
		ln := scanner.Text()
		fmt.Println(ln)
	}

	defer conn.Close()
}
