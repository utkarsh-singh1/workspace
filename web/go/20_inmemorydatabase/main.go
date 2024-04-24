package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)


func main() {

	li, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Panic("Error =",err)
	}

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println("Error =",err)
		}

		go handle(conn)
	}
	
	
}

func handle(conn net.Conn) {

	defer conn.Close()

	// Write to the connection
	io.WriteString(conn, "\r\nIn-Memory Database\r\n\r\n"+
		"Use:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDELETE key \r\n\r\n"+
		"Example: \r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")

	data := make(map[string]string)

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		if len(fs) < 1{
			continue
		}

		switch fs[0] {

		case "GET" :
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s\r\n",v)
		case "SET" :
			if len(fs) != 3{
				fmt.Fprintln(conn, "Expected a value\r\n")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v

		case "DELETE" :
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintln(conn, "Invalid Command "+fs[0]+"\r\n")
	}
	}
}
