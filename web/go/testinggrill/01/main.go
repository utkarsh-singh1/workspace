// 1st test , I will write to the connection and print it on my terminal by stdout

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {

		log.Fatalln("Error =",err)
	}

	defer li.Close()

	for {

		conn, err := li.Accept()
		if err != nil {

			log.Fatalln("Error =",err)
		}

		io.WriteString(conn, "Yes i am writing to the connection")
		writetoconn(conn)
		io.WriteString(conn, "Hola Amigo")

	}

}


func writetoconn(conn net.Conn) {

	io.WriteString(conn, "Hola Amigo")

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {

		ln := scanner.Text()

		fmt.Println(ln)
	}
	
}
