package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)


func main(){

	li ,err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Panic(err)
	}

	defer li.Close()


	// // What if you try to run code without putting below code in for loop, after an request seend to tcp server, the server stops listening
	// conn, err := li.Accept()

	// fmt.Fprintln(conn,"Hello from my side" )


	for{
		conn , err := li.Accept()
		if err != nil {
			log.Println("Error =")

		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()

		fmt.Println(ln)

		fmt.Fprintf(conn, "Hey I'm writing to the connection %s", ln)
	}

	defer conn.Close()
} 
