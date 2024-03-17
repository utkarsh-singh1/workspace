package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic("Error =",err)
	}

	defer li.Close()
	
	for {

		conn, err := li.Accept()
		if err != nil {
			log.Println("Error =",err)
		}

		go handle(conn)
		
	}

}

func handle(conn net.Conn) {
	//defer conn.Close()

	// Request to server
	request(conn)

	// Response to server
	response(conn)
}

func request(conn net.Conn) {
	i := 0

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {

		ln := scanner.Text()
		fmt.Println(ln)
		// request
		if i == 0 {
			m := strings.Fields(ln)[0]
			n := strings.Fields(ln)[1]
			fmt.Println("***Methods",m)
			fmt.Println("****URL",n)
		}
		if ln == "" {
			break
		}
		i++
	}
	
}

func response(conn net.Conn) {

	res := `<!DOCTYPE html>
<html lang="en-us">
<head>
<title></title>
</head>
<body>
<strong>Hello World</strong>
</body>
</html>
`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(res))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, res)
}



